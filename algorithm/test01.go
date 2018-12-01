package algorithm

import "fmt"

func T1() {
	nodeTrue := TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{3, nil, nil}},
	}

	nodeFalse := TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{2, &TreeNode{3, &TreeNode{6, nil, nil}, nil}, &TreeNode{3, nil, nil}},
	}

	nodeFalse2 := TreeNode{6,
		&TreeNode{1, nil, &TreeNode{4, nil, nil}},
		&TreeNode{6, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(isSymmetric(&nodeTrue))
	fmt.Println(isSymmetric(&nodeFalse))
	fmt.Println(isSymmetric(&nodeFalse2))
}

func T2() {
	// create a test tree
	var treeIs = `
		 3
	   / \
	  9    20
	     /   \
	    15     7
    `
	root := TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}
	depth := maxDepth(&root)
	fmt.Println("the tree shape is ")
	fmt.Println(treeIs)
	fmt.Printf(" and tree depth is %d ", depth)
}
