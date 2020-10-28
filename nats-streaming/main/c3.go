package main

import (
	"fmt"
	"github.com/nats-io/go-nats-streaming"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {

	sc, err := stan.Connect("test-cluster", "c3", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		panic(err)
	}
	subj := "testing"
	_, err = sc.Subscribe(subj+"2", func(msg *stan.Msg) {
		log.Printf("Received a message: %s\n", string(msg.Data))
		newMsg := strings.Replace( string(msg.Data), "o", "T", -1)
		err = sc.Publish(subj+"3", []byte(newMsg))
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj+"3", newMsg)

	}, stan.DurableName("dn3"))
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

//go run c3.go