package main

import "fmt"

func main() {
	//1. make 初始化
	slice1 := make([]int, 5)     // slice1: [0 0 0 0 0]
	slice2 := make([]int, 6, 10) //初始化长度为5,容量为10的切片 slice2: [0 0 0 0 0 0]

	//2. 简短定义
	slice3 := []int{1, 2, 3, 4} // slice3: [1 2 3 4]
	//数组: arr1 = [5]int{11, 12, 13, 14, 15}
	//数组与切片区别: []内有数字的数组,没有数字的是切片

	//3.使用数组来初始化切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice4 := arr[0:3] //左闭右开区间 slice4: [1 2 3]

	//4.使用切片来初始化切片
	sliceA := []int{11, 12, 13, 14, 15, 16}
	sliceB := sliceA[0:3] //A:[11 12 13 14 15 16] B:[11 12 13]

	//5. 长度和容量
	fmt.Println("len:", len(sliceA))
	fmt.Println("cap:", cap(sliceA))

	//6. 改变切片长度
	sliceA = append(sliceA, 17)

	fmt.Println("append after:")
	fmt.Println("len:", len(sliceA))
	fmt.Println("cap:", cap(sliceA))
	/*
		len: 6
		cap: 6
		append after:
		len: 7
		cap: 12*///底层数组容量不够时，会重新分配数组空间，通常为两倍

	//7.对底层数组的修改，将影响上层多个切片的值
	/*newsliceA:=sliceA[0:3]
	fmt.Println("sliceA:",sliceA)
	fmt.Println("newsliceA:",newsliceA)

	newsliceA[0]=10
	fmt.Println("after modifying:")
	fmt.Println("sliceA:",sliceA)
	fmt.Println("newsliceA:",newsliceA)*/

	/* 结果:
	sliceA: [11 12 13 14 15 16 17]
	newsliceA: [11 12 13]
	after modifying
	sliceA: [10 12 13 14 15 16 17]
	newsliceA: [10 12 13]
	*/

	//8.使用 copy 方法可以避免共享同一个底层数组
	newsliceA2 := make([]int, len(sliceA))
	copy(newsliceA2, sliceA)
	fmt.Println("sliceA: ", sliceA)
	fmt.Println("newsliceA2: ", newsliceA2)

	newsliceA2[0] = 111
	fmt.Println("after modifying: ")
	fmt.Println("sliceA: ", sliceA)
	fmt.Println("newsliceA2: ", newsliceA2)
	/*结果:
	    sliceA:  [11 12 13 14 15 16 17]
		newsliceA2:  [11 12 13 14 15 16 17]
		after modifying:
		sliceA:  [11 12 13 14 15 16 17]
		newsliceA2:  [111 12 13 14 15 16 17]
	*/

	//9.使用copy函数进行切片部分拷贝
	newsliceA3 := make([]int, len(sliceA))
	copy(newsliceA3, sliceA[3:]) //newsliceA3: [14 15 16 17 0 0 0]

	fmt.Println("newsliceA3:", newsliceA3)

	fmt.Printf("%v %v %v %v %v %v \n", slice1, slice2, slice3, slice4, sliceA, sliceB)

}
