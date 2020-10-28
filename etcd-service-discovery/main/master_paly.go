package main

import (
	"log"
	"time"

	sd "github.com/FrankHitman/playGo/etcd-service-discovery"
)

func main() {
	m, err := sd.NewMaster("sd-test", []string{
		"http://127.0.0.1:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("all ->", m.GetNodes())
		log.Println("all(strictly) ->", m.GetNodesStrictly())
		time.Sleep(time.Second * 2)
	}
}