package main

import "fmt"

//自定义类型
//声明: type City string

type City string
type Age int
type Height int

func main() {
	city := City("beijing")
	age := Age(12)
	height := Height(175)

	fmt.Println(" I live in", city+"上海")
	fmt.Println("city len:", len(city))
	if age > 14 {
		fmt.Println("age is bigger than 14")
	} else {
		fmt.Println("age is smaller than  14")
	}

	//因为 printAge 方法期望的是 int 类型，但是我们传入的参数是 Age，他们虽然具有相同的值，但为不同的类型
	printAge(int(age)) //显式的类型转换
	fmt.Println("除法计算: ", int(height)/int(age))
}

//函数参数
func printAge(ageNum int) {
	fmt.Println("Age is:", ageNum)
}
