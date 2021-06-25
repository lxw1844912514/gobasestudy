package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func doSomething(id int) { //总耗时9秒,每个任务是阻塞的
	log.Printf("before do job :(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job :(%d) \n", id)
}

func doSomething2(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("before do job :(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job :(%d) \n", id)
}

//使用goroutine 从9S 降低到3秒,提高了我们的效率,根据打印的输出结果得出:
//多个goroutine 的执行是随机的
//对于io密集型任务特别有效,比如文件,网络读写
func main() {
	/*go doSomething(1)
	go doSomething(2)
	go doSomething(3)
	time.Sleep(4 * time.Second)*/

	var wg sync.WaitGroup
	wg.Add(3) //fatal error: all goroutines are asleep - deadlock!

	go doSomething2(1, &wg)
	go doSomething2(2, &wg)
	go doSomething2(3, &wg)

	wg.Wait()
	log.Printf("finish all jobs \n")

	fmt.Println("========")

	//通过方法传参的方式,将i的值拷贝到新的变量v中,而在每个goroutine都对应了一个属于自己作用域的v变量,
	for i := 0; i < 3; i++ {
		go func(v int) {
			//fmt.Println(i)
			log.Printf("I is : %d \n", v)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
