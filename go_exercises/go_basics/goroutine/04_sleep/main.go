// 04_sleep: Go Tour 原版，每打印一次睡 100ms。
//
// 现象：hello / world 交错出现。
// 原因不是 go 关键字会「并行执行」，而是：
//   1. Sleep 让当前 goroutine 主动让出 CPU，调度器去跑另一个；
//   2. 整体耗时 500ms，main 迟迟不返回，world 才有机会跑完。
//
// 结论：这里的 Sleep 只是碰巧拖住了 main，不是正确的同步方式。
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i, "at", time.Now().Format("05.000"))
	}
}

func main() {
	go say("world")
	say("hello")
}
