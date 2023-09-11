package main

import (
	"fmt"
	"go-problems/bst"
	"go-problems/sorting"
)

func main() {
	input := []int{4, 1, 7, 3, 9, 33, 13, 54}
	output := sorting.MergeSort(input)
	fmt.Println("After merge sort :", output)
	output = sorting.QuickSort(input)
	fmt.Println("After quick sort :", output)
	output = sorting.BubbleSort(input)
	fmt.Println("After bubble sort :", output)
	output = sorting.InsertionSort(input)
	fmt.Println("After Insertion sort :", output)
	output = sorting.SelectionSort(input)
	fmt.Println("After selection sort :", output)
	BinaryTree()
}

func BinaryTree() {
	tree := bst.NewBinaryTree(30)
	tree.Insert(33)
	tree.Insert(32)
	tree.Insert(35)
	tree.Insert(25)
	tree.Insert(28)
	tree.Insert(23)
	tree.Insert(12)
	tree.Insert(50)
	tree.Insert(9)
	tree.InOrder()
	fmt.Println()
	tree.Delete(25)
	tree.PreOrder()
	fmt.Println()
	tree.PostOrder()
	fmt.Println()
	fmt.Println("Search result : ", tree.Search(12))
	noOfTimesToStretchNodes := 3
	bst.Stretch(tree, noOfTimesToStretchNodes)
	fmt.Printf("Tree after stretching nodes by %d is : ", noOfTimesToStretchNodes)
	tree.InOrder()
}
