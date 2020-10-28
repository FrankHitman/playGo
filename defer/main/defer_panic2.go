package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	go func() {
		defer func() {
			fmt.Println("defer begin")
			if err := recover(); err != nil{
				fmt.Println("get error: ", err)
				//return
			}
			fmt.Println("defer end")
		}()
		var user = os.Getenv("USER_")
		fmt.Println("user is: ", user)
		var i int
		for {
			fmt.Println("i is: ", i)
			i += 1
			time.Sleep(time.Second)

			//if user == "" {
			//	panic("should set user env.")
			//}
		}


	}()

	time.Sleep(6 * time.Second)
}

// -----output------
//user is:
//i is:  0
//defer begin
//get error:  should set user env.
//defer end



//user is:
//i is:  0
//i is:  1
//i is:  2
//i is:  3
//i is:  4
//i is:  5
