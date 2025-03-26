package domain

import "sync"

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client

	// 记录 uid 跟 client 的对应关系
	userClients map[int]*Client

	// 读写锁，保护 userClients 以及 clients 的读写
	sync.RWMutex
}
