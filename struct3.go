package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

type Student struct {
	Person
}

//声明一个成员的数据类型而不指名成员的名字,这类成员就叫匿名成员

func main() {
	per := Person{
		Age:  18,
		Name: "lxw",
	}
	per2 := Person{
		Age:  20,
		Name: "hello",
	}

	stu := Student{Person: per}
	stu2 := Student{Person: per2}

	fmt.Println("stu.Age: ", stu.Age)
	fmt.Println("stu2.Name: ", stu2.Name)
	//stu.Age:  18
	//stu.Name:  hello
}
