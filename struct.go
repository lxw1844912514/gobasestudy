package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person //匿名字段,默认包含person中所有字段
	sex    string
}

func main() {
	//person:=Person{"zhangsan",25} //{zhangsan 25}
	//person := Person{age: 22, name: "lxw"}

	//stu1 := Student{Person{name: "test", age: 11},  "男"}
	stu1 := Student{Person{"test", 113},  "男"}

	//fmt.Printf("%v", person)
	fmt.Printf("%v", stu1.Person.age)//113
	fmt.Printf("%v", stu1.name)//test
	fmt.Printf("%v", stu1.age)//113
	fmt.Printf("%v", stu1.sex)//男
}
/**

 */

//go 语言没有private,public,protected 这样的关键字
//要使某个符号对其他包package 可见(即可以访问),需要将该符号定义为以大写字母开头