package easysort

import (
	"sort"
)

type sortableSlice struct {
	len  int
	swap func(int, int)
	less func(int, int) bool
}

func (s *sortableSlice) Len() int {
	return s.len
}

func (s *sortableSlice) Swap(i, j int) {
	s.swap(i, j)
}

func (s *sortableSlice) Less(i, j int) bool {
	return s.less(i, j)
}

func OrderBy(len int, swap func(i, j int), less func(i, j int) bool) {
	data := sortableSlice{
		len:  len,
		swap: swap,
		less: less,
	}

	sort.Sort(&data)
}

func IndexBy(len int, less func(i, j int) bool) []int {
	index := make([]int, len)
	for i := 0; i < len; i++ {
		index[i] = i
	}

	data := sortableSlice{
		len: len,
		swap: func(i, j int) {
			index[i], index[j] = index[j], index[i]
		},
		less: func(i, j int) bool {
			return less(index[i], index[j])
		},
	}

	sort.Sort(&data)

	return index
}
