package main

import (
	"log"
	"strings"

	"github.com/levigross/grequests"
)

type KvItem struct {
	LockIndex   int64  `json:"lock_index"`
	Key         string `json:"key"`
	Flags       int64
	Value       string `json:"value"`
	CreateIndex int64  `json:"create_index"`
	ModifyIndex int64
}

func main() {
	// url := "http://101.91.224.210:8500/v1/kv/?recurse"
	// resp, err := grequests.Get(url, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// respStr := resp.String()
	// kvItems := make([]*KvItem, 0)
	//
	// if err := json.Unmarshal([]byte(respStr), &kvItems); err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// log.Printf("%+v",kvItems)
	// log.Printf("%+v",kvItems[0])
	//
	// for _, v := range kvItems {
	// 	decV, err := base64.StdEncoding.DecodeString(v.Value)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	v.Value = string(decV)
	// }
	// log.Println(kvItems)
	// log.Printf("%+v",kvItems[0])
	//

	url := "http://101.91.224.210:8500/v1/kv/foo4"
	r := strings.NewReader("0123456789")

	ro := &grequests.RequestOptions{RequestBody: r}
	resp, err := grequests.Put(url, ro)
	if err != nil {
		log.Fatalln(err)
	}

	respStr := resp.String()
	log.Println(respStr)

}

// use grequests call consul http api with raw text/plain format