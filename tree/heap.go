package tree

import (
	"errors"
	"fmt"
)

type Heap struct {
	arr   []int
	cap   int
	count int
}

func NewHeap(cap int) *Heap {
	arr := make([]int, cap)
	return &Heap{arr: arr, cap: cap, count: 0}
}

func (heap *Heap) Add(arr ...int) error {
	for _, v := range arr {
		err := heap.Insert(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (heap *Heap) Insert(v int) error {
	if heap.count >= heap.cap {
		return errors.New("容量已达上限")
	}

	heap.count++
	arr := heap.arr
	//sliceSet(heap.arr, heap.count, v)
	arr[heap.count] = v
	i := heap.count
	for i/2 > 0 && arr[i/2] < arr[i] {
		swap(arr, i, i/2)
		i = i / 2
	}

	return nil
}

func (heap *Heap) HeapArr() []int {
	return heap.arr[1 : heap.count+1]
}

func (heap *Heap) ToString() string {
	return fmt.Sprint(heap.HeapArr())
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func sliceSet(arr []int, i, v int) {
	if len(arr) == 0 {
		arr = append(arr, v)
		return
	}
	rear := append([]int{}, arr[i+1:]...)
	arr = append(append(arr[:i], v), rear...)
}
