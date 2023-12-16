package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Print("输入第一个值：")
	fmt.Scan(&x)
	fmt.Print("输入第二个值：")
	fmt.Scan(&y)
	z = x + y
	fmt.Println(z)
}
