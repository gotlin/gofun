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

	if nil == root {
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

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDp := 1 + maxDepth(root.Left)
	rightDp := 1 + maxDepth(root.Right)

	if leftDp > rightDp {
		return leftDp
	} else {
		return rightDp
	}
}

func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return is(root.Left, root.Right)
}

func is(left, right *TreeNode) bool {

	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false

	}

	return is(left.Left, right.Right) && is(left.Right, right.Left)
}

func AsPathSum(root *TreeNode, sum int) bool {

	var fun func(node *TreeNode, tem int) bool
	fun = func(node *TreeNode, tem int) bool {

		if node == nil {
			return false
		}

		if tem+node.Val == sum && node.Left == nil && node.Right == nil {
			return true
		}

		return fun(node.Left, tem+node.Val) || fun(node.Right, tem+node.Val)
	}

	return fun(root, 0)
}
