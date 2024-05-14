package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	// 打开文件（使用 OpenFile）
	file, err := os.OpenFile("wenzi2.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer file.Close()
		fmt.Println("成功打开文件:", file.Name())
	}

	// 重新定位文件指针到文件开头
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("重新定位文件指针出错:", err)
		return
	}

	// 读取文件
	time1 := time.Now()

	for i := 1; i <= 1000000; i++ {
		// 写入文件
		dataToWrite := []byte("123456789123456789123456789")
		bytesWritten, err := file.Write(dataToWrite)
		if err != nil {
			fmt.Println("写入文件出错:", err)
			return
		}
		fmt.Printf("成功写入 %d 字节到文件\n", bytesWritten)
		//重新定位文件指针到文件开头
		_, err = file.Seek(27, 27*i)
		if err != nil {
			fmt.Println("重新定位文件指针出错:", err)
			return
		}
		//读取文件
		readBuffer := make([]byte, 27)
		bytesRead, err := file.Read(readBuffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("读取文件出错:", err)
				return
			}
		}

		fmt.Printf("从文件中读取 %d 字节的数据: %s\n", bytesRead, readBuffer[:bytesRead])
	}
	time2 := time.Now()

	myfile, err := os.OpenFile("wenzi.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer myfile.Close() //关闭文件
	//初始化缓存区与文件关联
	writer := bufio.NewWriter(myfile)

	//初始阅读器
	myfile, _ = os.Open("wenzi.txt")
	defer myfile.Close()
	reader := bufio.NewReader(myfile)
	time3 := time.Now()
	//写入字符切片到缓存区
	for i := 1; i <= 1000000; i++ {
		data := []byte("123456789123456789123456789")
		_, err = writer.Write(data)
		if err != nil {
			fmt.Println("写入数据到缓冲区出错:", err)
			return
		}
		read_data := make([]byte, 27)
		n, err := reader.Read(read_data)
		if err != nil {
			if err != io.EOF {
				fmt.Println("从缓冲区读取数据出错:", err)
				return
			}
		}
		fmt.Printf("从缓冲区读取 %d 字节的数据: %s\n", n, read_data)

	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区出错:", err)
		return
	}
	fmt.Println("刷新缓冲区成功")
	time4 := time.Now()
	shijiancha1 := time2.Sub(time1) //运算time2-time1
	fmt.Println("无缓冲写入读取1000000次耗时:", shijiancha1)
	shijiancha2 := time4.Sub(time3) //运算time4-time3
	fmt.Println("有缓冲写入读取1000000次耗时:", shijiancha2)

	//
}
