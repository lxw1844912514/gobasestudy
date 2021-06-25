package main

import "fmt"

func main() {
	//1 定义长度为5的数组
	var arr1 [5]int

	for i := 0; i < 5; i++ { //i<3 可以,i<10 报错
		arr1[i] = i
	}

	//printHelper("arr1", arr1)
	arr1 = [5]int{11, 12, 13, 14, 15}//// 长度和元素类型都相同，可以正确赋值

	//2. 简写模式,在定义的同时给出赋值
	arr2 := [5]int{22, 23, 24, 26}
	//printHelper("arr2", arr2)

	fmt.Println(arr2==arr1)

	//3. 不显示定义数组长度,由编译器完成数组长度的计算
	var arr3  =[...]int{33,34,35,36,37}
	printHelper("arr3",arr3)

	//4. 定义前三个元素为默认值0,第四个为2,最后一个元素为-1
	var arr4=[...]int{3: 2, -1}
	printHelper("arr4",arr4)

	//5.多维数组
	var twoArr [2][3]int
	for i:=0;i<2 ;i++  {
		for j:=0;j<3 ;j++  {
			twoArr[i][j]=i+j
		}
	}
	fmt.Println("twoArr:",twoArr)
}

func printHelper(name string, arr [5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v[%v]: %v\n", name, i, arr[i])
	}

	//len 获取长度
	fmt.Printf("len of %v: %v\n", name, len(arr))

	//cap 也可以用来获取数组长度
	fmt.Printf("cap of %v: %v\n", name, cap(arr))

	fmt.Println("arr: " ,arr)
	slice := arr[0:5]
	fmt.Println("slice:",slice)
}
