package main
import "fmt"

func main() {
	//function1()
	f()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() string{
	fmt.Printf("Function2: Deferred until the end of the calling function!")
	return "end in function2"
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("i is  ", i)
	}
}

//i is   4
//i is   3
//i is   2
//i is   1
//i is   0

// 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
// 为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。


//关闭文件流 （详见 第 12.2 节）
//// open a file
//defer file.Close()
//
//解锁一个加锁的资源 （详见 第 9.3 节）
//mu.Lock()
//defer mu.Unlock()
//
//打印最终报告
//printHeader()
//defer printFooter()
//
//关闭数据库链接
//// open a database connection
//defer disconnectFromDB()
