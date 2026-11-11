# GO-LEARNING.md — 做 ft_irc_go 需要学的 Go 知识点

> 目标：从「Go 基本语法」走到「能独立完成 IRC 服务器的 JOIN / PRIVMSG / MODE」。
> 用法：每学完一项把 `[ ]` 改成 `[x]`，并按「验收标准」自测；不达标就回到对应练习。
> 对照文件：`main.go`（当前实现）、`plan.md`（完整设计）、`README.md`（功能清单）。

---
[text](.)
## 0. 开始前先确认一件事

**评审是否接受 goroutine-per-client？**

| 答案 | 你要走的路线 | epoll 是否要学 |
|---|---|---|
| 接受（idiomatic Go） | 路线 A：goroutine + channel | ❌ 不用 |
| 不接受（严格 42「一个 poll()」规则） | 路线 B：显式 epoll 单事件循环 | ✅ 第 4 层必须学 |

> 两条路线**不要混用**。先问清楚，再决定第 4 层学不学。
> 本文默认按路线 A 组织，路线 B 的额外内容集中在第 4 层末尾。

---

## 1. 基础语法（必须先会）

学完这层，你应该能**完整读懂 `main.go`，并解释每一行在干什么**。

- [ ] **Go 工具链**：`go mod init` / `go run` / `go build` / `go vet` / `go test`
      ⚠️ 本仓库目前**没有 `go.mod`**，第一件事：
      ```bash
      cd ft_irc_go && go mod init ft_irc_go
      ```
- [ ] **包与导入**：`package main`、`import (...)`、导出规则（大写首字母才导出）
- [ ] **变量与零值**：`var` / `:=`、`int` 零值 `0`、`string` 零值 `""`、`*T` 零值 `nil`
- [ ] **基本类型**：`string` / `byte` (= `uint8`) / `rune` / `int` / `bool`，以及三者转换
- [ ] **流程控制**：`if`、`for`（Go 只有 `for`）、`switch`（无 `break` 自动跳出）
      对应：`main.go:96` 的命令分发 `switch cmd`
- [ ] **函数**：多返回值、命名返回值、`defer`
      对应：`defer conn.Close()`（`main.go:41,55`）
- [ ] **error 惯用法**：`if err != nil { return err }`、`errors.New`、`fmt.Errorf("%w")`
- [ ] **struct 与方法**：字段、构造器 `NewServer`、**指针接收者 vs 值接收者**
      对应：`Client` / `Server`（`main.go:12-34`）
- [ ] **map 与 slice**：`make`、读不存在的 key 返回零值、`delete`、`append`、遍历无序
      对应：`map[string]*Client`（`main.go:24`）
- [ ] **strings 常用函数**：`SplitN` / `Split` / `Fields` / `TrimSpace` / `HasPrefix` /
      `ToUpper` / `TrimPrefix` / `Index`
      对应：`main.go:77,89,90`
- [ ] **接口**：`net.Conn` 是什么、`error` 接口、类型断言 `v, ok := x.(T)`

**验收标准**：不看注释，能口述 `main.go` 从 `main()` 到 `parseCommand()` 的完整调用链。

---

## 2. 网络与协议（项目核心）

- [ ] **TCP 服务器骨架**：`net.Listen` → `Accept` 循环 → `go handle(conn)` → `Close`
      对应：`Start()`（`main.go:36-52`）
- [ ] **TCP 是字节流，没有消息边界**（粘包 / 拆包）
      → IRC 靠 `\r\n` 分隔，必须按行读，不能假设一次 `Read` = 一条命令
- [ ] **`bufio.Reader`**：`NewReader` / `ReadString('\n')` / `ReadBytes`
      对应：`main.go:71-76`
- [ ] **`bufio.Writer` 与 `Flush`**：为什么批量写要 flush
- [ ] **IRC 报文格式**：`[:prefix] COMMAND [params...] [":" trailing]`
      trailing 可以含空格 → 所以 `SplitN(line, " ", 2)` 不能直接切到底
- [ ] **协议拼装**：`fmt.Sprintf` 拼 `:<server> 001 <nick> :Welcome...`
- [ ] **字节与字符串**：`[]byte(s)` / `string(b)`、`conn.Write([]byte(...))`
- [ ] **超时控制**：`SetReadDeadline` / `SetWriteDeadline`（防慢客户端卡死服务器）
- [ ] **`net.Pipe()`**：测试时当假连接用（第 5 层会用到）

**验收标准**：独立写出一个 echo server，`nc 127.0.0.1 9000` 输入什么回什么，
并用 `printf 'a\r\nb\r\n' | nc` 验证按行处理正确。

---

## 3. 并发（这个项目的灵魂）

- [ ] **goroutine**：`go f()`、生命周期、主 goroutine 退出会带走一切
      对应：`go s.handleClient(conn)`（`main.go:50`）、写协程（`main.go:64-68`）
- [ ] **channel**：无缓冲 vs 有缓冲、发送/接收阻塞语义、`range ch`、`close(ch)`
      对应：`sendChan: make(chan string, 100)`（`main.go:60`）
- [ ] **`select`**：多路等待、`default` 做非阻塞、`time.After` 做超时
- [ ] **背压**：channel 满了怎么办（丢弃？断开？）—— 慢客户端策略
- [ ] **`sync.Mutex` / `RWMutex`**：保护共享 map
      对应：`Server.mu`（`main.go:25`）—— **目前声明了但从未使用**
- [ ] **`sync.WaitGroup`**：关闭服务器时等所有 goroutine 退出
- [ ] **`context`**：`WithCancel` / `WithTimeout`、把 `ctx` 传给 goroutine
- [ ] **数据竞争**：`go test -race`、`go run -race`，理解 map 并发读写会 panic
- [ ] **goroutine 泄漏**：channel 永不关闭 → 接收方永久阻塞

**验收标准**：能指出并修掉当前 `main.go` 里的这 3 个真实 bug：

1. **数据竞争**：`clients` map 被多个 goroutine 读写，`mu` 从未 `Lock()`。
2. **goroutine 泄漏**：写协程 `for msg := range client.sendChan` 永不退出，
   因为没有任何地方 `close(client.sendChan)`。
3. **关闭顺序**：`QUIT` 只是 `return`，连接关掉了但写协程还在写、客户端没被移出 map、
   也没有广播 `QUIT` 给同频道的人。

---

## 4. 工程化（决定能否过评审）

- [ ] **多文件 / 多包**：`internal/server`、`internal/irc`、`internal/state`、`internal/commands`
      （结构见 `plan.md` §1；规则：只有 `main` 能 import `server`）
- [ ] **测试**：table-driven test、`t.Run`、`go test -run/-v/-race`、`net.Pipe()` 假连接
- [ ] **模糊测试**：`go test -fuzz`（`plan.md` Phase 5 用在协议解析器上）
- [ ] **命令行参数**：`os.Args` 或 `flag`，实现 `./ircserv <port> <password>`
      （README 里的第一个 TODO）
- [ ] **信号处理**：`os/signal` + `signal.Notify`，`SIGINT`/`SIGTERM` 优雅退出
- [ ] **日志**：`log` 或 `slog`，分级 + 带客户端标识
- [ ] **代码质量**：`gofmt -l .` 无输出、`go vet ./...` 干净、`go test -race ./...` 全绿

### 路线 B 额外内容（仅当学校强制「一个 poll()」）

- [ ] `syscall.Conn` → `SyscallConn()` → `RawConn.Control` 拿到 fd
- [ ] `golang.org/x/sys/unix`：`EpollCreate1` / `EpollCtl` / `EpollWait`
- [ ] 非阻塞 fd、`EAGAIN` / `EWOULDBLOCK` 循环读到空
- [ ] 事件掩码：`EPOLLIN` / `EPOLLOUT` / `EPOLLERR` / `EPOLLHUP` / `EPOLLRDHUP`
- [ ] 自己维护每连接的输入/输出 `[]byte` 缓冲
- [ ] 单 goroutine 拥有全部状态（循环内不需要 mutex）

**暂时不用学**：泛型、反射、`cgo`、`unsafe`、Docker/K8s。

---

## 5. 学习顺序与里程碑

| 阶段 | 内容 | 产出 / 验收 | 建议用时 |
|---|---|---|---|
| M1 | 第 1 层全部 | 能读懂并讲清 `main.go` | 3–5 天 |
| M2 | 第 2 层 + 练习 1–2 | 独立写出 echo server，`nc` 可测 | 2–3 天 |
| M3 | 第 3 层 + 练习 3–4 | 修掉上面 3 个 bug，`-race` 通过 | 3–5 天 |
| M4 | 第 4 层 + 练习 5–6 | `./ircserv <port> <password>` 可跑 | 3–5 天 |
| M5 | 重构 + JOIN / PRIVMSG | 两个 `nc` 客户端能互相发消息 | 1 周 |
| M6 | MODE / KICK / INVITE / TOPIC | 对照 `plan.md` §4 检查表 | 1 周 |

> 路线 B 的 epoll 建议单独排在 M2 之前或之后整块学 3–4 天，别和并发混着学。

---

## 6. 练习题清单（放在 `ex/` 下）

每个练习一个独立目录，带 `main.go` 和自己的 `go test`。

| # | 练习 | 覆盖知识点 | 验收 |
|---|---|---|---|
| 1 | `ex/01_basics/` — 字符串解析器：把 `PASS secret`、`NICK alice` 拆成命令 + 参数 | 第 1 层 | `go test ./ex/01_basics` 全绿 |
| 2 | `ex/02_echo/` — TCP echo server，按行回显 | 第 2 层 | `nc` 手动测试通过 |
| 3 | `ex/03_queue/` — 给 echo server 加 `sendChan` + 写协程 + 优雅关闭 | 第 3 层（channel / close） | `close` 后写协程能退出，无泄漏 |
| 4 | `ex/04_lock/` — 用 mutex 保护 `map[string]*Client`，并写并发测试 | 第 3 层（race） | `go test -race` 全绿 |
| 5 | `ex/05_args/` — `./prog <port> <password>` 参数校验 + 信号退出 | 第 4 层 | 参数错误时退出码 1，`Ctrl-C` 干净退出 |
| 6 | `ex/06_parse/` — 完整 IRC 消息解析（prefix / command / params / trailing） | 第 1+2 层 | table-driven 测试覆盖边界 |
| 7 | `ex/07_gomod/` —（可选）用 `net.Pipe` 写一个 `parseCommand` 的单元测试 | 第 4 层 | 无需真 TCP 即可测协议 |

---

## 7. 新手最容易踩的坑（对照检查）

- [ ] `map` 并发读写 → 直接 panic，且 `-race` 会报；永远加锁或用单 goroutine 拥有
- [ ] 忘记 `defer conn.Close()` 或重复 `Close()` 导致写 panic
- [ ] 写阻塞：客户端不读数据时 `conn.Write` 会卡住整个写协程 → 需要 deadline
- [ ] `strings.TrimSpace` 会吃掉 trailing 里有意义的空格（`:Alice Wonderland`）
- [ ] 用 `len(s)` 当字符数（中文/emoji 要 `utf8.RuneCountInString`）
- [ ] `for msg := range ch` 永远不退出 → 忘了 `close`
- [ ] 在 goroutine 里修改外层变量却没同步 → 数据竞争
- [ ] 把 `Accept` 返回的 `err` 直接 `continue`，服务器关闭时变成死循环刷 CPU

---

## 8. 每次提交前的自检命令

```bash
gofmt -l .              # 无输出
go vet ./...            # 无输出
go test -race ./...     # 全绿
go build -o ircserv .
```

---

## 9. 参考

- RFC 1459 / RFC 2812（IRC 协议）
- `plan.md` —— 完整阶段划分与评审检查表
- `README.md` —— 当前功能与 TODO
