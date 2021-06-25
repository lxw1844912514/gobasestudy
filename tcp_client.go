package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//通过 `net.DialTCP("tcp", nil, addr)` 尝试创建到 TCP Sever 的连接。
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		//通过 `conn.Write([]byte("hello\n"))` 向服务端发送数据。
		_, err := conn.Write([]byte("hello\n"))
		if err != nil {
			log.Fatal(err)
		}

		//通过 `reader.ReadLine()` 读取服务端响应数据。
		dat, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("server:", string(dat))
		time.Sleep(time.Second * 5)

	}
}

//bogon:web liutao$ go run tcp_client.go
//server: world
//server: world
//server: world
//server: world