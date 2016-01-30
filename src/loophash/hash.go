package loophash

/**
循环 一致性hash 简单的实现
这个差距有点不均匀呀
stats:node6,44
stats:node7,67
stats:node5,163
stats:node8,26
stats:node1,105
stats:node0,132
stats:node3,107
stats:node2,110
stats:node9,97
stats:node4,149
 */
import (
	"hash/crc32"
//"sort"
	"fmt"
	"strings"
)
var virtual_node int = 18

type HashNode struct {

	nodes          map[string]uint32 //节点列表
	nodes_id       map[uint32]string //节点列表
	pos            uint32            //id 计数器
	nodes_id_array *SortInt32Array  //id 列表
}

func NewHashNode() *HashNode {
	h := &HashNode{
		pos:0,
		nodes: make(map[string]uint32),
		nodes_id: make(map[uint32]string),
		nodes_id_array: NewSortInt32Array(),
	}
	return h
}


func (h * HashNode)AddNode(name string) uint32 {
	var last_node uint32
	for i := 0; i < virtual_node; i++ {
		last_node = h._AddNode(fmt.Sprintf("%s-%d", name, i))
	}
	return last_node
}
func (h * HashNode)_AddNode(name string) uint32 {
	var node_id uint32
	node_id, ok := h.nodes[name]
	if ok {
		return node_id
	}

	//虚拟节点 8个
	node_id = h.getHashNodeId(name)
	println("get new node id:", name, node_id)
	h.nodes[name] = node_id
	h.nodes_id[node_id] = name
	h.nodes_id_array.Insert(node_id)

	return h.nodes[name]
}

func (h * HashNode)RemoveNode(name string) uint32 {
	var last_node uint32
	for i := 0; i < virtual_node; i++ {
		last_node = h.RemoveNode(fmt.Sprintf("%s-%d", name, i))
	}
	return last_node
}

func (h * HashNode)_RemoveNode(name string) uint32 {
	node_id, ok := h.nodes[name]
	if !ok {
		return 0
	}
	h.nodes_id_array.Remove(node_id)
	delete(h.nodes, name)
	delete(h.nodes_id, node_id)
	return node_id
}

func (h * HashNode)getHashNodeId(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))

}
func (h * HashNode)FindHashNode(str string) string {



	id2:=crc32.ChecksumIEEE([]byte(str))

	index := h.nodes_id_array.FindNextValue(id2)
	name:= h.nodes_id[index]

	//虚拟节点的问题
	names := strings.Split(name,"-")
	name=strings.Join(names[0:len(names)-1],"-")
	return name
}

