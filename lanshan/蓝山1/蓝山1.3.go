package main

import "fmt"

func main() {
	var a int
	fmt.Print("请输入一个数，我们将会判断这个数是否为质数：")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		if (a%i) == 0 && i == a {
			fmt.Println(a, "是质数")
			break
		}
		if (a%i) == 0 && i < a && i != 1 {
			fmt.Println(a, "不是质数", "该数能被", i, "整数")
			break
		}

	}
}
