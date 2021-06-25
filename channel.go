package main

import (
	"fmt"
	"time"
)

//channel 是go 中实现并发的重要机制,channel 是goroutine 之间进行通信的重要桥梁.

//1. 使用内建函数make创建channel
//ch:=make(chan int) //channel 必须定义其传递的数据类型
//也可以用var 声明channel
//var ch chan int
//以上声明的channel 都是双向的,可以向该channel 发送数据,也可以接受数据
//发送和接受是channel 的两个基本操作
//ch <-x //channel 接受数据x
//x <-ch //channel 发送数据并赋值给x
//<- ch //channel 发送数据,忽略接受者

//注意的是，程序中必须同时有不同的 goroutine 对非缓冲通道进行发送和接收操作，否则会造成阻塞。
//main 函数是一个 goroutine, 在这一个 goroutine 中发送了数据给非缓冲通道，但是却没有另外一个 goroutine 从非缓冲通道中里读取数据， 所以造成了阻塞或者称为死锁

//2.单向channel :即限定了该channel 只能接受或者发送数据,单向通道通常作为函数的参数



func receive(receiver chan<- string, msg string) {
	receiver <- msg
}
func send(sender <-chan string, receiver2 chan<- string) {
	msg := <-sender
	receiver2 <- msg
	//receiver2 <- sender
	//fmt.Println(receiver2)
}

func strWorker(ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("do something with strWorker")
	ch <- "str"
}
func intWorker(ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("do something with intWorker")
	ch <- 1
}

func worker(done chan bool){
	fmt.Println("start working")
	done<-true
	fmt.Println("end working")
}

func main() {
	/*ch := make(chan string)
	go func() { //ge
		ch <- "ping"
	}()
	fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock!*/

	/*ch := make(chan int, 3)
	ch <- 1 //chan 接受数据 1
	ch <- 2
	ch <- 3
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) // 发送数据
	}*/

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	receive(ch1, "pass message")
	send(ch1, ch2)
	fmt.Println(<-ch2)

	//3.channel 遍历和关闭 close 函数用于关闭channel ,关闭后channel有缓冲数据,依然可以读取,
	// 但无法再发送数据给已经关闭的channel
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	res := 0
	for v := range ch {
		res += v
	}
	fmt.Println(res)

	//4.select 语句
	//select 专门用于通道发送和接受操作,看起来和switch相似,但是进行选择和判断的方法完全不同
	//通过select 的使用,保证了worker 中事务可以执行完毕后才退出main函数

	/*chStr := make(chan string)
	chInt := make(chan int)

	go strWorker(chStr)
	go intWorker(chInt)
	for i := 0; i < 2; i++ {
		select {
		case <-chStr:
			fmt.Println("get value from strWorker")
		case <-chInt:
			fmt.Println("get value from intWorker")
		}
	}*/


	done:=make(chan bool ,1)
	go worker(done)
	fmt.Println(<-done) //注释掉后,main 函数会提前退出
}
