package algorithms

import "quickDsp/dataStructures"

func DFSSearch(tree *dataStructures.Tree, toFind int) *dataStructures.TreeNode {
	if tree == nil || tree.Root == nil {
		return nil
	}
	helpStack := dataStructures.GetNewTreeNodeStack(2000)
	helpStack.Push(tree.Root)

	currNode := helpStack.Pop()
	for currNode != nil {
		if currNode.Data == toFind {
			return currNode
		}
		if rightNode := currNode.Right; rightNode != nil {
			helpStack.Push(rightNode)
		}
		if leftNode := currNode.Left; leftNode != nil {
			helpStack.Push(leftNode)
		}
		currNode = helpStack.Pop()
	}
	return nil
}
