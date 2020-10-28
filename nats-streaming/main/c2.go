package main

import (
	"fmt"
	"github.com/nats-io/go-nats-streaming"
	"log"
	"os"
	"os/signal"
)

func main() {

	sc, err := stan.Connect("test-cluster", "c2", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		panic(err)
	}
	subj := "testing"
	_, err = sc.Subscribe(subj, func(msg *stan.Msg) {
		log.Printf("Received a message: %s\n", string(msg.Data))
		newMsg := string(msg.Data) + "_abc"
		err = sc.Publish(subj+"2", []byte(newMsg))
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj+"2", newMsg)

	}, stan.DurableName("dn2"))
	if err != nil {
		panic(err)
	}

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}

//go run c2.go
