package loophash


import (
	"testing"
	"fmt"
)

func Test_Run(t *testing.T) {
	h := NewHashNode()

	for a := 0; a < 10; a++{
		name:=fmt.Sprintf("node%d" ,a)
		println("add node:",name)
		h.AddNode(name)
	}

	node_stats :=make(map[string]int)

	for a := 0; a < 1000; a++{
		name := h.FindHashNode(fmt.Sprintf("abc-%d" ,a))
		_,ok:=node_stats[name]
		if !ok{
			node_stats[name]=0
		}
		node_stats[name]++

	}
	fmt.Printf("stats:%v",node_stats)
}