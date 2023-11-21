package main

import (
	"os"
	"fmt"
)

type node struct {
    value int
    left  *node
    right *node
}

type BinaryTree struct {
    root *node
}

func (bt *BinaryTree) add(value int) {
    if bt.root == nil {
        bt.root = &node{value: value}
        return
    }

    currentNode := bt.root
    for {
        if value < currentNode.value {
            if currentNode.left == nil {
                currentNode.left = &node{value: value}
                return
            }
            currentNode = currentNode.left
        } else {
            if currentNode.right == nil {
                currentNode.right = &node{value: value}
                return
            }
            currentNode = currentNode.right
        }
    }
}

func (bt *BinaryTree) search(value int) *node {
    currentNode := bt.root
    for currentNode != nil {
        if value == currentNode.value {
            return currentNode
        } else if value < currentNode.value {
            currentNode = currentNode.left
        } else {
            currentNode = currentNode.right
        }
    }
    return nil
}

func (bt *BinaryTree) print() {
    bt.root.print(0, 'S')
}

func (bt *node) print(ns int, branch rune) {
    if bt == nil {
        return
    }

    for i := 0; i < ns; i++ {
        fmt.Fprint(os.Stdout, " ")
    }
    // fmt.Printf("p %p\n", bt)
    fmt.Fprintf(os.Stdout, "%c:%v\n", branch, bt.value)
    bt.left.print(ns+2, 'L')
    bt.right.print(ns+2, 'R')
}

func (tree *BinaryTree) delete(value int) {
	tree.root = tree.root.delete(value)
}

func (root *node) delete(value int) *node {
    if root == nil {
        return nil
    }

    if value < root.value {
        root.left = root.left.delete(value)
    } else if value > root.value {
        root.right = root.right.delete(value)
    } else {
        // If the node has one or no children
        if root.left == nil {
            return root.right
        } else if root.right == nil {
            return root.left
        }

        // If the node has two children, find the minimum value in the right subtree
        minRight := root.right
        for minRight.left != nil {
            minRight = minRight.left
        }

        // Replace the value of the node to be deleted with the minimum value from the right subtree
        root.value = minRight.value

        // Delete the node with the minimum value from the right subtree
        root.right = root.right.delete(minRight.value)
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
	tree.delete(100)
	tree.print()
}