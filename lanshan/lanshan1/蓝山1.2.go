package main

import "fmt"
import "math"

func main() {
	var r, s float64
	fmt.Print("输入圆的半径：")
	fmt.Scan(&r)
	s = math.Pi * math.Pow(r, 2)
	fmt.Println("圆的面积是：", s)
}
