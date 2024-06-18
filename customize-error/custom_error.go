package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型
type MyCustomError struct {
	Code    int
	Message string
}

// 实现 error 接口
func (e *MyCustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// 生成自定义错误的函数
func doSomething(n int) error {
	if n < 0 {
		return &MyCustomError{
			Code:    400,
			Message: "Negative value not allowed",
		}
	}
	if n == 0 {
		return errors.New("Generic error: value cannot be zero")
	}
	return nil
}

// 处理函数，根据错误类型进行特殊处理
func handleError() {
	values := []int{-1, 0, 1}
	for _, val := range values {
		err := doSomething(val)
		if err != nil {
			var myErr *MyCustomError
			if errors.As(err, &myErr) {
				fmt.Println("Custom Error occurred:", myErr)
				// 进行自定义错误的特殊处理
				// 比如，记录日志、重试操作等
			} else {
				fmt.Println("Generic Error occurred:", err)
			}
		} else {
			fmt.Println("Operation successful for value:", val)
		}
	}
}
func main() {
	handleError()
}

// output
// Custom Error occurred: Error 400: Negative value not allowed
// Generic Error occurred: Generic error: value cannot be zero
// Operation successful for value: 1
// MyCustomError支持定义成空的 struct{}
