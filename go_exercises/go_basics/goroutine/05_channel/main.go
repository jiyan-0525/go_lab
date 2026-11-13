// 05_channel: 不用 WaitGroup，用 channel 等待。
// channel 是 Go 独有的东西，C 里没有直接对应物（最接近的是管道 + 条件变量）。
package main

import "fmt"

func say(s string, done chan<- struct{}) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	done <- struct{}{} // 干完了，发一个信号（struct{} 不占内存，只需信号本身）
}

func main() {
	// 缓冲 2：两个 goroutine 各发一次，即使 main 还没开始收也不会阻塞
	done := make(chan struct{}, 2)

	go say("world", done)
	go say("hello", done)

	<-done // 收到第一个信号
	<-done // 收到第二个信号，等价于 join 了两个
	fmt.Println("--- all done ---")
}
