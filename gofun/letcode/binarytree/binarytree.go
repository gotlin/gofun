package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {

	var res []int

	if root == nil{
		return res
	}

	res = append(res,root.Val)

	left:=root.Left
	if left != nil{
		res=append(res,PreorderTraversal(left)...)
	}

	right:=root.Right
	if right != nil{
		res=append(res,PreorderTraversal(right)...)
	}

	return res

}
