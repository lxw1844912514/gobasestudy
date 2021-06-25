package main

import "fmt"

//结构体是由零个或多个任意类型的值聚合成的实体
//数组、切片和 Map 可以用来表示同一种数据类型的集合，但是当我们要表示不同数据类型的集合时就需要用到结构体。
//注意:
//1.字段名是唯一的
//2.结构体中字段占一行,但类型相同的字段,可以放在一行
//3.结构体中的字段如果是小写字母开头,那么其他package就无法使用该字段


//使用 type,struct 关键字定义结构体

type Student struct {
	Age  int
	Name string
	sex string
}

//结构体嵌套结构体,如果嵌入的结构体是本身,那么只能是指针
type Tree struct {
	value int
	left,right *Tree
}

func main() {
	stu := Student{
		Age:  11,
		Name: "lxw",
	}
	fmt.Println("学生1:", stu) //学生: {11 lxw}

	//在赋值的时候,字段名可以忽略
	fmt.Println("赋值: ",Student{20, "lxw22",""})

	//嵌套结构体
	tree := Tree{
		value: 1,
		left: &Tree{
			value: 1,
			left:  nil,
			right: nil,
		},
		right: &Tree{
			value: 2,
			left:  nil,
			right: nil,
		},
	}

	fmt.Printf(">>> %#v\n", tree)

	//结构体比较,前提是字段类型是可以比较的
	tree1:=Tree{
		value: 222,
	}
	tree2:=Tree{
		value: 333,
	}
	fmt.Printf(">>> %v\n",tree1==tree2)


	//结构体内嵌匿名成员
	//匿名成员:声明一个成员对应的数据类型而不指明成员的名字
	

	return
}
