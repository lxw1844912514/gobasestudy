package main

import (
	"fmt"
)

func main() {
	//1. 变量
	var (
		x, y int //同时声明 x,y 为整数
		z    float64
	)
	//
	d := 33       //简短声明变量
	e := int64(2) //声明e为 2 的64位整数

	//2. 常量
	const a = 64
	const (
		b = 3
		c = 0.1
	)

	//3.数据类型 布尔类型
	var bb = true
	cc := false

	//数字类型:分为有符号数,无符号数,有符号数可以表示负数,不同位数代表它们实际存储占用空间,以及取值范围
	var (
		ii uint8 = 1
		pp int8  = 18 // 超出范围报错 : constant 128 overflows int8
		oo int   = 64 // int uint,由操作系统决定,在32位系统中他们的长度为32位,在64位系统中,长度为64位
	)

	//字符串
	var str = "this a 测试"
	var oldstr = "\""                    //原样输出
	var line = ` 
	line1
	line2
	` //多行

	var str1 = "hello,世界"
	substr := str1[0]
	substr1 := str1[1]
	substr2 := str1[2]
	substr3 := str1[3]
	substr4 := str1[4]

	fmt.Printf("%v %v %v %v %v ", substr, substr1, substr2, substr3, substr4) // 104 101 108 108 111

	fmt.Println("\n\r")
	fmt.Printf("%v %v %v %v %v", string(substr), string(substr1), string(substr2), string(substr3), string(substr4)) //h e l l o
	//fmt.Println(substr) //104
	fmt.Println(len(str))         // 13 字符串有多少个字节  汉字算3个字节,空格算一个
	fmt.Println(len([]rune(str))) // 9 查看有多少个字符

	fmt.Println("=======\r")
	//byte: uint8 别名,表示二进制数据的byte
	//rune :int32 别名,用于表示一个符号
	for _,char:=range str{
		fmt.Printf(" %T ",char)
	}
	fmt.Println("\r\n=======")

	fmt.Printf("%v %v %v ", str, oldstr, line)

	fmt.Printf("%v %v %v ", ii, pp, oo) //1 1 64

	fmt.Println("\n\r")
	fmt.Printf("a= %v b= %v c= %v  %v %v %v %v %v %v %v ", a, b, c, x, y, z, d, e, bb, cc)

	fmt.Println("hello word")
}
