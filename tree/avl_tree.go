package tree

type AvlNode struct {
	Data   interface{}
	Left   *AvlNode
	Right  *AvlNode
	bf     int
	parent *AvlNode
}

type AvlTree struct {
	Root *AvlNode
}

func (tree *AvlTree) Insert(value int) {
	tree.insertNode(value, tree.Root)
}

func (tree *AvlTree) insertNode(value int, node *AvlNode) {
	tree.Root = &AvlNode{Data: value, bf: 0}
	nValue, _ := tree.Root.Data.(int)
	if value < nValue {
		//TODO
	}
}

func (tree *AvlTree) leftRotate(node *AvlNode) {
	subNode := node.Right

	left := false
	if node.parent != nil {
		subNode.parent = node.parent
		if node.parent.Right == node {
			left = true
		}
	} else {
		subNode.parent = nil
	}
	node.Right = subNode.Left
	node.parent = subNode
	subNode.Left = node
	node = subNode
	if node.parent == nil {
		tree.Root = node
	} else {
		if left {
			node.parent.Left = node
		} else {
			node.parent.Right = node
		}
	}
}

func (tree *AvlTree) rightRotate(node *AvlNode) {
	subNode := node.Left

	left := false
	if node.parent != nil {
		subNode.parent = node.parent
		if node.parent.Left == node {
			left = true
		}
	} else {
		subNode.parent = nil
	}
	node.Left = subNode.Right
	node.parent = subNode
	subNode.Right = node
	node = subNode
	if node.parent == nil {
		tree.Root = node
	} else {
		if left {
			node.parent.Left = node
		} else {
			node.parent.Right = node
		}
	}
}

func (tree *AvlTree) rightBalance(node *AvlNode) {
	subNode := node.Right
	switch subNode.bf {
	case -1:
		node.bf = 0
		subNode.bf = 0
		tree.leftRotate(node) //TODO
		break
	case 1:
		node.bf = 1
		subNode.bf = 0
		break
	case 0:
		node.bf = 0
		subNode.bf = 0
		break

	}
}
