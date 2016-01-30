package loophash

/**
循环 一致性hash 简单的实现
 */
import (
	"hash/crc32"
//"sort"
	"fmt"
	"strings"
)
var virtual_node int = 18

type HashNode struct {

	nodes          map[string]int32 //节点列表
	nodes_id       map[int32]string //节点列表
	pos            int32            //id 计数器
	nodes_id_array *SortInt32Array  //id 列表
}

func NewHashNode() *HashNode {
	h := &HashNode{
		pos:0,
		nodes: make(map[string]int32),
		nodes_id: make(map[int32]string),
		nodes_id_array: NewSortInt32Array(),
	}
	return h
}


func (h * HashNode)AddNode(name string) int32 {
	var last_node int32
	for i := 0; i < virtual_node; i++ {
		last_node = h._AddNode(fmt.Sprintf("%s-%d", name, i))
	}
	return last_node
}
func (h * HashNode)_AddNode(name string) int32 {
	var node_id int32
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

func (h * HashNode)RemoveNode(name string) int32 {
	var last_node int32
	for i := 0; i < virtual_node; i++ {
		last_node = h.RemoveNode(fmt.Sprintf("%s-%d", name, i))
	}
	return last_node
}

func (h * HashNode)_RemoveNode(name string) int32 {
	node_id, ok := h.nodes[name]
	if !ok {
		return -1
	}
	h.nodes_id_array.Remove(node_id)
	delete(h.nodes, name)
	delete(h.nodes_id, node_id)
	return node_id
}

func (h * HashNode)getHashNodeId(str string) int32 {
	crc := crc32.New(crc32.IEEETable)
	crc.Write([]byte(str))
	id2 := int32(crc.Sum32())
	return id2
}
func (h * HashNode)FindHashNode(str string) string {


	crc := crc32.New(crc32.IEEETable)
	crc.Write([]byte(str))
	id2 := int32(crc.Sum32())

	index := h.nodes_id_array.FindNextValue(id2)
	name:= h.nodes_id[index]

	//虚拟节点的问题
	names := strings.Split(name,"-")
	name=strings.Join(names[0:len(names)-1],"-")
	return name
}

