package main

import (
	"fmt"
	"os"
	"time"
)

// 定义的Sum函数通过计算处理变量和os操作的时间差，
// 验证了os操作要慢于变量操作
func Sum(n int) {
	f, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() //创立并延时关闭f文件
	a := 0
	t := time.Now() //计时t开始
	for i := 0; i < n; i++ {
		a++
	}

	fmt.Println("操作变量a的时间为：", time.Since(t)) //打印自从时间点t以来经过的时间
	t = time.Now()
	//下面的_, err = f.WriteString("a")也可以分开写出来
	if _, err = f.WriteString("a"); err != nil { //将a写到了text.txt文件中
		panic(err)
	}
	fmt.Println("写入变量a的时间为：", time.Since(t))

}

// 开启n个线程，每个都打印100次
func ConcurrentPrint(n int) {
	for i := 0; i < n; i++ {
		go Print()
	}
}
func Print() {
	for i := 0; i < 100; i++ {
		fmt.Println("a")
	}
}
func thistime() time.Time {
	t := time.Now()
	return t
}
func sub(t1 time.Time, t2 time.Time) time.Duration {
	subt := t1.Sub(t2)
	return subt
}
func bijiaoshijian() {
	t1 := thistime()
	ConcurrentPrint(10)
	t2 := thistime()

	t3 := thistime()
	for i := 0; i < 10; i++ {
		for x := 0; x < 100; x++ {
			fmt.Println("a")
		}
	}
	t4 := thistime()
	fmt.Println(sub(t1, t2))
	fmt.Println(sub(t3, t4))
}

// n个线程轮流打印1到10
func ConcurrentPrint2() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%d ", i)
		}()
	}
	time.Sleep(100000000)
	//这个并发函数与主函数同等，需要等待不然来不及打印
}

// 累计全局变量a自增
var a = 0

func Adda() {
	go Add()
	go Add()
	time.Sleep(100000000)
}
func Add() {
	for i := 0; i < 10000; i++ {
		a++
	}
}
func cheshiAdda() {
	Adda()
	fmt.Println(a)

}

// 原子操作
func compare_and_swap(addr *int, old, new int) (swapped bool) {
	oldVal := *addr
	if oldVal == old {
		*addr = new
		return true
	}
	return false
}

func main() {
	cheshiAdda()

}
