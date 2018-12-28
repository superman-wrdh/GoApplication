package algorithm

import "fmt"

/**
*
 * 题目：从上往下打印出二叉树的每个节点，同层节点从左至右打印
 * 广度优先层次遍历，
*/

func PrintBF(root *TreeNode) {
	if root != nil {
		fmt.Printf(" %v ", root.Val)
	}
	if root.Left != nil {
		PrintBF(root.Left)
		//nodeList = append(nodeList, node.Left)
	}
	if root.Right != nil {
		PrintBF(root.Right)
		//nodeList = append(nodeList, node.Right)
	}

}

func CreateTree() TreeNode {
	var treeIs = `
		 3
	   / \
	  9    20
	     /   \
	    15     7
    `
	fmt.Println("create a tree like\n", treeIs)
	root := TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}
	return root
}

func TestCaseBf() {
	tree := CreateTree()
	PrintBF(&tree)

}
