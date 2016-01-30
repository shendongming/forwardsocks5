package loophash

import (

)

type SortInt32Array struct {

	data []int32
}

func NewSortInt32Array() *SortInt32Array {
	s := &SortInt32Array{data:make([]int32, 0)}
	return s
}

func (s *SortInt32Array) Insert(value int32) int {

	pos := s.search(value)
//	fmt.Printf("data :%v value:%v\n", (s.data), value)
//	fmt.Printf("search value:%v,pos:%v,len:%v\n", value, pos, len(s.data))
	s.insert_pos(pos, value)


	return pos
}


//删除元素
func (s *SortInt32Array) Remove(value int32) int {
	pos := s.search(value)
	if pos >= len(s.data) {
		return -1
	}
	if s.data[pos] != value {
		return -1
	}

	tmp := append([]int32{}, s.data[:pos]...)
	s.data = append(tmp, s.data[pos + 1:]...)

	return pos
}


func (s *SortInt32Array) insert_pos(pos int, value int32) {
	if pos > len(s.data) {
		s.data = append(s.data, value)
		return
	}
	//分解为2个部分
	tmp := append([]int32{}, s.data[pos:]...)

	s.data = append(s.data[0:pos], value)

	s.data = append(s.data, tmp...)


}

func (s *SortInt32Array) FindNextValue(value int32) int32   {
 	pos:=s.FindNextIndex(value)
	return s.data[pos]
}
//找到最近邻的一个元素
func (s *SortInt32Array) FindNextIndex(value int32) int   {

	//这里使用顺序查询
	last_pos:=-1
	for pos,item :=range s.data{
		if value<item{
			last_pos=pos-1
			break
		}
	}
	//区别是最后一个
	if last_pos == -1{
		last_pos=len(s.data)-1
	}
	return last_pos
}

//升序保存
func (s *SortInt32Array) search(value int32) int {
	last_pos:=-1
	for pos,item :=range s.data{
		if value<=item{
			last_pos=pos
			break
		}
	}
	//为了需要增加
	if last_pos == -1{
		last_pos=len(s.data)
	}
	return last_pos
}

//todo: 二分查找到合适的位置
func (s *SortInt32Array) search_quick(value int32) int {
	data_len := len(s.data)
	if data_len == 0 {
		return 0
	}
	start := 0
	end := data_len
	mid := (start + end) / 2
	for start >= 0  && end <= data_len {
		if value > s.data[mid] {
			start = mid + 1
		}else if value < s.data[mid] {
			end = mid
		}else {
			break
		}

		mid = (start + end) / 2
		if start == end {
			break
		}

	}
	return mid

}
