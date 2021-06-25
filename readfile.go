package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//1.全量读
/*func main(){
	data,err:=ioutil.ReadFile("./lock.go")
	fmt.Println(err)
	fmt.Println(string(data))
}*/

//2.带缓冲区读: 输出文件的前 16个字节内容
/*func main()  {
	f,_:=os.Open("./lock.go")
	defer f.Close()

	buf:=make([]byte,16)
	f.Read(buf)
	fmt.Println(string(buf))
}*/

/*
package main
im
*/

// 3.任意位置读:一个文件特定地方特定长度的内容

func main() {
	f, _ := os.Open("./lock.go")
	defer f.Close()

	//第一种: f.Seek+f.Read 非并发安全的
	//File.Seek(offset, whence)，设置光标的未知
	//offset步长值,偏移量
	//whence，从哪开始：0距离文件开头，1当前位置，2距离文件末尾
	/*for i := 0; i < 5; i++ {
		go func() {
			b1 := make([]byte, 2) //5 读取几个字节

			//f.Seek(5, 0)
			//f.Read(b1)

			f.ReadAt(b1,5)
			fmt.Println(string(b1))

			//f.Seek(2, 0)
			//f.Read(b1)

			f.ReadAt(b1,2)
			fmt.Println(string(b1))
		}()
	}
	time.Sleep(time.Second)*/

	//第二种 使用f.ReadAt 并发安全的
	//f.ReadAt(b1,5) //5 步长值,从哪个位置开始读取文件

	//实战一: 使用buf实现ioutil.ReadFile 效果
	/*content := make([]byte, 0)
	buf := make([]byte, 16)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		if n == 16 {
			content = append(content, buf...)
		} else {
			content = append(content, buf[0:n]...)
		}
	}
	fmt.Println(string(content))*/

	//实战二:使用 bufio 实现行统计
	br := bufio.NewReader(f)
	totalLine := 0
	for {
		_, isPrefix, err := br.ReadLine()
		if !isPrefix {
			totalLine += 1
		}

		if err == io.EOF {
			break
		}
		//fmt.Println(line) //打印字母的 Unicode 字符编码

	}
	fmt.Println("total line is", totalLine)
}
