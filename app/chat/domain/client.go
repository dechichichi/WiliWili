package domain

import "golang.org/x/net/websocket"

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	// 新增字段
	uid int
}
