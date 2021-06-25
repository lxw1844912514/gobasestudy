package main

import "fmt"

//在面向对象语言中,用"类"来封装属于自己的数据和函数,这些类的函数就叫做方法.

//方法(method)的声明和函数很相似,只不过它必须指定接受者
//声明 func(t T)F(){}
//注意:
//接受者的类型只能为用关键字type 定义的类型,例如自定义类型,结构体
//同一个接受者的方法名不能重复(没有重载),如果是结构体,方法名还不能和字段名重复
//值作为接受者无法修改其值,如果有更改需求,需要使用指针类型

//type T int64
//func (t T) F()  {} //T 接受者不是任意类型,它只能为用关键字type定义的类型(例如自定义类型,结构体)

//type T struct{}

//结构体方法名不能和字段重复
/*type T struct {
	F string
}
func (T)  F(){}
//func(T) F(a string){}

func main() {
	//t := T(10)
	t:=T{}
	t.F()
}*/

//2.接收者可以同时为值和指针

/*func (T) F(){}
func (*T)N()  {}

func main()  {
	t:=T{}
	t.F()
	t.N()

	t1:=&T{} //无论值类型T 还是指针类型 &T 都可以同时访问F和N方法
	t1.F()
	t1.N()
}*/

//3.值和指针作为接受者的区别
type T struct {
	value int
}

func (m T) stayTheSame() {
	m.value = 3
}

func (m *T) Update(){
	m.value=3
}

func main()  {
	m:=T{0}
	fmt.Println(m)//{0}

	m.stayTheSame()
	fmt.Println(m)//{0}

	m.Update()
	fmt.Println(m)//{3}
}
