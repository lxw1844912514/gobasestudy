package main

import (
	"errors"
	"fmt"
)

//1.defer
// defer 通常用于延迟调用指定的函数,被defer 调用的函数称为"延迟函数"
//defer 常用场景:
//defer 常被用于处理成对的操作,如打开和关闭,链接和断开链接,加锁和释放锁.恰当使用defer 能保证资源正确释放.

func outFunc() {
	defer fmt.Println(" world \n")
	fmt.Print("Hello")
}

func testDefer() (i int) {
	defer func() {
		fmt.Println(i)
		i = 4
	}()

	i = 2
	return i
}

func printNum() {
	for i := 0; i < 5; i++ {
		//1. defer fmt.Println(i) //4 3 2 1 0

		//2.
		/*defer func() {
			fmt.Println(i) //5 5 5 5 5 等到执行 defer 调用的函数时i 已经是5 了
		}()*/

		//3.
		defer func(v int) {
			fmt.Println(v) //4 3 2 1 0 等同于第一种情况,延迟直接打印值
		}(i)

	}
}

//2.panic
func panicFunc() {
	defer func() {
		//recover 返回的是一个interface类型的结果,如果捕获了panic事件,该结果就为非nil
		if p := recover(); p != nil {
			fmt.Println("panic info: ",p)
			fmt.Println("recover panic")
		}
	}()
	panic(errors.New("this is a test for panic"))
}

//recover
//recover 函数能使当前程序从panic中恢复,recover能够拦截panic事件,使得程序不会因为意外而触发panic事件而完全退出


func printNumbers(){
	for i:=0;i<5;i++{
		defer func( n int) {
			fmt.Printf("I is %v ,res is %v \n",i,n)
		}(i*2)
	}
}

func printArr(){
	defer func() {
		if f:=recover();f!=nil{
			fmt.Printf("超出数组长度越界:%v \n",f)
		}
	}()
	//当程序遇到致命错误导致无法继续运行时就会触发panic,如数组越界,空指针
	s := []int{1, 2, 3}
	for i := 0; i <= 4; i++ {
		fmt.Printf("数组s的第%v元素: %v \n ",i,s[i])
	}
}
func main() {

	outFunc()
	//printNum()

	printNumbers()

	testDefer() //2

	printArr()


	fmt.Println("before panic")
	panicFunc() //主动调用panic函数
	fmt.Println("after panic")

}
