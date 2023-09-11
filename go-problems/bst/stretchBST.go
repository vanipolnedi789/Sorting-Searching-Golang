package bst

import "fmt"

// type TreeNode struct {
// 	left  *TreeNode
// 	right *TreeNode
// 	val   int
// }

func (node *TreeNode) PrintTree() {
	if node.left != nil {
		node.left.PrintTree()
	}
	fmt.Println(node.val)
	if node.right != nil {
		node.right.PrintTree()
	}
}

func stretchChildren(root *TreeNode, k int) {
	if root == nil {
		return
	}
	stretchChildren(root.left, k)
	stretchChildren(root.right, k)

	if root.left != nil {
		stretchByKNodes(root.left, k, "left")
	}

	if root.right != nil {
		stretchByKNodes(root.right, k, "right")
	}
}

func Stretch(root *TreeNode, k int) {
	stretchChildren(root, k)
	stretchByKNodes(root, k, "right")
}

func stretchByKNodes(root *TreeNode, k int, direction string) *TreeNode {
	var treeRootLeft, treeRootRight *TreeNode
	if direction == "left" {
		treeRootLeft = root.left
	} else {
		treeRootRight = root.right
	}
	stretchVal := root.val / k
	root.val = stretchVal
	for i := 0; i < k-1; i++ {
		if direction == "left" {
			root.left = &TreeNode{val: stretchVal}
			root = root.left
		}
		if direction == "right" {
			root.right = &TreeNode{val: stretchVal}
			root = root.right
		}
	}

	if direction == "left" {
		root.left = treeRootLeft
	} else {
		root.right = treeRootRight
	}
	return root
}
