package datastructure

import (
	"sort"
)

type TopN[T Trier] struct {
	data Data[T]
}

func (t TopN[T]) GetData() Data[T] {
	return t.data
}

func MakeTopN[T Trier]() TopN[T] {
	return TopN[T]{make([]*T, 0)}
}

func MakeTopN1[T Trier](t *T) TopN[T] {
	return TopN[T]{data: []*T{t}}
}

func MakeTopN2[T Trier](data []*T) TopN[T] {
	sort.Sort(Data[T](data))
	data = data[:MinInt(len(data), DEFAULT_TOPN)]
	return TopN[T]{data}
}

type Data[T Trier] []*T

func (t Data[T]) Len() int {
	return len(t)
}

func (t Data[T]) Less(i int, j int) bool {
	return (*(t[i])).Less(*(t[j]))
}

func (t Data[T]) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *TopN[T]) Merge(t2 TopN[T]) {
	t.data = append(t.data, t2.data...)
	sort.Sort(t.data)
	t.data = t.data[:MinInt(DEFAULT_TOPN, len(t.data))]
}

func (t *TopN[T]) Sort() {
	sort.Sort(t.data)
}

func MinInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func overlap[T Trier](d1 Data[T], d2 Data[T]) bool {
	for i := 0; i < d1.Len(); i++ {
		for j := 0; j < d2.Len(); j++ {
			if (*(d1[i])).Compare(*(d2[j])) == Equal {
				return true
			}
		}
	}
	return false
}
