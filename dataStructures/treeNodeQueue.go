package dataStructures

type treeNodeQueue struct {
	queue       []*TreeNode
	currentSize int
	maxSize     int
}

func GetTreeNodeQueue(size int) *treeNodeQueue {
	toReturn := &treeNodeQueue{
		queue:   make([]*TreeNode, 0, size),
		maxSize: size,
	}
	return toReturn
}

func (qu *treeNodeQueue) Push(val *TreeNode) {
	if qu.currentSize == qu.maxSize {
		return
	}
	qu.queue = append(qu.queue, val)
	qu.currentSize += 1
}

func (qu *treeNodeQueue) Pop() *TreeNode {
	if qu.currentSize == 0 {
		return nil
	}
	qu.currentSize -= 1
	toReturn := qu.queue[0]
	qu.queue = qu.queue[1:]
	return toReturn
}
