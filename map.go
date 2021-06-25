package main

import (
	"fmt"
	"sync"
)

func main() {
	//1. 使用var 声明
	var cMap map[string]int
	//cMap["beijing"] = 1
	fmt.Println(cMap == nil)    //true
	fmt.Println("cMap: ", cMap) //cMap:  map[]

	//线程不安全, 一个 goroutine 在对 map 进行写的时候，另外的 goroutine 不能进行读和写操作
	//panic: assignment to entry in nil map
	// goroutine 1 [running]:

	//2. 使用make 声明
	//在使用make初始化map的时候,可以指定初始容量,这在能预估map key数量的情况下,减少动态分配的次数,从而提升性能
	cMap2 := make(map[string]int, 10)
	cMap2["beijing"] = 2
	fmt.Println("cMap2: ", cMap2) //cMap2:  map[beijing:2]

	//3.简短声明
	cMap4 := map[string]int{"beijing": 1}
	fmt.Println("cMap4: ", cMap4) //cMap4:  map[beijing:1]

	//4.map 基本操作
	cMap5 := map[string]int{}
	cMap5["beijing"] = 1    //写
	cMap5["guangzhou"] = 12 //写

	code := cMap5["beijing"]    //读
	fmt.Println("code: ", code) //code:  1

	code2 := cMap5["guangzhou"]  //读不存在的key
	fmt.Println("code2:", code2) //code2: 0

	code3, ok := cMap5["guangzhou"] //检查key 是否存在
	if ok {
		fmt.Println("code3: ", code3)
	} else {
		fmt.Println("key not exist")
	}

	delete(cMap5, "beijing")      //删除key
	fmt.Println(cMap5["beijing"]) //0

	//5.循环和无序性
	cMap6 := map[string]int{"beijing": 1, "shanghai": 2, "shengzhen": 3}
	for city, code6 := range cMap6 {
		fmt.Printf("%s : %d ", city, code6)
		fmt.Println()
	}
	//beijing 1
	//shanghai 2
	//shengzhen 3

	//6.线程不安全
	cMap7 := make(map[string]int)
	var wg sync.WaitGroup
	//var wg sync.Map
	wg.Add(2)

	go func() {
		cMap7["beijing"] = 1
		wg.Done()
	}()

	go func() {
		cMap7["shanghai"] = 2
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("cMap7: %v \n", cMap7)//cMap7: map[beijing:1 shanghai:2]

	//7.map 嵌套
	provices := make(map[string]map[string]int)

	provices["beijing"] = map[string]int{
		"东城区": 11,
		"西城区": 12,
		"朝阳区": 13,
		"海淀区": 14,
	}

	fmt.Println(provices["beijing"])//map[东城区:11 朝阳区:13 海淀区:14 西城区tl:12]

}
