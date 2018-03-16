package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {

	var res []int

	if root == nil {
		return res
	}

	res = append(res, root.Val)

	left := root.Left
	if left != nil {
		res = append(res, PreorderTraversal(left)...)
	}

	right := root.Right
	if right != nil {
		res = append(res, PreorderTraversal(right)...)
	}

	return res

}

func InorderTraversal(root *TreeNode) []int {

	var res []int

	if root == nil {
		return res
	}

	left := root.Left
	if left != nil {
		res = append(res, InorderTraversal(left)...)
	}

	res = append(res, root.Val)

	right := root.Right
	if right != nil {
		res = append(res, InorderTraversal(right)...)
	}

	return res
}



func PostorderTraversal(root *TreeNode) []int {

	var res []int

	if root == nil {
		return res
	}

	left := root.Left
	if left != nil {
		res = append(res, PostorderTraversal(left)...)
	}


	right := root.Right
	if right != nil {
		res = append(res, PostorderTraversal(right)...)
	}

	res = append(res, root.Val)

	return res
}