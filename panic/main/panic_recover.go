// panic_recover.go
package main

import (
	"log"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	log.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func main() {
	log.Printf("Calling test\r\n")
	test()
	//protect(badCall)
	log.Printf("Test completed\r\n")
}

//----test() output----
//2019/05/07 09:38:10 Calling test
//2019/05/07 09:38:10 Panicing bad end
//2019/05/07 09:38:10 Test completed

func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() //   possible runtime-error
}

//----protect(badCall) output-----
//2019/05/07 09:35:50 Calling test
//2019/05/07 09:35:50 start
//2019/05/07 09:35:50 done
//2019/05/07 09:35:50 run time panic: bad end
//2019/05/07 09:35:50 Test completed

