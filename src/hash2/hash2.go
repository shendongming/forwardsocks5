package hash2

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"

)

const DEFAULT_REPLICAS = 256

type HashRing []uint32

func (c HashRing) Len() int {
	return len(c)
}

func (c HashRing) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c HashRing) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//节点只要唯一标识 就好了,具体逻辑 这里不关心

type Node struct {
	Key    string
	Weight int
}

func NewNode(key string, weight int) *Node {
	return &Node{
		Key:       key,
		Weight:   weight,
	}
}

type Consistent struct {
	Nodes     map[uint32]Node
	numReps   int
	Resources map[string]bool
	ring      HashRing
	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent{
		Nodes:     make(map[uint32]Node),
		numReps:   DEFAULT_REPLICAS,
		Resources: make(map[string]bool),
		ring:      HashRing{},
	}
}

func (c *Consistent) Add(node *Node) bool {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Key]; ok {
		return false
	}

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		c.Nodes[c.hashStr(str)] = *(node)
	}
	c.Resources[node.Key] = true
	c.sortHashRing()
	return true
}

func (c *Consistent) sortHashRing() {
	c.ring = HashRing{}
	for k := range c.Nodes {
		c.ring = append(c.ring, k)
	}
	sort.Sort(c.ring)
}

func (c *Consistent) joinStr(i int, node *Node) string {
	return node.Key + "*" + strconv.Itoa(node.Weight) +
	"-" + strconv.Itoa(i)
}

// MurMurHash算法 :https://github.com/spaolacci/murmur3
func (c *Consistent) hashStr(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) Get(key string) Node {
	c.RLock()
	defer c.RUnlock()

	hash := c.hashStr(key)
	i := c.search(hash)

	return c.Nodes[c.ring[i]]
}

func (c *Consistent) search(hash uint32) int {

	i := sort.Search(len(c.ring), func(i int) bool { return c.ring[i] >= hash })
	if i < len(c.ring) {
		if i == len(c.ring) - 1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1
	}
}

func (c *Consistent) Remove(node *Node) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Key]; !ok {
		return
	}

	delete(c.Resources, node.Key)

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		delete(c.Nodes, c.hashStr(str))
	}
	c.sortHashRing()
}

func main() {

	cHashRing := NewConsistent()

	for i := 0; i < 10; i++ {
		si := fmt.Sprintf("%d", i)
		cHashRing.Add(NewNode("node-" + si, 10))
	}
	//权重高
	cHashRing.Add(NewNode("node_ok", 20))
	cHashRing.Add(NewNode("node_00", 1))

	for k, v := range cHashRing.Nodes {
		fmt.Println("Hash:", k, " KEY:", v.Key)
	}

	ipMap := make(map[string]int, 0)
	for i := 0; i < 10000; i++ {
		si := fmt.Sprintf("key%d", i)
		k := cHashRing.Get(si)
		if _, ok := ipMap[k.Key]; ok {
			ipMap[k.Key] += 1
		} else {
			ipMap[k.Key] = 1
		}
	}

	for k, v := range ipMap {
		fmt.Println("Node Key:", k, " count:", v)
	}

}

/*
分布:

Node Key: node_ok  count: 1977
Node Key: node-5  count: 812
Node Key: node-1  count: 673
Node Key: node-3  count: 982
Node Key: node-2  count: 642
Node Key: node-7  count: 1044
Node Key: node-4  count: 846
Node Key: node-9  count: 942
Node Key: node_00  count: 65
Node Key: node-0  count: 678
Node Key: node-6  count: 560
Node Key: node-8  count: 779

*/