package main

import (
	"log"
	"net"
	"time"
)

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Printf("%d: dial error: %s", i, err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}

func main() {
	var sl []net.Conn
	for i := 1; i < 1000; i++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}

	time.Sleep(time.Second * 10000)
}

// first execute go run server.go, then go run client.go
// 可以看出Client初始时成功地一次性建立了128个连接，然后后续每阻塞近10s才能成功建立一条连接。
// 也就是说在server端 backlog满时(未及时accept)，客户端将阻塞在Dial上，直到server端进行一次accept。
// 至于为什么是128，这与darwin 下的默认设置有关：