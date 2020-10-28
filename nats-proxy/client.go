package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats"
	"gopkg.in/sohlich/nats-proxy.v1"
)

func main() {
	clientConn, _ := nats.Connect(nats.DefaultURL)
	natsClient, _ := natsproxy.NewNatsClient(clientConn)
	// Subscribe to URL /user/info
	natsClient.GET("/user/info", func(c *natsproxy.Context) {
		user := struct {
			Name string
		}{
			"Alan",
		}
		c.JSON(200, user)
	})
	defer clientConn.Close()

	// Waiting for signal to close the client
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C for exit.")
	<-sig
}
