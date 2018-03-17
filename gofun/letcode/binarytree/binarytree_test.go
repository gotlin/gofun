package binarytree

import (
	"testing"
	"fmt"
)

func TestPreorderTraversal(t *testing.T) {

	var tree *TreeNode
	PreorderTraversal(tree)
}

func TestLevelOrder(t *testing.T) {
	var tree = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	res := LevelOrder(tree)
	fmt.Println(res)
}

func TestIsSymmetric(t *testing.T) {
	var tree = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(IsSymmetric(tree))

}

func TestAsPathSum(t *testing.T) {
	var tree = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(AsPathSum(tree,3))

}
