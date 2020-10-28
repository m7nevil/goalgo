package main

import (
	"github.com/m7nevil/goalgo/tree"
	"log"
)

func main() {
	//arr1 := []int{1,3,4,7}
	//arr2 := []int{3,5,6,7,9}
	//
	//sort.Merge(arr1, arr2)
	//log.Println(arr1, arr2)

	//arr := []int{4, 6, 8, 1, 3, 88, 45, 58, 3456, 76, 457, 79, 21, 9, 47, 677, 31}
	//
	//log.Println(arr[0:])
	//res := sort.Bubble(arr)
	//res := sort.Insertion(arr)
	//res := sort.Selection(arr)
	//res := sort.Merge(arr)
	//res := sort.Quick(arr)
	//log.Println(res)

	//node11 := tree.Node{Data:4}
	//node12 := tree.Node{Data:5}
	//node1 := tree.Node{Data:2, Left:&node11, Right:&node12}
	//node21 := tree.Node{Data:6}
	//node2 := tree.Node{Data:3, Left:&node21}
	//root := tree.Node{1, &node1, &node2}
	//
	//tree.PreOrderTraverse(&root)
	//log.Println("===============")
	//tree.MidOrderTraverse(&root)
	//log.Println("===============")
	//tree.PostOrderTraverse(&root)

	//sTree := tree.SortedTree{}
	//sTree.Add(5, 7, 3, 2, 4, 9, 10, 6)

	//tree.PreOrderTraverse(sTree.Root)
	//log.Println("===============")
	//tree.MidOrderTraverse(sTree.Root)
	//log.Println("===============")
	//tree.PostOrderTraverse(sTree.Root)

	//log.Println(sTree.Find(9))

	heap := tree.NewHeap(20)
	heap.Add(1, 4, 5, 76, 8, 9, 3, 45, 89, 43, 31)
	log.Println(heap.ToString())
}
