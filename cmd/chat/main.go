package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/websocket"
)

var addr = flag.String("addr", "0.0.0.0:6666", "http service address")

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true // 允许所有跨域请求
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type wsConnection struct {
	wsSocket  *websocket.Conn // 底层websocket
	inChan    chan *wsMessage // 读队列
	outChan   chan *wsMessage // 写队列
	mutex     sync.Mutex      // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

// 用户连接映射
var userConnections = make(map[string]*wsConnection)
var userMutex sync.Mutex

// 消息格式
type Message struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func (wsConn *wsConnection) wsReadLoop() {
	for {
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			log.Println("read message error:", err)
			wsConn.wsClose()
			return
		}
		req := &wsMessage{
			messageType: msgType,
			data:        data,
		}
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			log.Println("close chan received, exiting read loop")
			return
		}
	}
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		select {
		case msg := <-wsConn.outChan:
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				log.Println("write message error:", err)
				wsConn.wsClose()
				return
			}
			msg.data = nil // 显式释放内存
		case <-wsConn.closeChan:
			log.Println("close chan received, exiting write loop")
			return
		}
	}
}

func (wsConn *wsConnection) wsClose() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
		close(wsConn.inChan)  // 关闭读队列
		close(wsConn.outChan) // 关闭写队列
	}
}

func wsHandler(_ context.Context, c *app.RequestContext) {
	// 获取用户标识（例如从 URL 参数中获取）
	userID := c.Query("user")
	if userID == "" {
		log.Println("user ID is required")
		return
	}

	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		wsConn := &wsConnection{
			wsSocket:  conn,
			inChan:    make(chan *wsMessage, 1000),
			outChan:   make(chan *wsMessage, 1000),
			closeChan: make(chan byte),
			isClosed:  false,
		}
		// 将用户连接存储到映射中
		userMutex.Lock()
		userConnections[userID] = wsConn
		userMutex.Unlock()

		// 处理器
		go wsConn.procLoop()
		// 读协程
		go wsConn.wsReadLoop()
		// 写协程
		go wsConn.wsWriteLoop()

		// 等待用户断开连接
		<-wsConn.closeChan
		// 从映射中移除用户连接
		userMutex.Lock()
		delete(userConnections, userID)
		userMutex.Unlock()
	})
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
}

func (wsConn *wsConnection) procLoop() {
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			log.Println("read fail:", err)
			wsConn.wsClose()
			return
		}
		// 解析消息
		var message Message
		if err := json.Unmarshal(msg.data, &message); err != nil {
			log.Println("invalid message format:", err)
			continue
		}

		// 查找目标用户连接
		userMutex.Lock()
		targetConn, exists := userConnections[message.To]
		userMutex.Unlock()

		if !exists {
			log.Printf("target user %s not found", message.To)
			continue
		}

		// 将消息转发到目标用户
		err = targetConn.wsWrite(msg.messageType, []byte(message.Message))
		if err != nil {
			log.Println("write fail:", err)
			wsConn.wsClose()
			return
		}
	}
}

func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
		return nil, errors.New("websocket closed")
	}
}

func main() {
	flag.Parse()
	h := server.Default(server.WithHostPorts(*addr))
	h.NoHijackConnPool = true
	h.GET("/ws", wsHandler)
	h.Spin()
}
