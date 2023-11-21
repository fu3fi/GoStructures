package main

import (
	"fmt"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (bt *BinaryTree) add(data int) {
	if bt.root == nil {
		bt.root = &BinaryNode{
			left:  nil,
			right: nil,
			data:  data,
		}
		// fmt.Printf("%d - %p\n", data, bt.root)
	} else {
		bt.root.add(data)
	}
}

func (bn *BinaryNode) add(data int) {
	if data < bn.data {
		if bn.left == nil {
			bn.left = &BinaryNode{
				left:  nil,
				right: nil,
				data:  data,
			}
		} else {
			bn.left.add(data)
		}
	} else if data > bn.data {
		if bn.right == nil {
			bn.right = &BinaryNode{
				left:  nil,
				right: nil,
				data:  data,
			}
		} else {
			bn.right.add(data)
		}
	}
}

func minRightSubTreeNode(subTree *BinaryNode) *BinaryNode {
	tmpNode := subTree
	for tmpNode.left != nil {
		tmpNode = tmpNode.left
	}
	return tmpNode
}

func (bt *BinaryTree) delete(data int) {
	bt.root.delete(data)
}

func (bt *BinaryNode) delete(data int) *BinaryNode {

	if bt.data == data {
		if bt.left == nil && bt.right == nil {
			bt = nil
			return bt
		}
		if bt.left == nil && bt.right != nil {
			bt = bt.right
			return bt
		}
		if bt.left != nil && bt.right == nil {
			bt = bt.left
			return bt
		}

		tmpNode := minRightSubTreeNode(bt.right)
		bt.data = tmpNode.data
		bt.right = bt.right.delete(bt.data)
		return bt
	}

	if bt.data > data {
		bt.left = bt.left.delete(data)
	}

	if bt.data < data {
		bt.right = bt.right.delete(data)
	}

	return bt
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
	fmt.Fprintf(os.Stdout, "%c:%v\n", branch, bt.data)
	bt.left.print(ns+2, 'L')
	bt.right.print(ns+2, 'R')
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
	fmt.Println(tree.search(-20))
	fmt.Println(tree.search(-19))
	// tree.print()

}

func (bn *BinaryNode) search(data int) *BinaryNode {
	if data == bn.data {
		return bn
	}
	if data < bn.data {
		if bn.left == nil {
			return nil
		} else {
			return bn.left.search(data)
		}
	} else if data > bn.data {
		if bn.right == nil {
			return nil
		} else {
			return bn.right.search(data)
		}
	}

	return nil
}

func (bt *BinaryTree) search(data int) *BinaryNode {
	return bt.root.search(data)
}
