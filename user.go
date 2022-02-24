package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn

	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
		server: server,
	}
	// 启动监听user
	go user.ListenMessage()
	return user
}

// 用户上线业务
func (this *User) Online() {
	// 当前用户上线, 将用户加入onlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) Offline() {
	// 当前用户下线, 将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播用户下线消息
	this.server.BroadCast(this, "已下线")
}

func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}


// 监听当前User，一旦有消息，直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <- this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}