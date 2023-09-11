package bst

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func NewBinaryTree(val int) *TreeNode {
	return &TreeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (root *TreeNode) Insert(val int) {
	if val < root.val {
		if root.left == nil {
			root.left = NewBinaryTree(val)
		} else {
			root.left.Insert(val)
		}
	} else {
		if root.right == nil {
			root.right = NewBinaryTree(val)
		} else {
			root.right.Insert(val)
		}
	}
}

func (root *TreeNode) Search(val int) bool {
	if root == nil {
		return false
	}
	if val < root.val {
		return root.left.Search(val)
	}
	if val > root.val {
		return root.right.Search(val)
	}
	return true
}

func (root *TreeNode) Delete(val int) *TreeNode {
	if root == nil {
		return root
	}
	if val < root.val {
		root.left = root.left.Delete(val)
	} else if val > root.val {
		root.right = root.right.Delete(val)
	} else {
		if root.val == val {
			if root.left != nil && root.right == nil {
				temp := root.left
				root = nil
				root = temp
				return root
			}
			if root.left == nil && root.right != nil {
				temp := root.right
				root = nil
				root = temp
				return root
			}
			if root.left == nil && root.right == nil {
				root = nil
				return root
			}

			minNode := findMin(root.right)
			root.val = minNode.val
			root.right = root.right.Delete(minNode.val)
		}
	}
	return root
}

func (root *TreeNode) PreOrder() {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.val)
	root.left.PreOrder()
	root.right.PreOrder()
}

func (root *TreeNode) InOrder() {
	if root == nil {
		return
	}

	root.left.InOrder()
	fmt.Printf("%d ", root.val)
	root.right.InOrder()
}

func (root *TreeNode) PostOrder() {
	if root == nil {
		return
	}

	root.left.PostOrder()
	root.right.PostOrder()
	fmt.Printf("%d ", root.val)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.left)
	right := maxDepth(root.right)

	return MaxInt(left, right) + 1
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	} else {
		var min = root
		for root != nil && root.left != nil {
			min = root.left
		}
		return min
	}
}
