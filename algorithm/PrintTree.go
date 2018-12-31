package algorithm

import "fmt"

/**
*
 * 题目：从上往下打印出二叉树的每个节点，同层节点从左至右打印
 * 广度优先层次遍历，或者 层次遍历
*/
// 前序遍历
func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Printf(" %v ", root.Val)
	}
	if root.Left != nil {
		PreOrder(root.Left)
		//nodeList = append(nodeList, node.Left)
	}
	if root.Right != nil {
		PreOrder(root.Right)
		//nodeList = append(nodeList, node.Right)
	}

}

//中序遍历
func InOrder(root *TreeNode) {
	if root.Left != nil {
		InOrder(root.Left)
	}

	if root != nil {
		fmt.Print(root.Val, " ")
	}

	if root.Right != nil {
		InOrder(root.Right)
	}

}

// 后序遍历
func PostOrder(root *TreeNode) {
	if root.Left != nil {
		PostOrder(root.Left)
	}
	if root.Right != nil {
		PostOrder(root.Right)
	}
	fmt.Print(root.Val, " ")
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
func CreateTree2() TreeNode {
	var treeIs = `
		   1
	     /   \
	   3       2
      / \     /  \
	 7   6   5   4
    /
   0
    `
	fmt.Println("create a tree like\n", treeIs)
	root := TreeNode{
		1,
		&TreeNode{3, &TreeNode{7,
			&TreeNode{0, nil,
				nil}, nil},
			&TreeNode{6, nil, nil}},
		&TreeNode{2,
			&TreeNode{5, nil, nil},
			&TreeNode{4, nil, nil}}}
	return root
}

func TestCaseInorder() {
	//tree := CreateTree()
	tree := CreateTree2()
	fmt.Println("前序遍历")
	PreOrder(&tree)
}

func TestCasePreOrder() {
	tree := CreateTree2()
	fmt.Println("中序遍历")
	InOrder(&tree)
}

func TestCasePostOrder() {
	tree := CreateTree2()
	fmt.Println("后序遍历")
	PostOrder(&tree)
}
