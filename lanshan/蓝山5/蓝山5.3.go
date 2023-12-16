package main

//这个代码满足了多个线程应用锁保护同一个变量的基本要求
//采用了sync.WattGroup来控制各分任务
import (
	"fmt"
	"sync"
)

// CounterCapacity 柜台的容量
// breadCount 面包的数量
var CounterCapacity, breadCount = 5, 0
var l sync.Mutex

func Eat(wg *sync.WaitGroup) {
	defer wg.Done() //defer 表示延迟执行，等待Eat一切完毕再执行Done
	for {
		l.Lock()
		//柜台没面包
		if breadCount <= 0 {
			//不能占用面包师放面包
			l.Unlock()
			continue
		}
		//有面包，吃一个面包
		breadCount--
		fmt.Printf("吃")
		l.Unlock()
	}
}

func Make(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		l.Lock()
		if breadCount >= CounterCapacity {
			l.Unlock()
			continue
		}
		//放面包到柜台
		breadCount++
		fmt.Printf("做")
		l.Unlock()
	}

}

func Concurrent() {
	var wg sync.WaitGroup //线程计数器
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go Make(&wg)
		go Eat(&wg)
	}
	wg.Wait()
}
func main() {
	Concurrent()
}
