// 01_basic: 用 sync.WaitGroup 实现「启动 + 等待」。
// 对应 C 里的 pthread_create + pthread_join。
package main

import (
	"fmt"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done() // 函数退出时计数 -1，等价于线程「结束报告」

	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 本次要等 2 个 goroutine

	go say("world", &wg) // ≈ pthread_create
	go say("hello", &wg) // ≈ pthread_create

	wg.Wait() // ≈ pthread_join，阻塞到计数归零

	// 注意：两个 goroutine 的打印顺序每次都可能不同，这是正常的「并发」。
	fmt.Println("--- all done ---")
}
