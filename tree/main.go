package main

import (
	"fmt"
	"os"
)

type BinaryNode struct {
	leftChield  *BinaryNode
	rightChield *BinaryNode
	value       int
}

type BinaryTree struct {
	root *BinaryNode
}

// func NewTree(root int) *Node {
// 	node := &Node{}
// 	node.value = root
// 	return node
// }

func (bn *BinaryNode) add(value int) {
	if value < bn.value {
		if bn.leftChield == nil {
			bn.leftChield = &BinaryNode{
				leftChield:  nil,
				rightChield: nil,
				value:       value,
			}
		} else {
			bn.leftChield.add(value)
		}
	} else if value > bn.value {
		if bn.rightChield == nil {
			bn.rightChield = &BinaryNode{
				leftChield:  nil,
				rightChield: nil,
				value:       value,
			}
		} else {
			bn.rightChield.add(value)
		}
	}

}

func (bt *BinaryTree) add(value int) {
	if bt.root == nil {
		bt.root = &BinaryNode{
			leftChield:  nil,
			rightChield: nil,
			value:       value,
		}
	} else {
		bt.root.add(value)
	}
}

func (bt *BinaryTree) print() {
	bt.root.print(0, 'S')
}

func (bt *BinaryNode) print(ns int, branch rune) {
	if bt == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(os.Stdout, " ")
	}
	// fmt.Printf("p %p\n", bt)
	fmt.Fprintf(os.Stdout, "%c:%v\n", branch, bt.value)
	bt.leftChield.print(ns+2, 'L')
	bt.rightChield.print(ns+2, 'R')
}

func (bt *BinaryTree) delete(value int) {
	bt.root.delete(value)
}

func minRightSubTreeNode(subTree *BinaryNode) *BinaryNode {
	tmpNode := subTree
	for tmpNode.leftChield != nil {
		tmpNode = tmpNode.leftChield
	}
	return tmpNode
}

func (bn *BinaryNode) delete(value int) *BinaryNode {
	if bn.value < value {
		bn.rightChield = bn.rightChield.delete(value)
	} else if bn.value > value {
		bn.leftChield = bn.leftChield.delete(value)
	} else {
		if bn.leftChield == nil {
			return bn.rightChield
		} else if bn.rightChield == nil {
			return bn.leftChield
		}

		tmpNode := minRightSubTreeNode(bn.rightChield)
		bn.value = tmpNode.value
		bn.rightChield = bn.rightChield.delete(bn.value)
	}
	return bn
}

func (bt *BinaryTree) search(value int) *BinaryNode {
	return bt.root.search(value)
}

func (bn *BinaryNode) search(value int) *BinaryNode {
	if value == bn.value {
		return bn
	}
	if value < bn.value {
		if bn.leftChield == nil {
			return nil
		} else {
			return bn.leftChield.search(value)
		}
	} else if value > bn.value {
		if bn.rightChield == nil {
			return nil
		} else {
			return bn.rightChield.search(value)
		}
	}
	return nil
}

func NewTree(root int) *BinaryNode {
	node := &BinaryNode{}
	node.value = root
	return node
}

func deleteNode(root *BinaryNode, value int) *BinaryNode {
	if root == nil {
		return nil
	}
	if root.value == value {
		if root.leftChield == nil {
			return root.rightChield
		} else if root.rightChield == nil {
			return root.leftChield
		}
		left := root.leftChield
		for left.rightChield != nil {
			left = left.rightChield
		}
		left.rightChield = root.rightChield
		root = root.leftChield
	} else if root.value < value {
		root.rightChield = deleteNode(root.rightChield, value)
	} else if root.value > value {
		root.leftChield = deleteNode(root.leftChield, value)
	}
	return root
}

func main() {

	tree := &BinaryTree{}
	tree.add(100)
	tree.add(-20)
	tree.add(-50)
	tree.add(-15)
	tree.add(-17)
	tree.add(-60)
	tree.add(50)
	tree.add(60)
	tree.add(55)
	tree.add(85)
	tree.add(15)
	tree.add(5)
	tree.add(-10)

	tree.print()
	tree.delete(50)
	tree.print()
	fmt.Println(tree.search(13))

}
