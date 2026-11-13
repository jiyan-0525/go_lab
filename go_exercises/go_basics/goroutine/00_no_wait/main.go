// 00_no_wait: 你原来的写法。
// go 关键字只负责「启动」，不负责「等待」。
// main 一返回，进程立刻退出，其他 goroutine 全部被杀。
package main

import "fmt"

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	// 多跑几次你会发现：world 一行都不打印，或者只打印几行。
	// 因为 main 的 5 次循环几微秒就跑完了，调度器还没来得及切过去。
}
