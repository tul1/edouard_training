package main 

import "fmt"

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func (BinaryTree *BinaryTree) Insert(data int) {
	temp_pointer := BinaryTree.root
	var samevaluefound bool

	if temp_pointer == nil {
		BinaryTree.root = &TreeNode{value: data}
	} else {
		switch {
		case temp_pointer.value > data:
				if temp_pointer.left == nil {
					temp_pointer.left = &TreeNode{value: data, right: temp_pointer}
				} else {
					for temp_pointer.left != nil {
						temp_pointer = temp_pointer.left
						if temp_pointer.value == data {
							temp_pointer.value = data
							samevaluefound = true
						}
					}
					if samevaluefound == false {
					temp_pointer.left = &TreeNode{value: data, right: temp_pointer}
					}
				}
			
		case temp_pointer.value == data:
			temp_pointer.value = data
		case temp_pointer.value < data:
			if temp_pointer.right == nil {
				temp_pointer.right = &TreeNode{value: data, left: temp_pointer}
			} else {
				for temp_pointer.right != nil {
					temp_pointer = temp_pointer.right
					if temp_pointer.value == data {
						temp_pointer.value = data
						samevaluefound = true
					}
				}
				if samevaluefound == false {
				temp_pointer.right = &TreeNode{value: data, left: temp_pointer}
				}
			}
		}
	}	
	
}

func main() {
	binaryTree := BinaryTree{}
	binaryTree.Insert(0)
	binaryTree.Insert(-1)
	binaryTree.Insert(-1)
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(2)
	temp_pointer := binaryTree.root
	for temp_pointer.left != nil {
		temp_pointer = temp_pointer.left
	}
	fmt.Print("In-Order Traversal: ", temp_pointer.value)

	for temp_pointer.right != nil {
		temp_pointer = temp_pointer.right
		fmt.Print(" -> ",temp_pointer.value)
	}
}
