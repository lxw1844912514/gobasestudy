package main

import "fmt"

//分支循环: if switch for 进行条件判断和流程控制

func main() {
	age := 82

	//1. if分支
	if age > 6 && age <= 12 {
		fmt.Println("it's primary school")
	} else if age > 12 && age <= 15 {
		fmt.Println("it is middle school")
	} else if age < 20 {
		fmt.Println(" it is high school")
	} else {
		fmt.Println("年龄超限")
	}

	//2. switch 分支
	switch age {
	case 5:
		fmt.Printf("his age is %v \n", age)
	case 6, 7, 8: //case 条件可以是多个值,同一个 case 中的多值不能重复
		fmt.Printf("his age is %v \n", age)
	default:
		fmt.Println("The age is unknown ")
	}

	//2.1 还可以使用if ...else 作为case 条件
	switch {
	case age >= 5 && age <= 12:
		fmt.Println("小学生")
	case age > 12 && age <= 15:
		fmt.Println("中学生")
	case age > 40 || age < 5:
		fmt.Println("不是学生")
	default:
		fmt.Println("age is unknown")
	}

	//2.2 调用
	checkType(8)
	checkType("hello world")

	//3.for 循环
	for i := 0; i < 2; i++ {
		fmt.Println("loop with index", i)
	}

	//3.1 使用for..range 对数组,切片 map 字符串等进行循环操作
	numbers := [5]int{1, 2, 3}
	for i, v := range numbers { //这里的i v 是切片元素的位置索引和值
		fmt.Printf("numbers[%d] is %d \n", i, v)
	}

	//3.2 循环 map
	cityCode := map[string]int{
		"北京": 1,
		"上海": 2,
		"广州": 3,
	}
	for i, v := range cityCode { //i  v 是map 键值对的键和值
		fmt.Printf("%s is %d \n", i, v)
	}

	//3.3 使用break 和continue 对循环进行控制
	//break 会结束所有循环
	//continue 会跳过当前循环直接进入下一次循环
	num := []int{1, 2, 3, 4, 5}
	for i, v := range num {
		if v == 6 {
			break
		}

		if v%2 == 0 {
			continue
		}

		fmt.Printf("num[%d] is %d \n", i, v)
	}

}

//2.2使用switch 对interface 进行断言
func checkType(i interface{}) {
	switch v := i.(type) { //.(type)只能在switch 中使用
	case int:
		//fmt.Printf("%v is a int\n", i.(type))   //报错: use of .(type) outside type switch
		fmt.Printf("%v is a int\n", v)
	case string:
		fmt.Printf("%v is a string\n", v)
	default:
		fmt.Printf("%v's type is unkown\n", v)
	}
}
