### does the code below panic?
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)
	int_chan := make(chan int, 1)
	string_chan := make(chan string)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
	}
}

```
### answer
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
```
It will panic at ```string_chan <- "hello"```, because this channel with no buffer, when put something in it ,
it will check the consumer is already or not, when consumer is not ready, it will panic.
```int_chan <- 1``` will panic if add ```int_chan <- 2``` before it.