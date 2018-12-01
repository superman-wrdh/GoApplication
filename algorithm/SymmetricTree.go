package algorithm

func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricHelp(root.Left, root.Right)
}

func isSymmetricHelp(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelp(left.Left, right.Right) && isSymmetricHelp(left.Right, right.Left)
}


