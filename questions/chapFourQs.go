package questions

import "quickDsp/dataStructures"

//1

func isBalancedTree(tree *dataStructures.Tree) bool {
	if tree == nil || tree.Root == nil {
		return true
	}
	maxDepth := tree.GetMaxDepth(tree.Root)
	minDepth := tree.GetMinDepth(tree.Root)
	return maxDepth-minDepth == 1
}

//3

func createBinaryTreeMinimalHeight(incoming []int) *dataStructures.Tree {
	if len(incoming) == 0 {
		return nil
	}
	tree := dataStructures.GetTree()
	for _, val := range incoming {
		tree.InsertToBinaryTree(val)
	}
	return tree
}

// 4 HELP

func getLSsByTreeLevels(tree *dataStructures.Tree) {
	if tree == nil || tree.Root == nil {
		return
	}
	allTreeValsArray := []*dataStructures.TreeNode{}
	helpQueue := dataStructures.GetTreeNodeQueue(2000)
	helpQueue.Push(tree.Root)

	currNode := helpQueue.Pop()
	for currNode != nil {
		allTreeValsArray = append(allTreeValsArray, currNode)
		if rightNode := currNode.Right; rightNode != nil {
			helpQueue.Push(rightNode)
		}
		if leftNode := currNode.Left; leftNode != nil {
			helpQueue.Push(leftNode)
		}
		currNode = helpQueue.Pop()
	}

	//for _, node := range allTreeValsArray {
	//
	//}
	return
}
