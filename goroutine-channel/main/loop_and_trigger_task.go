package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// GPT prompt:
// 1. 每5秒钟执行一次定时任务； 2. 可以通过 channel 从外部传入信号量进行触发执行任务；
// 3. 这是一个永久的 goroutine 线程，在退出的时候要求优雅的进行退出。
// 4. 定时任务是一个 HTTP request 请求。 5. 执行结果放入 channel 传到另外一个 线程里面。

// 定时任务：发送 HTTP 请求
func performHTTPRequest() (string, error) {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Printf("%s : Response status: %s \n", time.Now().String(), resp.Status)
	return fmt.Sprintf("%s: Response status: %s", time.Now().String(), resp.Status), nil

}

// 定时任务 goroutine
func startTask(ctx context.Context, signalChan <-chan struct{}, resultChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 定时器信号
			result, err := performHTTPRequest()
			if err != nil {
				resultChan <- fmt.Sprintf("Error: %v", err)
			} else {
				resultChan <- result
			}
		case <-signalChan:
			// 外部信号
			result, err := performHTTPRequest()
			if err != nil {
				resultChan <- fmt.Sprintf("Error: %v", err)
			} else {
				resultChan <- result
			}
		case <-ctx.Done():
			// 上下文取消信号
			fmt.Printf("%s Exiting goroutine...\n", time.Now().String())
			// fmt.Println(" Exiting goroutine...")
			return
		}
	}
}
func main() {
	// 创建上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())
	// 创建信号通道和结果通道
	signalChan := make(chan struct{})
	resultChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	// 启动定时任务 goroutine
	go startTask(ctx, signalChan, resultChan, &wg)
	// 模拟外部信号触发任务执行
	go func() {
		time.Sleep(10 * time.Second)
		signalChan <- struct{}{}
		time.Sleep(10 * time.Second)
		signalChan <- struct{}{}
		// 关闭上下文，优雅退出
		cancel()
	}()
	// 处理结果，如果没有以下四行代码，程序会运行不到 "Main function exiting..."
	go func() {
		wg.Wait()         // 等待所有任务完成
		close(resultChan) // 关闭结果通道
	}()
	// 输出结果
	func() {
		for result := range resultChan {
			fmt.Println(result)
		}
	}()
	// 确保所有任务在程序退出前完成
	time.Sleep(25 * time.Second)
	fmt.Println("Main function exiting...", time.Now().String())
}

// output
// 2024-06-13 16:30:42.355825 +0800 CST m=+9.796349625 : Response status: 200 OK
// 2024-06-13 16:30:42.356158 +0800 CST m=+9.796682047: Response status: 200 OK
// 2024-06-13 16:30:42.865435 +0800 CST m=+10.305960149 : Response status: 200 OK
// 2024-06-13 16:30:42.865453 +0800 CST m=+10.305978666: Response status: 200 OK
// 2024-06-13 16:30:43.275769 +0800 CST m=+10.716295481 : Response status: 200 OK
// 2024-06-13 16:30:43.275786 +0800 CST m=+10.716311965: Response status: 200 OK
// 2024-06-13 16:30:47.884018 +0800 CST m=+15.324555016 : Response status: 200 OK
// 2024-06-13 16:30:47.88406 +0800 CST m=+15.324596700: Response status: 200 OK
// 2024-06-13 16:30:52.826775 +0800 CST m=+20.267323045 : Response status: 200 OK
// 2024-06-13 16:30:52.826817 +0800 CST m=+20.267365678: Response status: 200 OK
// 2024-06-13 16:30:53.209626 +0800 CST m=+20.650174775 : Response status: 200 OK
// 2024-06-13 16:30:53.209664 +0800 CST m=+20.650213019 Exiting goroutine...
// 2024-06-13 16:30:53.209652 +0800 CST m=+20.650200570: Response status: 200 OK
// Main function exiting... 2024-06-13 16:31:18.209994 +0800 CST m=+45.650600650
