package main

import "fmt"

type LRU struct {
	Length   int
	Elements map[string]int
}

func (lru *LRU) get(ele string) {
	if _, isExist := lru.Elements[ele]; isExist{
		lru.Elements[ele] += 1
	}
}

func (lru *LRU) put(ele string) {
	_, isExist := lru.Elements[ele]
	if isExist == false {
		if len(lru.Elements) >= lru.Length {
			var min int
			var minKey string
			for minKey, min = range lru.Elements {
				break
			}
			for k, v := range lru.Elements {
				if v < min {
					min = v
					minKey = k
				}
			}
			delete(lru.Elements, minKey)
		}
		lru.Elements[ele] = 1
	} else {
		lru.get(ele)
	}
}

func New() *LRU {
	lru := &LRU{
		Length: 3,
		Elements:make(map[string]int,3),
	}
	return lru
}

func main() {
	lru := New()

	lru.put("hello")
	lru.put("world")
	lru.put("ha")
	lru.get("ha")
	lru.put("ha")
	lru.get("hello")
	lru.put("new")
	lru.put("new")


	for k, v :=range lru.Elements{
		fmt.Println(k,v)
	}
}