# goroutine 实验目录

对照 C 的 pthread 学习 Go 并发。每个子目录是一个可独立运行的程序。

## 运行方式

```bash
cd goroutine

go run ./00_no_wait    # 原版：world 根本不打印（main 返回杀进程）
go run ./01_basic      # WaitGroup = create + join，正确等待
go run ./02_mutex      # Mutex 保护 counter++，结果恒为 1000
go run -race ./03_race # 数据竞争，race detector 报错，counter < 1000
go run ./04_sleep      # tour 原版：靠 sleep 拖住 main，输出交错
go run ./05_channel    # 用 channel 代替 WaitGroup 做等待
```

`03_race` 一定要带 `-race`。不带的话它“看起来能跑”，但结果通常是错的。
带 `-race` 时会打印 `WARNING: DATA RACE`，最后一行 `exit status 66`（`go run` 自己会返回 1）。

## 和 C 的对应关系

| C | Go | 在哪里看 |
|---|---|---|
| `pthread_create` | `go f()` | 全部 |
| `pthread_join` | `wg.Wait()` / `<-done` | 01、05 |
| `pthread_mutex_lock/unlock` | `mu.Lock()/Unlock()` | 02 |
| 数据竞争 | 一样存在 | 03 |
| 线程栈固定 8MB | goroutine 栈 2KB 起、按需增长 | — |
| 1 线程 = 1 内核调度实体 | M:N 调度（G-M-P） | — |

## 建议阅读顺序

1. `00_no_wait` — 先看清 main 返回的后果
2. `01_basic` — 学会正确等待
3. `02_mutex` / `03_race` — 理解为什么需要锁
4. `04_sleep` — 回头再看，明白它只是歪打正着
5. `05_channel` — 进入 Go 特有的通信方式
