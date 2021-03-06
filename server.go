package main

import (
	"fmt"
	"net"
	"sync"
	"io"
	"time"
)

type Server struct {
	IP string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	// 锁
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

// 服务端监听消息
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

// 广播用户消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 业务入口
func (this *Server) Handler(conn net.Conn) {
	// 当前用户上线, 将用户加入onlineMap
	user := NewUser(conn, this)
	user.Online()
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			// 提取用户消息(去除'\n')
			msg := string(buf[:n - 1])
			// 针对message进行消息处理
			user.DoMessage(msg)
			isLive <- true
		}
	}()
	
	for {
		select{
		case <- isLive :
			// 当前用户是活跃的，重制定时器
			// 不做任何事情，为了激活select，更新定时器
		case <-time.After(time.Second * 60):
			// 超时
			user.SendMsg("你已超时，程序关闭...\n")
			// 销毁资源
			close(user.C)
			conn.Close()
			return
		}
	}
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