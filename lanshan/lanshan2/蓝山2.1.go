package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Calculator(num1 int, num2 int, f func(int, int) int) int {
	result := f(num1, num2)
	return result
}
func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumbera := rand.Intn(101)
	randomNumberb := rand.Intn(101)
	fmt.Println(randomNumbera)
	fmt.Println(randomNumberb)
	fmt.Println(Calculator(randomNumbera, randomNumberb, func(a, b int) int { return a - b }))

}
