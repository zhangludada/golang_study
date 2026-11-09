package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("recvStr:", recvStr)
		conn.Write([]byte("ok")) // 把收到的数据返回给客户端
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 启动 gotine 去处理链接
		go process(conn)
	}

}
