package main

import "fmt"

func main() {
	/*for i:=1;i<=5 ;i++  {
		defer fmt.Println(i)
	}
	defer fmt.Println("出现在 5 之前")
	defer fmt.Println("第二早")
	fmt.Println("before defer:")*/

	//	闭包函数
	/*defer func() {
		fmt.Println("After ")
	}()

	//等价于
	f := func() {
		fmt.Println("After 2")
	}
	defer f()

	//	go 异常处理
	panic("I am wrong !") //抛出异常不影响,后续的defer 操作,在退出之前执行defer

	fmt.Println("before")*/

	//	函数也可以做为值,类型
	type sum func(x, y int) int

	var f sum = func(x, y int) int {
		return x + y
	}
	fmt.Println(f(3, 4))
}
