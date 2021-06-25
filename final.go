package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "hello,世界"
	//
	//for _, char := range str {
	//	fmt.Printf("%v is %T \n",string(char), char)
	//}

	var a int8 = 127
	fmt.Println(a)

	fmt.Println(len(str))
	fmt.Println(len([]int32(str))) //  ===>	fmt.Println(len([]rune(str)))

	arr4 := [...]int{3: 0}
	fmt.Println(arr4)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	arr5 := []int64{2, 3, 4, 5,6}
	fmt.Println(arr5)
	fmt.Printf("arr5 类型: %T ;",arr5) //arr5 类型: []int64 ;
	fmt.Println(reflect.TypeOf(arr5)) //[]int64

	slice := make([]int64, 2)
	copy(slice, arr5[2:4])
	fmt.Println(slice)

	//var cMap map[string]int  // 只定义, 此时 cMap 为 nil
	//fmt.Println(cMap == nil)
	//cMap["北京"] = 1
	//cMap1:=make(map[string]int,10)
	//cMap1["bj"]=1
	//
	//for k,i:=range str {
	//	cMap1[string(i)]=k
	//}
	//
	//fmt.Println(cMap1)
	cMap := map[string]int{"北京": 1, "上海": 2, "广州": 3, "深圳": 4}

	for city, code := range cMap {
		fmt.Printf("%s:%d \n", city, code)
	}


	//map 嵌套
	provinces := make(map[string]map[string]int)
/*	provinces["北京"] = map[string]int{
		"东城区": 1,
		"西城区": 2,
		"朝阳区": 3,
		"海淀区": 4,
	}
	provinces["河北"] = map[string]int{
		"邯郸": 11,
		"邢台": 21,
		"衡水": 31,
		"石家庄": 14,
	}*/
	provinces=map[string]map[string]int{
		"北京":{
			"东城区": 1,
			"西城区": 2,
			"朝阳区": 3,
			"海淀区": 4,
		},
		"河北":{
			"邯郸": 11,
			"邢台": 21,
			"衡水": 31,
			"石家庄": 14,
		},
	}

	fmt.Println(provinces["北京"])
}
