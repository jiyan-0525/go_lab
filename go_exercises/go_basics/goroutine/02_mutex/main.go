// 02_mutex: counter++ 不是原子操作，必须加锁。
// 对照 C 的 pthread_mutex_lock / pthread_mutex_unlock。
package main

import (
	"fmt"
	"sync"
)

func main() {
	const N = 1000

	var (
		mu      sync.Mutex
		counter int
		wg      sync.WaitGroup
	)

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()   // pthread_mutex_lock
			counter++   // 读 - 改 - 写 三步，不加锁会丢更新
			mu.Unlock() // pthread_mutex_unlock
		}()
	}
	wg.Wait()

	// 因为加锁了，这里一定是 1000
	fmt.Println("counter =", counter)
}
