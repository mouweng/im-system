package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex

	// 消息广播channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP : ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

func (this *Server) ListenMessager() {
	for {
		msg := <- this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// 当前用户上线, 将用户加入onlineMap
	user := NewUser(conn)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	this.BroadCast(user, "已上线")

	select{}
}

// 启动服务器接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessager()
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}

}