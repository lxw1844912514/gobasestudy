package main

import "fmt"

func swap(a ,b int)( int, int){
	return  b,a
}

//指针
func add(a *int)* int  {
	*a=*a+1
	return a
}

//for 循环
func getSum(num [6]int)int{
	sum:=0
	for i:=0;i<len(num) ;i++  {
		sum+=num[i]
	}
	return sum
}

func main()  {
	/*a:=1
	b:=3
	a,b=swap(a,b)
	fmt.Printf("%d %d",a,b)*/

	/*a:=1
	add(&a)
	fmt.Printf("%d",a)*/

	num:=[6]int{1,2,3,4,5}
	fmt.Print(getSum(num))
}
