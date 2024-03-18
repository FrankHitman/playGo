// this file will show the implementation of consistent hashing with virtual nodes
// assume that the real node locations are located by crc32, the virtual node locations are located by crc64
package main

import (
	"fmt"
	"hash/crc64"
	"sort"
)

type Node64 struct {
	ID     string
	HashId uint64
}

type Ring64 struct {
	Nodes Nodes64
}

func NewRing64() *Ring64 {
	return &Ring64{Nodes: Nodes64{}}
}

func NewNode64(id string) *Node64 {
	return &Node64{
		ID:     id,
		HashId: crc64.Checksum([]byte(id), crc64.MakeTable(crc64.ISO)),
	}
}

func (r *Ring64) AddNode(id string) {
	node := NewNode64(id)
	r.Nodes = append(r.Nodes, *node)
	sort.Sort(r.Nodes)
}

// Declear Nodes64 type to implement sort.Interface: Len, Less, Swap
// otherwise sort.Sort(r.Nodes64) indicats error.
type Nodes64 []Node64

func (n Nodes64) Len() int {
	return len(n)
}
func (n Nodes64) Less(i, j int) bool {
	return n[i].HashId < n[j].HashId
}
func (n Nodes64) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (r *Ring64) Get(key string) string {
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= crc64.Checksum([]byte(key), crc64.MakeTable(crc64.ISO))
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	// if reach the end of the ring(array), jump to the start of the ring(array)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].ID
}

func main() {
	hashRing := NewRing64()
	hashRing.AddNode("node1")
	hashRing.AddNode("node2")
	hashRing.AddNode("node3")

	fmt.Println("he is at", crc64.Checksum([]byte("he"), crc64.MakeTable(crc64.ISO)), hashRing.Get("he"))
	fmt.Println("she is at", crc64.Checksum([]byte("she"), crc64.MakeTable(crc64.ISO)), hashRing.Get("she"))
	fmt.Println("node0 is at", crc64.Checksum([]byte("node0"), crc64.MakeTable(crc64.ISO)), hashRing.Get("node0"))
	fmt.Println("node1 is at", crc64.Checksum([]byte("node1"), crc64.MakeTable(crc64.ISO)), hashRing.Get("node1"))
	fmt.Println("node2 is at", crc64.Checksum([]byte("node2"), crc64.MakeTable(crc64.ISO)), hashRing.Get("node2"))
	fmt.Println("node3 is at", crc64.Checksum([]byte("node3"), crc64.MakeTable(crc64.ISO)), hashRing.Get("node3"))
}

// output
// Franks-Mac:system-design frank$ go run consistent-hashing.go
// he is at 3508889223 node2
// she is at 956988259 node3
// node0 is at 4075296214 node2
// node1 is at 2247042368 node1
// node2 is at 484865274 node2
// node3 is at 1809925228 node3
// Franks-Mac:system-design frank$ go run consistent-hashing2.go
// he is at 3663695889051942912 node3
// she is at 3663744542441472000 node3
// node0 is at 4801632614575243264 node1
// node1 is at 4833157811966836736 node1
// node2 is at 4738582219792056320 node2
// node3 is at 4644006627617275904 node3
