package main

import "fmt"

//函数是语句序列的集合,能够将一个大的工作分解成小的任务,对外隐藏了实现细节
//函数组成: 函数名,参数列表(parameter-list), 返回值(result-list) ,函数体(body)

//语法
/*func name(param-list)(result-list)  {
	body
}*/

//1. 单返回值函数
func plus(a, b int) (res int) {
	return a + b
}

//2.多返回值函数
func multi() (string, int) {
	return "lxw", 18
}

//3.命名返回值
func namedReturnValue() (name string, height int) {
	name = "xiaoming"
	height = 180
	return
}

//4.参数可变函数
func sum(nums ...int) int {
	fmt.Println("\n len of nums is:", len(nums))
	res := 0

	for _, v := range nums {
		res += v
	}
	return res
}

//6.闭包函数
func addInt(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

//7.函数作为参数
func sayHello(name string) {
	fmt.Println("Hello ", name)
}
func logger(f func(string), newName string) {
	fmt.Println("开始调用sayHello:")
	f(newName)
	fmt.Println("结束调用 sayHello \n")
}

//8. 传值和传引用
func sendValue(name string)  {
	name="lxw"
}

func sendAddress(name *string){ //传引用
	*name="lxw1234"
}

func main() {
	fmt.Println("单返回值函数plus: ", plus(1, 2)) //3

	name, age := multi()
	fmt.Printf("多返回值函数multi:[name:%v age:%v]\n", name, age)
	name2, _ := multi()
	fmt.Printf("多返回值函数multi:[name:%v ]\n", name2)

	name3, height := namedReturnValue()
	fmt.Printf("命名返回值: [name:%v height:%v]\n", name3, height)

	fmt.Printf("参数可变函数结果1: %v \n", sum(1))
	fmt.Printf("参数可变函数结果2: %v \n", sum(1, 2))
	fmt.Printf("参数可变函数结果3: %v \n", sum(1, 2, 3))

	//5.匿名函数
	func(name string) {
		fmt.Println(name)
	}("\n这个是匿名函数")

	//6.闭包函数
	addOne := addInt(1)
	fmt.Printf("闭包函数加1: %v \n", addOne())
	fmt.Printf("闭包函数加1: %v \n", addOne())
	fmt.Printf("闭包函数加1: %v \n\n", addOne())

	addTwo := addInt(2)
	fmt.Printf("闭包函数加2: %v \n", addTwo())
	fmt.Printf("闭包函数加2: %v \n", addTwo())
	fmt.Printf("闭包函数加2: %v \n", addTwo())

	//7.函数作为参数
	logger(sayHello, "lxw")

	//8.传值和传引用
	str:="北京"
	fmt.Println("开始调用sendValue,str : ",str) //开始调用sendValue,str :  北京
	sendValue(str)
	fmt.Println("结束调用sendValue,str : ",str) //结束调用sendValue,str :  北京

	fmt.Println("开始调用sendAddress,str : ",str) //开始调用sendAddress,str :  北京
	sendAddress(&str) //
	fmt.Println("结束调用sendAddress,str : ",str) //结束调用sendAddress,str :  lxw1234

}
