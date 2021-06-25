package main

import (
	"fmt"
	"math"
)

//接口类型是一种抽象类型,是方法的集合,其他类型实现了这些方法就是实现了接口

//打印矩形和圆的面积和周长

//定义接口
type tmp interface {
	area() float32      //面积
}
type geometry interface {
	//area() float32      //面积
	tmp  //接口中内嵌接口
	perimeter() float32 //周长
}

//定义结构体: 矩形
type rect struct {
	length, width float32
}

func (r rect) area() float32 {
	return r.length * r.width
}

func (r rect) perimeter() float32 {
	return 2 * (r.length + r.width)
}

//定义结构体:圆形

type circle struct {
	radius float32 //半径
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

//展示调用
func show(name string, param interface{}) {
	//fmt.Println(param)
	switch param.(type) {
	case geometry:
		//类型断言
		fmt.Printf(" %v的面积是 %v \n", name, param.(geometry).area())
		fmt.Printf(" %v的周长是 %v \n", name, param.(geometry).perimeter())
	default:
		fmt.Println("wrong type")
	}
}

func main() {

	rec := rect{
		length: 1,
		width:  2,
	}
	show("矩形", rec)

	cir := circle{radius: 1}
	show("圆形", cir)

	show("测试:", "test param")
}
