package loophash


import (
	"testing"
)
func Test_Sort_Array_insert(t *testing.T) {
	var i int
	t.Log("start test Test_Sort_Array_insert ")
	h := NewSortInt32Array()
	i = h.Insert(2)
	if i != 0 {
		t.Error("insert fail")
	}
	i = h.Insert(6)
	if i != 1 {
		t.Error("insert fail")
	}
	i = h.Insert(3)
	if i != 1 {
		t.Error("insert fail")
	}
	i = h.Insert(4)
	if i != 2 {
		t.Error("insert fail")
	}
	i = h.Insert(5)
	if i != 3 {
		t.Error("insert fail")
	}
	i = h.Insert(7)
	if i != 5 {
		t.Error("insert fail")
	}
	t.Log("ok:", h)
}

func Test_Sort_Array_insert2(t *testing.T) {

	t.Log("start test Test_Sort_Array_insert ")
	h := NewSortInt32Array()
	test_cace := [][2]int{
		{5, 0},
		{4, 0},
		{3, 0},
		{2, 0},
		{1, 0},
	}
	for _, v := range test_cace {
		if v[1] != h.Insert(uint32(v[0])) {
			t.Error("test insert tail")
		}
	}
	t.Log("ok:", h)
}


func Test_Sort_Array_Find(t *testing.T) {
	h := NewSortInt32Array()
	test_cace := [][2]int{
		{66, 0},
		{44, 0},
		{33, 0},
		{22, 0},
		{11, 0},
	}
	for _, v := range test_cace {
		if v[1] != h.Insert(uint32(v[0])) {
			t.Error("test insert tail")
		}
	}
	test_cace2 := [][2]int{
		{12, 0},
		{6, 4},
		{23, 1},
		{88, 4},
	}
	for _, v := range test_cace2 {
		ret := h.FindNextIndex(uint32(v[0]))
		if v[1] != ret {
			t.Error("test find tail:", v[0], v[1], ret)
		}
	}

}
func Test_Sort_Array_remove(t *testing.T) {

	var i int
	t.Log("start test Test_Sort_Array_insert ")
	h := NewSortInt32Array()
	test_cace := [][2]int{
		{6, 0},
		{4, 0},
		{3, 0},
		{2, 0},
		{1, 0},
	}
	for _, v := range test_cace {
		if v[1] != h.Insert(uint32(v[0])) {
			t.Error("test insert tail")
		}
	}

	if -1 != h.Remove(9) {
		t.Error("find not exists value test fail")
	}
	if -1 != h.Remove(5) {
		t.Error("find not exists value test fail")
	}

	i = h.Remove(6)
	if i != 4 {
		t.Error("find end pos test fail")
	}
	i = h.Remove(6)
	if i != -1 {
		t.Error("find end pos test fail")
	}

	i = h.Remove(1)
	if i != 0 {
		t.Error("find first pos test fail ", i)
	}
	i = h.Remove(1)
	if i != -1 {
		t.Error("find first pos test fail", i)
	}

}
