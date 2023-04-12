package dataStructures

import "math"

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Data  int
}

func GetTree() *Tree {
	return new(Tree)
}

func (t *Tree) InsertToBinarySearchTree(incomingData int) {
	if t.Root == nil {
		t.Root = &TreeNode{Data: incomingData}
	} else {
		t.Root.insertBinarySearchTree(incomingData)
	}
}

func (t *TreeNode) insertBinarySearchTree(incomingData int) {
	nodeData := t.Data
	if incomingData < nodeData {
		if t.Left == nil {
			t.Left = &TreeNode{Data: incomingData}
		} else {
			t.Left.insertBinarySearchTree(incomingData)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Data: incomingData}
		} else {
			t.Right.insertBinarySearchTree(incomingData)
		}
	}

}

///////////////////////////////////////////

func (t *Tree) InsertToBinaryTree(incomingData int) {
	if currNode := t.Root; currNode == nil {
		currNode = &TreeNode{Data: incomingData}
	} else {
		for currNode != nil {
			if currNode.Left == nil {
				currNode.Left = &TreeNode{Data: incomingData}
				return
			}

			if currNode.Right == nil {
				currNode.Right = &TreeNode{Data: incomingData}
				return
			}
			currNode = currNode.Left
		}
	}
}

///////////////////////////////////////////

func (t *Tree) SumTreeValues() int {
	if t.Root == nil {
		return 0
	}
	sum := 0
	return t.Root.sumDataValues(sum)
}

func (t *TreeNode) sumDataValues(sum int) int {
	sum += t.Data
	if t.Right != nil {
		return t.Right.sumDataValues(sum)
	}
	if t.Left != nil {
		return t.Left.sumDataValues(sum)
	}
	return sum
}

//////////////////////////////////////////

func (t *Tree) Find(toFind int) *TreeNode {
	if t.Root == nil {
		return nil
	}
	return t.Root.find(toFind)
}

func (t *TreeNode) find(toFind int) *TreeNode {
	if toFind == t.Data {
		return t
	}
	if toFind < t.Data && t.Left != nil {
		return t.Left.find(toFind)
	}
	if toFind > t.Data && t.Right != nil {
		return t.Right.find(toFind)
	}
	return nil
}

////////////////////////////////////////////

func (t *Tree) GetMaxDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Max(t.GetMaxDepth(root.Left), t.GetMaxDepth(root.Right))
}

func (t *Tree) GetMinDepth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Min(t.GetMinDepth(root.Left), t.GetMinDepth(root.Right))
}
