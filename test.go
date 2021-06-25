package main

import "fmt"

func main() {
	//变量
	var x int
	x = 2
	fmt.Printf("%d", x) //2

	var y [10]int //数组
	y[1] = 1
	y[9] = 10
	fmt.Print(y) //[0 1 0 0 0 0 0 0 0 10]

	z := [102] int64{1, 3, 4}
	fmt.Printf("%v", len(z)) //%v 变量 10

	s := make([]int, 3, 3) //make([]类型,len长度,cap容量)
	s = append(s, 3, 4, 5, 6, 7)
	fmt.Print("___")
	fmt.Printf("%v", cap(s))
	fmt.Print("___")

	var student map[string]float32
	student=make(map[string]float32) //make 用于内建类型	(map,channel,slice)的内存分配
	student["zs"]=19.2

	fmt.Printf("%v",student)
}
