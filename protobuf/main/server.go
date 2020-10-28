package main

import (
	"fmt"
	stProto "github.com/FrankHitman/playGo/protobuf"
	"net"
	"os"
	//protobuf编解码库,下面两个库是相互兼容的，可以使用其中任意一个
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
	//fmt.Scanln()

}

//接收消息
func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		//读消息
		cnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		stReceive := &stProto.UserInfo{}
		pData := buf[:cnt]
		//protobuf解码
		err = proto.Unmarshal(pData, stReceive)
		if err != nil {
			panic(err)
		}
		fmt.Println("receive", conn.RemoteAddr(), stReceive)
		if stReceive.Message == "stop" {
			os.Exit(1) // deferred functions will not run.
		}
	}
}

// first go run server then go run client
// -----output----
//[sdy@centos playGo]$ go run protobuf/main/server.go
//
//new connect 127.0.0.1:50362
//receive 127.0.0.1:50362 message:"hello" length:5 cnt:1
//world
//receive 127.0.0.1:50362 message:"world" length:5 cnt:2
//receive 127.0.0.1:50362 message:"stop" length:4 cnt:5
//exit status 1


// reference https://yushuangqi.com/blog/2017/golangshi-yong-protobuf.html
//// Exit causes the current program to exit with the given status code.
//// Conventionally, code zero indicates success, non-zero an error.
//// The program terminates immediately; deferred functions are not run.
//func Exit(code int) {
//	if code == 0 {
//		// Give race detector a chance to fail the program.
//		// Racy programs do not have the right to finish successfully.
//		runtime_beforeExit()
//	}
//	syscall.Exit(code)
//}