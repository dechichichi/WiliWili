package utils

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"
	"wiliwili/kitex_gen/chat"

	"github.com/gorilla/websocket"
)

// 存储在线用户的 WebSocket 连接信息
var onlineUsers = make(map[string]*websocket.Conn)
var mu sync.RWMutex

// 处理用户上线
func UserOnline(uid string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	onlineUsers[uid] = conn
}

// 处理用户下线
func UserOffline(uid string) {
	mu.Lock()
	defer mu.Unlock()
	if conn, exists := onlineUsers[uid]; exists {
		conn.Close()
		delete(onlineUsers, uid)
	}
}

// 检查用户是否在线
func IsUserOnline(uid string) bool {
	mu.RLock()
	defer mu.RUnlock()
	_, exists := onlineUsers[uid]
	return exists
}

// 向特定用户发送消息
func SendMessageToUser(uid string, msg string) error {
	mu.RLock()
	defer mu.RUnlock()
	if conn, exists := onlineUsers[uid]; exists {
		resp := &chat.ChatResp{
			Uid:       uid,
			Msg:       msg,
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		}
		// 将响应结构体转换为 JSON 格式
		respJSON, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		// 使用 WebSocket 连接发送消息
		err = conn.WriteMessage(websocket.TextMessage, respJSON)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
