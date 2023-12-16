package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 创建了一个计算方差的计算方法
func FangCha(a ...float64) float64 {
	var b float64
	for _, k := range a {
		b += k
	}
	var c, d float64
	PingJunShu := b / float64(len(a))
	for _, k := range a {
		c = (k - PingJunShu) * (k - PingJunShu)
		d += c
	}
	FangCha := d / float64((len(a)))
	return FangCha
}
func main() {
	fmt.Println("你可以输入任意个你想输入的数据，输入一次按一次回车，终止输入时按两次回车，可以计算出你输入的所有数据的方差")
	scanner := bufio.NewScanner(os.Stdin) //创建一个新的扫描器（Scanner），用于从标准输入（os.Stdin）读取输入
	var a []float64
	for scanner.Scan() { //在每次迭代时检查扫描器是否还有更多的数据可供读取。如果扫描器可以继续读取数据，则进入循环体内执行
		inputStr := scanner.Text() //将当前扫描器所在位置的文本内容存储到 inputStr 变量中。
		// scanner.Text() 方法会返回最后一次调用 Scan() 方法后的扫描器的文本内容。
		if inputStr == "" {
			break
		}
		inputArr := strings.Split(inputStr, " ") //将字符串 inputStr 按照空格进行拆分，
		// 并将拆分后的子字符串存储在 inputArr 切片中
		for _, x := range inputArr {
			num, err := strconv.ParseFloat(x, 64)
			if err == nil {
				// 转换成功，将 num 添加到 float64 切片中
				a = append(a, num)
			} else {
			} //这段代码将输入的字符串切片按照空格进行拆分，并将拆分后的每个子字符串 x 转换为 float64 类型
			// 并将转换后的结果添加到 a 切片中
		}
	}
	y := FangCha(a...)
	fmt.Println("你输入数据的方差为：", y)
}
