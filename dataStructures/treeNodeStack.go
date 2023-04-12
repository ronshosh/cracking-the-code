package dataStructures

type TreeNodeStack struct {
	StackArray   []*TreeNode
	StackMaxSize int
	CurrentSize  int
}

func GetNewTreeNodeStack(size int) *TreeNodeStack {
	return &TreeNodeStack{
		StackArray:   make([]*TreeNode, 0, size),
		StackMaxSize: size,
		CurrentSize:  0,
	}
}

func (tns *TreeNodeStack) Push(val *TreeNode) bool {
	if tns.CurrentSize == tns.StackMaxSize {
		return false
	}
	tns.StackArray = append(tns.StackArray, val)
	tns.CurrentSize += 1
	return true
}

func (tns *TreeNodeStack) Pop() *TreeNode {
	if tns.CurrentSize == 0 {
		return nil
	}
	output := tns.StackArray[tns.CurrentSize-1]
	tns.StackArray = tns.StackArray[:tns.CurrentSize-1]
	tns.CurrentSize -= 1
	return output
}

func (tns *TreeNodeStack) IsEmpty() bool {
	return tns.CurrentSize == 0
}
