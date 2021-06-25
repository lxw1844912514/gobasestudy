package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80") //与google 简历tcp 链接
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // 退出关闭链接

	//通过 链接发送get请求,访问首页
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dat))

}
