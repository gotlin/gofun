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

func LevelOrder(root *TreeNode) [][]int {


	var toTra []*TreeNode
	var isEnd = false
	var res [][]int

	if nil == root{
		return res
	}

	toTra = append(toTra, root)

	for ; !isEnd; {
		var level []int
		var toTraTmp []*TreeNode

		for _, node := range toTra {
			level = append(level, node.Val)

			if nil != node.Left {
				toTraTmp = append(toTraTmp, node.Left)
			}
			if nil != node.Right {

				toTraTmp = append(toTraTmp, node.Right)
			}
		}

		res = append(res, level)

		if len(toTraTmp) > 0 {
			toTra = toTraTmp
			isEnd = false
		} else {
			isEnd = true
		}

	}

	return res

}
