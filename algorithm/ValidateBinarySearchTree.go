package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var latest int = -1

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var flag = isValidBST(root.Left)
	if latest == -1 && root.Val > latest {
		latest = root.Val
		return flag && isValidBST(root.Right)
	}
	return false
}
