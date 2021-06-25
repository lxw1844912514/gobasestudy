package main

import "fmt"

/*func compare(a, b int) bool {
	return a < b
}*/

type Integer struct {
	value int
}

func (a Integer) compare(b Integer)bool{
	return a.value<b.value
}

func main() {
	var a int = 1
	var b int = 2
fmt.Printf("%v",compare(a,b))
}
