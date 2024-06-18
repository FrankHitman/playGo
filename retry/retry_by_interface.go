package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// 定义一个泛型函数类型，返回一个值和 error，并接受变长参数
type RetryableFunc func(args ...interface{}) (interface{}, error)

// 创建一个高阶函数，添加错误重试机制，返回值和 error，并接受变长参数
func withRetry(fn interface{}, retries int, delay time.Duration) RetryableFunc {
	return func(args ...interface{}) (interface{}, error) {
		var result interface{}
		var err error
		fnValue := reflect.ValueOf(fn)
		fnType := fnValue.Type()
		if fnType.Kind() != reflect.Func {
			return nil, errors.New("fn is not a function")
		}
		for i := 0; i < retries; i++ {
			inputs := make([]reflect.Value, len(args))
			for j, arg := range args {
				inputs[j] = reflect.ValueOf(arg)
			}
			outputs := fnValue.Call(inputs)
			if len(outputs) != 2 {
				return nil, errors.New("fn must return two values")
			}
			result = outputs[0].Interface()
			err, _ = outputs[1].Interface().(error)
			if err == nil {
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
		fmt.Printf("Operation succeeded with result: %d\n", result.(int))
	}
}

// output
// Retry 1/5 failed with error: operation failed. Retrying in 2s...
// Retry 2/5 failed with error: operation failed. Retrying in 2s...
// Retry 3/5 failed with error: operation failed. Retrying in 2s...
// Retry 4/5 failed with error: operation failed. Retrying in 2s...
// Retry 5/5 failed with error: operation failed. Retrying in 2s...
// Operation failed after retries: after 5 attempts, last error: operation failed
