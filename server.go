package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP : ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	// 业务
	fmt.Println("连接建立成功")
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