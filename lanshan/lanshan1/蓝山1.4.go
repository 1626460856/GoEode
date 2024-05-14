package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成[0,100)范围内的随机整数
	//fmt.Println("0-98之间的随机整数：", rand.Intn(98))

	// 生成[0.0,1.0)范围内的随机浮点数
	//fmt.Println("0-1之间的随机浮点数：", rand.Float64())
	var a float64
	var d int
	a = 1 + float64(rand.Intn(98)) + rand.Float64()
	fmt.Println("1-100随机一个数：", a)
	fmt.Print("输入二分法的精度：")
	fmt.Scan(&d)
	var b float64 = 1
	var c float64 = 100
	for i := 1; i <= d; i++ {
		if (b+c)/2 < a {
			b = (b + c) / 2
		}
		if (b+c)/2 > a {
			c = (b + c) / 2
		}

	}
	fmt.Println("随机数的范围处于(", b, ",", c, ")")
}
