// 03_race: 故意制造数据竞争。
//
// 运行方式（必须加 -race）：
//   go run -race ./03_race
//
// 预期：打印 WARNING: DATA RACE 和一段调用栈，最后一行是 exit status 66
//（go run 会把它包装成自己的退出状态，所以 shell 里看到的是 $? = 1）。
// 去掉 -race 时会「看起来正常」——但 counter 通常小于 1000，这就是丢更新。
//
// 修复方式见 02_mutex。
package main

import (
	"fmt"
	"sync"
)

func main() {
	const N = 1000

	counter := 0 // 共享变量，没有任何保护
	var wg sync.WaitGroup

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			counter++ // 数据竞争：多个 goroutine 同时读-改-写
		}()
	}
	wg.Wait()

	fmt.Println("counter =", counter, "(期望 1000，实际经常更小)")
}
