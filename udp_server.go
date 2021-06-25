package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//1. 监听本地 UDP `127.0.0.1:8888`。
	pc, err := net.ListenPacket("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start server with: %s", pc.LocalAddr())

	defer pc.Close()

	clients := make([]net.Addr, 0)

	go func() {
		for {
			for _, addr := range clients {
				//4.通过 `pc.WriteTo([]byte("pong\n"), addr)` 向所有客户端发送消息。
				_, err := pc.WriteTo([]byte("pong\n"), addr)
				if err != nil {
					log.Println(err)
				}
			}
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		buf := make([]byte, 256)
		//2.使用 `pc.ReadFrom(buf)` 方法读取客户端发送的消息。
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		//3.使用 `clients` 来保存所有连上的客户端连接。
		clients = append(clients, addr)

		log.Println(string(buf[0:n]))
		log.Println(addr.String(), "connecting...", len(clients), "connected")
	}
}

//2021/06/24 21:00:43 start server with: 127.0.0.1:8888
//2021/06/24 21:00:48 ping..
//2021/06/24 21:00:48 127.0.0.1:63609 connecting... 1 connected
//2021/06/24 21:00:52 ping..
//2021/06/24 21:00:52 127.0.0.1:50738 connecting... 2 connected
//2021/06/24 21:01:32 ping..
//2021/06/24 21:01:32 127.0.0.1:65461 connecting... 3 connected