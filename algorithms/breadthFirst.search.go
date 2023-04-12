package algorithms

import "quickDsp/dataStructures"

func BFSSearch(tree *dataStructures.Tree, toFind int) *dataStructures.TreeNode {
	if tree == nil || tree.Root == nil {
		return nil
	}
	helpQueue := dataStructures.GetTreeNodeQueue(2000)
	helpQueue.Push(tree.Root)

	currNode := helpQueue.Pop()
	for currNode != nil {
		if currNode.Data == toFind {
			return currNode
		}
		if rightNode := currNode.Right; rightNode != nil {
			helpQueue.Push(rightNode)
		}
		if leftNode := currNode.Left; leftNode != nil {
			helpQueue.Push(leftNode)
		}
		currNode = helpQueue.Pop()
	}
	return nil
}
