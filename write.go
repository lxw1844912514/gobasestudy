package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//创建文件
//使用os.Create(name string)方法创建文件
//使用os.Stat(name string) 方法获取文件信息
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)
	fmt.Println("file content:", string(data))
}

/*func main() {
	path := "test.txt"
	newFile, err := os.Create(path)
	checkErr(err)
	defer newFile.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
			return
		}
	}
	fmt.Println("文件存在,文件名:", fileInfo.Name())
	fmt.Println("文件存在,文件大小:", fileInfo.Size())
	fmt.Println("文件存在,文件权限:", fileInfo.Mode())
}*/

//2. ioutil.WriteFile 一步完成文件创建和写入操作,如果文件存在,则覆盖原来内容

/*func main(){
	path:="test.txt"
	str:="hello world"

	err:=ioutil.WriteFile(path,[]byte(str),0644)
	checkErr(err)
	readFile(path)
}
*/

//3.使用 writeAt 在文件指定位置写入内容
/*func main() {
	path := "test.txt"
	str := "hello world"
	newStr := "123456 "

	//创建文件
	newFile, err := os.Create(path)
	checkErr(err)

	n1, err := newFile.WriteString(str)
	checkErr(err)
	fmt.Println("n1: ", n1)
	readFile(path)

	n2,err:=newFile.WriteAt([]byte(newStr),6) //返回写入的字节数和错误(如果有)
	checkErr(err)
	fmt.Println("n2: ",n2)
	readFile(path)

	n3,err:=newFile.WriteAt([]byte(newStr),33)//从0下标位置开始写入
	checkErr(err)
	fmt.Println("n3: ",n3)
	readFile(path)

//结果:
//n1:  11
//	file content: hello world
//n2:  7
//	file content: hello 123456
//n3:  7
//	file content: hello 123456 123456

}*/

//4.使用 Buffered Writer 可以避免太多次的磁盘io操作,写入的内容首先是存在内存中,当调用Flush()方法才会写入磁盘

func main() {
	path := "test.txt"
	str := "hello world"

	//创建文件
	newFile, err := os.Create(path)
	checkErr(err)
	defer newFile.Close()

	bufferWrite := bufio.NewWriter(newFile)

	for _, v := range str {
		written, err := bufferWrite.WriteString(string(v))
		checkErr(err)
		fmt.Println("written: ", written) //循环写入11次,每次1个字节 ; written:  1
	}
	readFile(path) //没有flush 之前内容为空,file content:

	unflushSize := bufferWrite.Buffered()
	fmt.Println("unflushSize: ", unflushSize) //unflushSize:  11

	bufferWrite.Flush()
	readFile(path) //file content: hello world
	//结果:
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//written:  1
//	file content:
//unflushSize:  11
//	file content: hello world

}
