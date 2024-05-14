package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mymap = make(map[int]int)

var oi sync.Mutex
var cond2 = sync.NewCond(&oi)
var counter int

func Get(m map[int]int, k int) {

	timeout := time.After(1 * time.Second) // 设置超时时间为1秒
	for {
		oi.Lock()
		for key, _ := range m {
			if k == key {
				fmt.Println("查询到：", key)
				cond2.Wait()
			}
		}
		fmt.Println("未查询到", k)
		cond2.Broadcast()
		oi.Unlock()

		select {
		case <-timeout:
			return // 如果超时，则退出循环
		default:
		}
		// 增加计数器
		counter++
		if counter%2 == 0 {
			break
		}
	}
}

func Put(m map[int]int, x int, y int) {

	timeout := time.After(1 * time.Second) // 设置超时时间为1秒
	for {
		oi.Lock()
		for key := range m {
			if x == key {
				fmt.Println("该序号已被使用:", x)
				cond2.Wait()

			}
		}

		m[x] = y
		fmt.Println("写入了号码：", x)
		cond2.Broadcast()
		oi.Unlock()

		select {
		case <-timeout:
			return // 如果超时，则退出循环
		default:
		}
		// 增加计数器
		counter++
		if counter%2 == 0 {
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for len(mymap) < 100 {

		go Get(mymap, rand.Intn(100))
		go Put(mymap, rand.Intn(100), rand.Intn(100))
	}
	select {}

}
