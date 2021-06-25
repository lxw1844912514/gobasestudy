package main

import (
	"encoding/json"
	"io/ioutil"
)

//在http,rpc 的微服务调用中,以 json方式对数据进行序列化和反序列化

/*func main() {
	var (
		data  = `23`
		value int64
	)
	err1 := json.Unmarshal([]byte(data), &value)
	fmt.Println("Unmarshal error is:", err1)
	fmt.Printf("Unmarshal value is :%T,%v \n", value, value)

	value2, err2 := json.Marshal(value)

	fmt.Println("Marshal error is: ", err2)
	fmt.Printf("Marshal value is: %s \n", string(value2))
	//注意：在实际应用中，我们在序列化和反序列化的时候，需要检查函数返回的 err, 如果 err 不为空，表示数据转化失败。
}*/

/*func printHelper(name string, value interface{}) {
	fmt.Printf("%s Unmarshal value is: %T ,%v \n", name, value, value)
}
func main() {
	var (
		d1 = `false`
		v1 bool
	)

	json.Unmarshal([]byte(d1), &v1)
	printHelper("d1", v1)

	var (
		d2 = `2`
		v2 int
	)
	json.Unmarshal([]byte(d2), &v2)
	printHelper("d2", v2)

	var (
		d3 = `3.14`
		v3 float32
	)
	json.Unmarshal([]byte(d3), &v3)
	printHelper("d3", v3)

	var (
		d4 = `[1,2]`
		v4 []int
	)
	json.Unmarshal([]byte(d4), &v4)
	printHelper("d4", v4)

	var (
		d5 = `{"a","b"}`
		v5 map[string]string
		v6 interface{}
	)
	json.Unmarshal([]byte(d5), &v5)
	printHelper("d5", v5)

	json.Unmarshal([]byte(d5),&v6)
	printHelper("d6(interface{})",v6)

}*/
type Result struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type Student struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Result []Result `json:"result"`
}

//反序列化
/*func main() {
	dat, _ := ioutil.ReadFile("data.json")
	var s Student
	json.Unmarshal(dat, &s)
	fmt.Printf("Student's result is:%v \n", s) //Student's result is:{1 小红 [{语文 90} {数学 100}]}
}*/

//序列化

func main() {
	s := Student{
		Id:   1,
		Name: "小明",
		Result: []Result{
			Result{
				Name:  "语文",
				Score: 90,
			},
			Result{
				Name:  "数学",
				Score: 98,
			},
		},
	}
	dat, _ := json.Marshal(s)
	ioutil.WriteFile("data2.json", dat, 0755)
}
