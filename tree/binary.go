package tree

import (
	"log"
)

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func PreOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	log.Println(node.Data)
	PreOrderTraverse(node.Left)
	PreOrderTraverse(node.Right)
}

func MidOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	MidOrderTraverse(node.Left)
	log.Println(node.Data)
	MidOrderTraverse(node.Right)
}

func PostOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	PostOrderTraverse(node.Left)
	PostOrderTraverse(node.Right)
	log.Println(node.Data)
}
