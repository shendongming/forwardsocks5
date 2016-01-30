package loophash

/**
循环 一致性hash 简单的实现
 */
import (
	"sync/atomic"
	"hash/crc32"
//"sort"
)


type Hash struct {

	nodes        map[string]int32 //节点列表
	nodes_id     map[int32]string //节点列表
	pos          int32            //id 计数器
	node_id_list []int32
}

func NewHash() *Hash {
	h := &Hash{
		pos:0,
		nodes: make(map[string]int32),
		nodes_id: make(map[int32]string),
	}
	return h
}


func (h * Hash)AddNode(name string) int32 {
	var node_id int32
	node_id, ok := h.nodes[name]
	if ok {
		return node_id
	}
	node_id = h.GetHashNode(name)
	println("get new node id:", name, node_id)
	h.nodes[name] = node_id
	h.nodes_id[node_id] = name

	pos := find_loop_insert_pos(h.node_id_list, node_id)
	println("find pos:", pos)
	h.node_id_list = append(h.node_id_list, node_id)

	println("add name:", name)
	for k, v := range h.node_id_list {
		println(k, v)
	}
	return h.nodes[name]
}

func (h * Hash)createId() int32 {
	newId := int32(atomic.AddInt32(&h.pos, 1))
	return newId
}
func (h * Hash)GetHashNode(str string) int32 {


	crc := crc32.New(crc32.IEEETable)
	hash_id, err := crc.Write([]byte(str))
	id2 := crc.Sum32()
	println("hi", hash_id, "err", err)
	println("id2", id2)
	return int32(id2)
}



func insert_silce(data_list []int32, pos int, value int32) {

}

func find_loop_insert_pos(data_list []int32, find_value int32) int {
	data_len := len(data_list)
	start := 0
	end := data_len
	mid := (start + end) / 2
	println("start:", start, "end:", end, "mid:", mid)
	for start >= 0  && end < data_len {
		if find_value > data_list[mid] {
			start = mid
		}else if find_value < data_list[mid] {
			end = mid
		}else {
			return mid
		}
		if start == end {
			return start
		}

	}
	return mid
}