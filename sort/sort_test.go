package sort

import (
	"log"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{4, 6, 8, 1, 3, 88, 45, 58, 3456, 76, 457, 79, 8, 9, 45, 677}
	result := Bubble(arr)
	log.Println(result)
}

func TestMerge1(t *testing.T) {
	arr1 := []int{1, 3, 4, 7}
	arr2 := []int{3, 5, 7, 9, 2}
	merge(arr1, arr2)
}

func TestSelect(t *testing.T) {

}

func TestQuick(t *testing.T) {

}

func TestMerge(t *testing.T) {

}
