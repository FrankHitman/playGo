// this file show the algorithm of consistent hashing
package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Node struct {
	ID     string
	HashId uint32
}

type Ring struct {
	Nodes Nodes
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

func NewNode(id string) *Node {
	return &Node{
		ID:     id,
		HashId: crc32.ChecksumIEEE([]byte(id)),
	}
}

func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	r.Nodes = append(r.Nodes, *node)
	sort.Sort(r.Nodes)
}

// Declear Nodes type to implement sort.Interface: Len, Less, Swap
// otherwise sort.Sort(r.Nodes) indicats error.
type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}
func (n Nodes) Less(i, j int) bool {
	return n[i].HashId < n[j].HashId
}
func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (r *Ring) Get(key string) string {
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= crc32.ChecksumIEEE([]byte(key))
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	// if reach the end of the ring(array), jump to the start of the ring(array)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].ID
}

func main() {
	hashRing := NewRing()
	hashRing.AddNode("node1")
	hashRing.AddNode("node2")
	hashRing.AddNode("node3")

	fmt.Println("he is at", crc32.ChecksumIEEE([]byte("he")), hashRing.Get("he"))
	fmt.Println("she is at", crc32.ChecksumIEEE([]byte("she")), hashRing.Get("she"))
	fmt.Println("node0 is at", crc32.ChecksumIEEE([]byte("node0")), hashRing.Get("node0"))
	fmt.Println("node1 is at", crc32.ChecksumIEEE([]byte("node1")), hashRing.Get("node1"))
	fmt.Println("node2 is at", crc32.ChecksumIEEE([]byte("node2")), hashRing.Get("node2"))
	fmt.Println("node3 is at", crc32.ChecksumIEEE([]byte("node3")), hashRing.Get("node3"))
}

// output
// Franks-Mac:system-design frank$ go run consistant-hashing.go
// he is at 3508889223 node2
// she is at 956988259 node3
// node0 is at 4075296214 node2
// node1 is at 2247042368 node1
// node2 is at 484865274 node2
// node3 is at 1809925228 node3
