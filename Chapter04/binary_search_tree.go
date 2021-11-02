package main

import (
	"fmt"
	"sync"
)

// TreeNode class
type TreeNode struct {
	key int
	value int
	leftNode *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock sync.RWMutex
}

// InsertElement method
func (tree *BinarySearchTree) InsertElement (key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

// insertTreeNode function
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	}
}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

// inOrderTraverseTree method
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}