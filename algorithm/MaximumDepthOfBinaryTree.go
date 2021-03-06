package algorithm

/**
MaximumDepthOfBinaryTree
 https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

//Definition for a binary tree node.


func maxDepth(root *TreeNode) int {
	if root != nil {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if leftDepth > rightDepth {
			return 1 + leftDepth
		}
		return 1 + rightDepth
	}
	return 0
}


