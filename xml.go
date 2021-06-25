package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Student struct {
	Name    string `xml:"name"`              //xml 标签
	Address string `xml:"address,omitempty"` //如果该字段为空就过滤掉
	Hobby   string `xml:"-"`                 // 进行xml 序列化的时候忽略该字段
	Father  string `xml:"parent>father"`     //xml 标签嵌套模式
	Mother  string `xml:"parent>mother"`     //xml 标签嵌套模式
	Note    string `xml:"note,attr"`         //xml 标签属性
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	stu1 := Student{
		Name:  "lxw",
		Hobby: "basketball",
		Note:  "userInfo",
	}
	//xml 序列化
	newData, err := xml.MarshalIndent(stu1, "", " ")
	checkErr(err)
	fmt.Println(string(newData))

	//写xml 文件
	err = ioutil.WriteFile("stu.xml", newData, 0777)
	checkErr(err)

	//读 xml 文件
	content, err := ioutil.ReadFile("stu.xml")
	stu2 := &Student{}
	fmt.Printf("stu2 指针:%#v \n",stu2) //空指针
	// stu2 指针:&main.Student{Name:"", Address:"", Hobby:"", Father:"", Mother:"", Note:""}


	//xml 反序列化
	err = xml.Unmarshal(content, stu2) //注意:第二个参数必须是指针
	checkErr(err)
	fmt.Printf("stu2: %#v\n", stu2)

}
