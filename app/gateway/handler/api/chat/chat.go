package chat

import (
	"context"
	"html/template"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
)

var upgrader = websocket.HertzUpgrader{} // use default options

// HandleWebSocket 处理 WebSocket 连接
func Echo(_ context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
}

func Home(_ context.Context, c *app.RequestContext) {
	c.SetContentType("text/html; charset=utf-8")
	homeTemplate.Execute(c, "ws://"+string(c.Host())+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse("home.html"))
