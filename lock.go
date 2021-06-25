package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

//使用channel 在多个goroutine 之间进行通信,较为常用通信方式是共享内存.

var name string

func printName() {
	log.Println("name is ", name)
}

func main() {
	//1.这就是最简单的通过共享变量(内存)的方式在多个goroutine 进行通信的方式
	name = "小明"
	go printName()
	go printName()
	time.Sleep(time.Second)

	name = "小红"
	go printName()
	go printName()
	time.Sleep(time.Second)

	//2.并发对同一个切片进行写操作的时候,会出现数据不一致的问题,这就是一个典型的共享变量的问题.
	//使用lock(锁)来修复,最后都包含0-9这10个数字

	//在go中还有一种读写锁,sync.RwMutex,对于我们的共享对象,如果可以分离出读和写两个互斥信号的情况,
	//可以考虑使用它来提高读的并发性能
	/*	var (
			wg      sync.WaitGroup
			numbers []int
			mux     sync.Mutex //互斥锁,只有一个信号标量,

		)
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {

				mux.Lock()
				numbers = append(numbers, i)
				mux.Unlock()

				wg.Done()
			}(i)

		}
		wg.Wait()
		fmt.Println("The numbers is", numbers)*/

	//3.
	var (
		mux    sync.Mutex
		state1 = map[string]int{
			"a": 65,
		}
		muxTotal uint64

		rw     sync.RWMutex
		state2 = map[string]int{
			"a": 65,
		}
		rwTotal uint64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				_ = state1["a"]
				mux.Unlock()
				atomic.AddUint64(&muxTotal, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				rw.RLock()
				_ = state2["a"]
				rw.RUnlock()
				atomic.AddUint64(&rwTotal, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("sync.Mutex readOps is", muxTotal)  //sync.Mutex readOps is 234 6154
	fmt.Println("sync.RwMutex readOps is", rwTotal) //sync.RwMutex readOps is 1780 6741
	//可以看到使用sync.RwMutex 的读的并发能力大概是 sync.Mutex 的十倍,从而大大提高了其并发能力

}
//总结:
//我们可以通过共享内存的方式实现多个goroutine 中的通信.
//	多个goroutine 对于共享的内存进行写操作的时候,可以使用Lock来避免数据不一致的情况.
//	对于可以分离为读写操作的共享数据可以考虑使用 sync,RWMutex 来提高其读的并发能力.