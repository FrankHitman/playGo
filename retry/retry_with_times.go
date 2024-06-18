package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 定义一个函数类型
type RetryableFunc func() error

// 创建一个高阶函数，添加错误重试机制
func withRetry(fn RetryableFunc, retries int, delay time.Duration) RetryableFunc {
	return func() error {
		var err error
		for i := 0; i < retries; i++ {
			if err = fn(); err == nil {
				return nil
			}
			fmt.Printf("Retry %d/%d failed with error: %v. Retrying in %v...\n", i+1, retries, err, delay)
			time.Sleep(delay)
		}
		return fmt.Errorf("after %d attempts, last error: %w", retries, err)
	}
}

// 示例函数，随机返回错误
func unstableOperation() error {
	if rand.Float32() < 0.8 { // 70% 的概率返回错误
		return errors.New("operation failed")
	}
	return nil
}
func main() {
	rand.Seed(time.Now().UnixNano())
	// 包装不稳定的操作函数，添加错误重试功能
	retryableOperation := withRetry(unstableOperation, 5, 2*time.Second)
	// 调用包装后的函数
	if err := retryableOperation(); err != nil {
		fmt.Println("Operation failed after retries:", err)
	} else {
		fmt.Println("Operation succeeded")
	}
}

// 90% possibility fail output
// Retry 1/5 failed with error: operation failed. Retrying in 2s...
// Retry 2/5 failed with error: operation failed. Retrying in 2s...
// Retry 3/5 failed with error: operation failed. Retrying in 2s...
// Retry 4/5 failed with error: operation failed. Retrying in 2s...
// Retry 5/5 failed with error: operation failed. Retrying in 2s...
// Operation failed after retries: after 5 attempts, last error: operation failed

// 80% possibility fail output
// Retry 1/5 failed with error: operation failed. Retrying in 2s...
// Operation succeeded
