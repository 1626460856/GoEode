package main

//这个代码满足了多个线程应用锁保护同一个变量的基本要求
//使用了条件变量解决线程等待和唤醒问题
//{在 Eat2() 函数中，当柜台没有面包时，
//线程会调用 cond.Wait() 进入等待状态，直到被唤醒。
//当有面包可供吃时，线程会减少 breadCount2 的数量，
//然后调用 cond.Broadcast() 唤醒其他线程，让它们有机会执行。}
import (
	"fmt"
	"sync"
)

// CounterCapacity 柜台的容量
// breadCount 面包的数量
var CounterCapacity2, breadCount2 = 5, 0
var l2 sync.Mutex
var cond = sync.NewCond(&l2)

func Eat2() {
	for {
		l2.Lock()
		//柜台没面包
		for breadCount2 <= 0 {
			//不能占用面包师放面包
			cond.Wait()
		}
		//有面包，吃一个面包
		breadCount2--
		fmt.Printf("吃")
		cond.Broadcast()
		l2.Unlock()
	}
}

func Make2() {

	for {
		l2.Lock()
		for breadCount2 >= CounterCapacity2 {
			cond.Wait()
		}
		//放面包到柜台
		breadCount2++
		fmt.Printf("做")
		cond.Broadcast()
		l2.Unlock()
	}

}

func Concurrent2() {

	for i := 0; i < 3; i++ {
		go Eat2()
		go Make2()

	}
	select {}
}
func main() {
	Concurrent2()
}
