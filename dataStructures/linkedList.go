package dataStructures

import "fmt"

type LinkedList struct {
	Head *LsNode
	Tail *LsNode
	Len  int
}

type LsNode struct {
	Data int
	Next *LsNode
	Prev *LsNode // exist in doubly-linked list
}

func (l *LinkedList) Insert(inData int) {
	cNode := &LsNode{Data: inData}
	if l.Head == nil {
		l.Head = cNode
		l.Tail = cNode
	} else {
		l.Tail.Next = cNode
		l.Tail = cNode
	}
	l.Len++
}

func (l *LinkedList) Display() {
	if currNode := l.Head; currNode != nil {
		for currNode != nil {
			fmt.Printf("val: %v", currNode.Data)
			currNode = currNode.Next
		}
	}
}

// o(1) if cuz we have reference to the lsNode (and its Prev and Next)
func (l *LinkedList) Delete(node *LsNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node = nil
}

//o(n)
func (l *LinkedList) FindNodeByData(data int) *LsNode {
	if currNode := l.Head; currNode != nil {
		for currNode != nil {
			if currNode.Data == data {
				fmt.Print("found lsnode")
				return currNode
			}
			currNode = currNode.Next
		}
	}
	fmt.Print("lsnode not found")
	return nil
}

//o(n) since we dont have reference to the lsNode we need to delete and to go over the Data till find
func (l *LinkedList) DeleteData(dataToDelete int) {
	nodeToDelete := l.FindNodeByData(dataToDelete)
	l.Delete(nodeToDelete)
}
