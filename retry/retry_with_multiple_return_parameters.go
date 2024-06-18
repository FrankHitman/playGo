package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 定义一个泛型函数类型，返回一个值和 error
type RetryableFunc[T any] func() (T, error)

// 创建一个高阶函数，添加错误重试机制，返回值和 error
func withRetry[T any](fn RetryableFunc[T], retries int, delay time.Duration) RetryableFunc[T] {
	return func() (T, error) {
		var result T
		var err error
		for i := 0; i < retries; i++ {
			if result, err = fn(); err == nil {
				return result, nil
			}
			fmt.Printf("Retry %d/%d failed with error: %v. Retrying in %v...\n", i+1, retries, err, delay)
			time.Sleep(delay)
		}
		return result, fmt.Errorf("after %d attempts, last error: %w", retries, err)
	}
}

// 示例函数，随机返回一个整数和错误
func unstableOperation() (int, error) {
	if rand.Float32() < 0.7 { // 70% 的概率返回错误
		return 0, errors.New("operation failed")
	}
	return rand.Intn(100), nil
}
func main() {
	rand.Seed(time.Now().UnixNano())
	// 包装不稳定的操作函数，添加错误重试功能
	retryableOperation := withRetry(unstableOperation, 5, 2*time.Second)
	// 调用包装后的函数
	result, err := retryableOperation()
	if err != nil {
		fmt.Println("Operation failed after retries:", err)
	} else {
		fmt.Printf("Operation succeeded with result: %d\n", result)
	}
}

// output
// Retry 1/5 failed with error: operation failed. Retrying in 2s...
// Operation succeeded with result: 3
