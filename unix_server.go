package main

import (
	"log"
	"net"
)

func main() {
	//使用 `net.Listen("unix", "/tmp/unix.sock")` 启动一个 Server。
	l, err := net.Listen("unix", "/tmp/unix2.sock")
	if err != nil {
		log.Fatal("listen error", err)
	}
	for {
		//使用 `conn, err := l.Accept()` 来接受客户端的连接。
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error", err)
		}
		//使用 `go helloServer(conn)` 来处理客户端连接，并读取客户端发送的数据 `hi` 并返回 `hello`
		go helloServer(conn)
	}
}

func helloServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		log.Println(string(data))

		_, err = c.Write([]byte("hello"))
		if err != nil {
			log.Fatal("write:", err)
		}
	}

}
