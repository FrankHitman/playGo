# 性能测试

### 并发数，QPS与性能的关系
衡量 API 性能的指标主要有 3 个：

- 并发数（Concurrent）

并发数是指某个时间范围内，同时正在使用系统的用户个数。

广义上的并发数是指同时使用系统的用户个数，这些用户可能调用不同的 API。严格意义上的并发数是指同时请求同一个 API 的用户个数。本小节所讨论的并发数是严格意义上的并发数。

- 每秒查询数（QPS）

每秒查询数 QPS 是对一个特定的查询服务器在规定时间内所处理流量多少的衡量标准。

QPS = 并发数 / 平均请求响应时间。

- 请求响应时间（TTLB）

请求响应时间指的是从客户端发出请求到得到响应的整个时间。这个过程从客户端发起的一个请求开始，到客户端收到服务器端的响应结束。在一些工具中，请求响应时间通常会被称为 TTLB（Time to last byte，意思是从发送一个请求开始，到客户端收到最后一个字节的响应为止所消费的时间）。请求响应时间的单位一般为"秒”或“毫秒”。


衡量 API 性能的最主要指标是 QPS，但是在说明 QPS 时，需要指明是多少并发数下的 QPS，否则毫无意义，因为不同并发数下的 QPS 是不同的。
比如单用户 100 QPS 和 100 用户 100 QPS 是两个不同的概念，前者说明 API 可以在一秒内串行执行 100 个请求，
而后者说明在并发数为 100 的情况下，API 可以在一秒内处理 100 个请求。当 QPS 相同时，并发数越大，说明 API 性能越好，并发处理能力越强。

在并发数设置过大时，API 同时要处理很多请求，会频繁切换进程，而真正用于处理请求的时间变少，使得 QPS 反而会降低。
并发数设置过大时，请求响应时间也会变大。API 会有一个合适的并发数，在该并发数下，API 的 QPS 可以达到最大，但该并发数不一定是最佳并发数，还要参考该并发数下的平均请求响应时间。

### 性能测试工具
Web 性能测试工具，常用的有 Jmeter、AB、Webbench 和 Wrk。此处使用wrk。

#### 安装wrk
```shell script
git clone https://github.com/wg/wrk
cd wrk && make
ln -s $PWD/wrk /usr/local/bin/
```

#### 使用wrk
使用帮助，展示使用所需参数与意义
```shell script
Franks-Mac:api frank$ wrk --help
Usage: wrk <options> <url>                            
  Options:                                            
    -c, --connections <N>  Connections to keep open   
    -d, --duration    <T>  Duration of test           
    -t, --threads     <N>  Number of threads to use   
                                                      
    -s, --script      <S>  Load Lua script file       
    -H, --header      <H>  Add header to request      
        --latency          Print latency statistics   
        --timeout     <T>  Socket/request timeout     
    -v, --version          Print version details      
                                                      
  Numeric arguments may include a SI unit (1k, 1M, 1G)
  Time arguments may include a time unit (2s, 2m, 2h)
```
- -t: 线程数（线程数不要太多，是核数的 2 到 4 倍即可，多了反而会因为线程切换过多造成效率降低(也可能报错：too many open files)）
- -c: 并发数
- -d: 测试的持续时间，默认为 10s
- -T: 请求超时时间
- -H: 指定请求的 HTTP Header，有些 API 需要传入一些 Header，可通过 Wrk 的 -H 参数来传入
- --latency: 打印响应时间分布
- -s: 指定 Lua 脚本，Lua 脚本可以实现更复杂的请求

例如如果是测试post请求的接口性能，那么需要使用lua脚本来实现，以下login.lua：
```lua
wrk.method = "POST"
wrk.body   = "{\"mobile\":\"17221008976\",\"password\":\"password\"}"
wrk.headers["Content-Type"] = "application/json"
```
使用示例
```shell script
wrk -t8 -c3000 -d60s -T30s --latency http://127.0.0.1:443/api/v1.0/login -s login.lua
```

#### 执行结果意义
```shell script
Franks-Mac:api frank$ wrk -t8 -c3000 -d60s -T30s --latency http://127.0.0.1:443/api/v1.0/login -s login.lua 
Running 1m test @ http://127.0.0.1:443/api/v1.0/login
  8 threads and 3000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.24s     5.10s   29.70s    77.15%
    Req/Sec     4.93      4.67    29.00     81.70%
  Latency Distribution
     50%    7.72s 
     75%   11.49s 
     90%   16.60s 
     99%   26.48s 
  887 requests in 1.00m, 392.39KB read
  Socket errors: connect 2755, read 158, write 0, timeout 25
Requests/sec:     14.76
Transfer/sec:      6.53KB
```
意义解释
```
8 threads and 3000 connections: 用 8 个线程模拟 3000 个连接，分别对应 -t 和 -c 参数
Thread Stats： 线程统计
Latency: 响应时间，有平均值、标准偏差、最大值、正负一个标准差占比
Req/Sec: 每个线程每秒完成的请求数, 同样有平均值、标准偏差、最大值、正负一个标准差占比
Latency Distribution: 响应时间分布
50%: 50% 的响应时间为：7.72s
75%: 75% 的响应时间为：11.49s
90%: 90% 的响应时间为：16.60s
99%: 99% 的响应时间为：26.48s
887 requests in 1.00m, 392.39KB read: 1min 完成的总请求数（887）和数据读取量（392.39KB）
Socket errors: connect 2755, read 158, write 0, timeout 25: 错误统计
Requests/sec: QPS
Transfer/sec: TPS
```
### pprof
#### pprof是什么
PProf 是一个 Go 程序性能分析工具，可以分析 CPU、内存等性能。
Go 在语言层面上集成了 profile 采样工具，只需在代码中简单地引入 runtime/ppro 或者 net/http/pprof 包即可获取程序的 profile 文件，并通过该文件来进行性能分析。
runtime/pprof 还可以为控制台程序或者测试程序产生 pprof 数据。
其实 net/http/pprof 中只是使用 runtime/pprof 包来进行封装了一下，并在 HTTP 端口上暴露出来。

#### pprof使用
在gin的路由中注册pprof的路由
```go
// pprof router
import "github.com/gin-contrib/pprof"
// ...
	Router = gin.New()
	pprof.Register(Router)

```
通过 go tool pprof http://127.0.0.1:8080/debug/pprof/profile，可以获取 profile 采集信息并分析。

也可以直接在浏览器访问 http://localhost:8080/debug/pprof 来查看当前 API 服务的状态，包括 CPU 占用情况和内存使用情况等。

通过 topN 的输出可以分析出哪些函数占用 CPU 时间片最多，这些函数可能存在性能问题。此命令用于显示 profile 文件中的最靠前的 N 个样本（默认前10个sample），它的输出格式各字段的含义依次是：
- 采样点落在该函数中的总时间
- 采样点落在该函数中的百分比
- 上一项的累积百分比
- 采样点落在该函数，以及被它调用的函数中的总时间
- 采样点落在该函数，以及被它调用的函数中的总次数百分比
- 函数名

如果觉得不直观，可以直接生成函数调用图，通过调用图来判断哪些函数耗时最久，在 pprof 交互界面，执行 svg 生成 svg 文件。
但是需要确保系统已经安装 graphviz 命令（Mac OS:brew install graphviz; CentOS: yum -y install graphviz.x86_64）。

### go test
#### 基本概念
Go 语言有自带的测试框架 testing，可以用来实现单元测试和性能测试，通过 go test 命令来执行单元测试和性能测试。

go test 执行测试用例时，是以 go 包为单位进行测试的。执行时需要指定包名，比如：go test 包名，如果没有指定包名，默认会选择执行命令时所在的包。go test 在执行时会遍历以 _test.go 结尾的源码文件，执行其中以 Test、Benchmark、Example 开头的测试函数。其中源码文件需要满足以下规范：

- 文件名必须是 _test.go 结尾，跟源文件在同一个包。
- 测试用例函数必须以 Test、Benchmark、Example 开头
- 执行测试用例时的顺序，会按照源码中的顺序依次执行
- 单元测试函数 TestXxx() 的参数是 testing.T，可以使用该类型来记录错误或测试状态
- 性能测试函数 BenchmarkXxx() 的参数是 testing.B，函数内以 b.N 作为循环次数，其中 N 会动态变化
- 示例函数 ExampleXxx() 没有参数，执行完会将输出与注释 // Output: 进行对比
- 测试函数原型：func TestXxx(t *testing.T)，Xxx 部分为任意字母数字组合，首字母大写，例如： TestgenShortId 是错误的函数名，TestGenShortId 是正确的函数名
- 通过调用 testing.T 的 Error、Errorf、FailNow、Fatal、FatalIf 方法来说明测试不通过，通过调用 Log、Logf 方法来记录测试信息：
```
t.Log t.Logf     # 正常信息 
t.Error t.Errorf # 测试失败信息 
t.Fatal t.Fatalf # 致命错误，测试程序退出的信息
t.Fail     # 当前测试标记为失败
t.Failed   # 查看失败标记
t.FailNow  # 标记失败，并终止当前测试函数的执行，需要注意的是，我们只能在运行测试函数的 Goroutine 中调用 t.FailNow 方法，而不能在我们在测试代码创建出的 Goroutine 中调用它
t.Skip     # 调用 t.Skip 方法相当于先后对 t.Log 和 t.SkipNow 方法进行调用，而调用 t.Skipf 方法则相当于先后对 t.Logf 和 t.SkipNow 方法进行调用。方法 t.Skipped 的结果值会告知我们当前的测试是否已被忽略
t.Parallel # 标记为可并行运算
```
#### Benchmark
- 性能测试函数名必须以 Benchmark 开头，如 BenchmarkXxx 或 Benchmark_xxx
- go test 默认不会执行压力测试函数，需要通过指定参数 -test.bench 来运行压力测试函数，-test.bench 后跟正则表达式，如 go test -test.bench=".*" 表示执行所有的压力测试函数
- 在压力测试中，需要在循环体中指定 testing.B.N 来循环执行压力测试代码

### 综合使用
#### 寻找项目中login接口在高并发时候报错`too many open files`的原因

- 把web服务run起来`go run main.go`;
- 接着使用go tool请求pprof的api`go tool pprof http://127.0.0.1:443/debug/pprof/profile`;
- 然后立马使用wrk运行性能测试脚本，供pprof采集样本，wrk的执行时长最好大于30s，因为pprof的采样时长为30s
`wrk -t8 -c3000 -d60s -T30s --latency http://127.0.0.1:443/api/v1.0/login -s login.lua`

go tool pprof 与 top 执行结果
```shell script
Franks-Mac:wallbox frank$ go tool pprof http://127.0.0.1:443/debug/pprof/profile
Fetching profile over HTTP from http://127.0.0.1:443/debug/pprof/profile
Saved profile in /Users/frank/pprof/pprof.samples.cpu.004.pb.gz
Type: cpu
Time: Dec 31, 2019 at 3:37pm (CST)
Duration: 30.18s, Total samples = 25.90s (85.81%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
(pprof) top
Showing nodes accounting for 25.50s, 98.46% of 25.90s total
Dropped 159 nodes (cum <= 0.13s)
Showing top 10 nodes out of 38
      flat  flat%   sum%        cum   cum%
    22.36s 86.33% 86.33%     22.36s 86.33%  golang.org/x/crypto/blowfish.encryptBlock
     1.43s  5.52% 91.85%      1.43s  5.52%  runtime.newstack
     0.73s  2.82% 94.67%     25.12s 96.99%  golang.org/x/crypto/bcrypt.expensiveBlowfishSetup
     0.60s  2.32% 96.99%     24.37s 94.09%  golang.org/x/crypto/blowfish.ExpandKey
     0.24s  0.93% 97.92%      0.27s  1.04%  syscall.syscall
     0.14s  0.54% 98.46%      0.14s  0.54%  runtime.pthread_cond_wait
         0     0% 98.46%      0.13s   0.5%  database/sql.withLock
         0     0% 98.46%     25.41s 98.11%  github.com/gin-gonic/gin.(*Context).Next
         0     0% 98.46%     25.41s 98.11%  github.com/gin-gonic/gin.(*Engine).ServeHTTP
         0     0% 98.46%     25.41s 98.11%  github.com/gin-gonic/gin.(*Engine).handleHTTPRequest
(pprof) 
```

通过以上观察，login这个接口慢的主要原因是在blowfish.encryptBlock这个方法的耗时比较大。
于是循迹找到他的调用者`bcrypt.CompareHashAndPassword`

那么我们需要弄明白这个函数单独运行耗时多少？编写CompareHashAndPassword函数的性能测试代码，以下这段代码是官方自带的测试用例，拿来主义>_<。
以下chap_bench_test.go：
```go
package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkEqual(b *testing.B) {
	b.StopTimer()
	passwd := []byte("somepasswordyoulike")
	hash, _ := bcrypt.GenerateFromPassword(passwd, 10)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bcrypt.CompareHashAndPassword(hash, passwd)
	}
}

// b.StopTimer()：调用该函数停止压力测试的时间计数
// b.StartTimer()：重新开始时间
// 在 b.StopTimer() 和 b.StartTimer() 之间可以做一些准备工作，这样这些时间不影响我们测试函数本身的性能。
```

执行`go test -test.bench=".*"`
```shell script
Franks-Mac:pprof frank$ go test -test.bench=".*"
goos: darwin
goarch: amd64
pkg: github.com/FrankHitman/playGo/pprof
BenchmarkEqual-8              20          59242343 ns/op
PASS
ok      github.com/FrankHitman/playGo/pprof     1.370s
Franks-Mac:pprof frank$ 
```
以上结果得知 BenchmarkEqual 执行了 20 次，每次的执行平均时间是 59242343 纳秒，也就是59ms，总执行时间 1.370s。。

查询到[Why does AES encryption take more time than decryption?](https://security.stackexchange.com/questions/38055/why-does-aes-encryption-take-more-time-than-decryption/38056)  
确定是因为把密码原文加密耗时大，验证密码是否相同的方法是把传入的密码原文再加密一次，与原来加密后的密码做比较。
而加密比解密要慢很多的，因为加密是顺序一段一段的加，解密可以并发的解。

### 参考资料
- [基于 Go 语言构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b18630ee51d4506cd4fd834)
- [Why does AES encryption take more time than decryption?](https://security.stackexchange.com/questions/38055/why-does-aes-encryption-take-more-time-than-decryption/38056)  


