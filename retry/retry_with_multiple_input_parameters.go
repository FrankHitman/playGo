package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 定义一个泛型函数类型，返回一个值和 error，并接受变长参数
type RetryableFunc[T any] func(args ...any) (T, error)

// 创建一个高阶函数，添加错误重试机制，返回值和 error，并接受变长参数
func withRetry[T any](fn RetryableFunc[T], retries int, delay time.Duration) RetryableFunc[T] {
	return func(args ...any) (T, error) {
		var result T
		var err error
		for i := 0; i < retries; i++ {
			if result, err = fn(args...); err == nil {
				return result, nil
			}
			fmt.Printf("Retry %d/%d failed with error: %v. Retrying in %v...\n", i+1, retries, err, delay)
			time.Sleep(delay)
		}
		return result, fmt.Errorf("after %d attempts, last error: %w", retries, err)
	}
}

// 示例函数，随机返回一个整数和错误，并接受一个整数参数
func unstableOperation(n int) (int, error) {
	if rand.Float32() < 0.7 { // 70% 的概率返回错误
		return 0, errors.New("operation failed")
	}
	return n + rand.Intn(100), nil
}
func main() {
	rand.Seed(time.Now().UnixNano())
	// 包装不稳定的操作函数，添加错误重试功能
	retryableOperation := withRetry(unstableOperation, 5, 2*time.Second)
	// 调用包装后的函数，传递参数
	arg := 10
	result, err := retryableOperation(arg)
	if err != nil {
		fmt.Println("Operation failed after retries:", err)
	} else {
		fmt.Printf("Operation succeeded with result: %d\n", result)
	}
}

// output
// # command-line-arguments
// ./retry_with_multiple_input_parameters.go:39:34: type func(n int) (int, error) of unstableOperation does not match RetryableFunc[T] (cannot infer T)
//
// Compilation finished with exit code 1
