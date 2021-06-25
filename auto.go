package main

import (
	"fmt"
	"sync/atomic"
	"time"
)
//一.原子操作

//automic 提供的原子操作能够确保任一时刻只有一个 goroutine 对变量进行操作,善用atomic能够避免程序中出现大量的锁操作
//automic 常见的操作有:增减, 载入,比较并变换,交换,存储
//例子: func AddInt32(addr *int32,delta int32)(new int32)
//第一个参数必须是指针类型的值,通过指针变量可以获取被操作数在内存中的地址,确保同一时间只有一个goroutine能够进行操作.

//1.使用锁和原子操作来实现多个goroutine 对同一个变量进行累加操作

/*func main() {
	var (
		mux   sync.Mutex
		total uint64
	)
	for i := 0; i < 1; i++ {
		go func() {
			for{
				mux.Lock()
				total += 1
				mux.Unlock()
				time.Sleep(time.Microsecond)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("the total numbers is", atomic.LoadUint64(&total))
}*/

//2.使用atomic 实现,累加操作
func main(){
	var total int64
	for i:=0;i<10 ;i++  {
		go func() {
			for{
				atomic.AddInt64(&total,-1)
				time.Sleep(time.Microsecond)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("the total numbers is ",total)
}

//二.载入操作:
//atomic 包中提供了如下以Load 为前缀的载入操作
//例如: func LoadInt32(addr *int32)(val int32)
//载入操作能够保证原子的读变量的值,当读取的时候,任何其他goroutine 都无法对该变量进行读写

//三.比较并交换:CAS(compare And Swap) 这类操作的前缀为CompareAndSwap

//总结:
//atomic 和锁的方式其性能没有太大区别
//atomic 写法相交于使用锁,更简单


