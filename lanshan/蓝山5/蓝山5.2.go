package main

import (
	"fmt"
	"sync"
)

//使用官方锁Mutex
//var m sync.Mutex
//for{
//m.Lock()
/*临界区*/
//m.Unlock()
//}
var count int     //全局计数器
var mu sync.Mutex //互斥锁

func increase() {
	for i := 0; i < 10000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}

}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increase()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

//{var wg sync.WaitGroup
//这行代码是声明一个 sync.WaitGroup 类型的变量 wg。
//WaitGroup 用于等待一组 goroutine 执行完毕，
//它提供了三个方法：Add()、Done() 和 Wait()。
//
//Add(n int)：表示向 WaitGroup 添加 n 个等待的 goroutine。
//Done()：表示一个等待的 goroutine 完成，相当于调用 Add(-1)。
//Wait()：阻塞当前的 goroutine，直到所有的等待的 goroutine 都调用了 Done()。
//在上面的示例中，我们使用 WaitGroup 来等待所有的 increase() goroutine 的执行完成。
//在 main() 函数中，我们使用 wg.Add(1) 来添加一个等待的 goroutine，
//然后在每个 increase() 函数的结尾调用 wg.Done() 来表示这个 goroutine 完成了。
//最后，我们调用 wg.Wait() 来阻塞 main() 函数，直到所有的等待的 goroutine 都执行完毕。
//
//通过 sync.WaitGroup 可以方便地进行并发控制和等待操作，
//确保所有的 goroutine 正确地执行完毕后再继续执行后续的操作。}
