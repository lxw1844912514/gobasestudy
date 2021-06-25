package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//我们先使用 net 包创建一个 TCP Server ，然后尝试连接 Server,
// 最后再通过客户端发送 hello 到 Server，同时 Server 响应 word
func main() {
	//通过 `net.Listen("tcp", "127.0.0.1:8888")` 新建一个 TCP Server。
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start server with : %s", l.Addr())

	defer l.Close()

	for {
		//通过 `l.Accept()` 获取创建的连接。
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 通过 `go handleConnection(c)` 新建的 goroutine 来处理连接。
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		dat, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("client:", string(dat))

		_, err = conn.Write([]byte("world\n"))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
//2021/06/24 21:04:39 Start server with : 127.0.0.1:8888
//client: hello
//client: hello