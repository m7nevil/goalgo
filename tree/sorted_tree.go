package tree

type SortedTree struct {
	Root *Node
}

func (tree *SortedTree) Add(values ...int) {
	for _, value := range values {
		tree.Insert(value)
	}
}

func (tree *SortedTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Data: value}
		return
	}

	node := tree.Root
	for {
		pValue, _ := node.Data.(int)
		if value < pValue {
			if node.Left == nil {
				node.Left = &Node{Data: value}
				return
			}
			node = node.Left
		} else if value > pValue {
			if node.Right == nil {
				node.Right = &Node{Data: value}
				return
			}
			node = node.Right
		}
		if node == nil {
			return
		}
	}
}

func (tree *SortedTree) Find(value int) *Node {
	node := tree.Root
	for {
		pValue, _ := node.Data.(int)
		if value < pValue {
			node = node.Left
		} else if value > pValue {
			node = node.Right
		} else {
			return node
		}

		if node == nil {
			return nil
		}
	}

	return nil
}
