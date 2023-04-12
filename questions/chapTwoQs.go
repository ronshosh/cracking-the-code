package questions

import (
	"log"
	"quickDsp/dataStructures"
)

//type lsNode struct {
//	data int
//	next *lsNode
//	prev *lsNode // exist in doubly-linked list
//}
//
//type LinkedList struct {
//	head *lsNode
//	tail *lsNode
//	len  int
//}

func questionOne(list *dataStructures.LinkedList) {
	if list.Head == nil || list.Head.Next == nil {
		return
	}

	var currentNode, nodeForComparison *dataStructures.LsNode
	currentNode = list.Head
	nodeForComparison = list.Head.Next

	for currentNode != nil {
		nodeForComparison = currentNode.Next
		for nodeForComparison != nil {
			if currentNode.Data == nodeForComparison.Data {
				currentNode.Next = nodeForComparison.Next
				nodeForComparison.Next.Prev = currentNode
				nodeForComparison = currentNode.Next
			} else {
				nodeForComparison = nodeForComparison.Next
			}
		}
		currentNode = currentNode.Next
	}
}

func questionTwo(head *dataStructures.LsNode, n int) *dataStructures.LsNode {
	if head == nil || n < 1 {
		return nil
	}
	p1, p2 := head, head
	for i := 0; i < (n - 1); i++ { // move p2 forward until creating n size gap between p1 and p2
		if p2 == nil {
			return nil // if linkedList length is smaller than n return null
		}
		p2 = p2.Next
	}
	for p2.Next != nil { // move both p1 p2 forward while keeping n size gap between them, until p2 is in the end of the list
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

func questionThree(node *dataStructures.LsNode) {
	if node == nil || node.Next == nil {
		return
	}
	node.Data = node.Next.Data
	node.Next = node.Next.Next
}

func QuestionFour(numOne, numTwo *dataStructures.LinkedList) *dataStructures.LinkedList {
	current1 := numOne.Head
	current2 := numTwo.Head
	if current1 == nil && current2 == nil {
		return nil
	}
	var result []int
	var addLater int
	for current1 != nil || current2 != nil {
		sumRes := 0
		if current1 != nil {
			sumRes += current1.Data
		}
		if current2 != nil {
			sumRes += current2.Data
		}
		log.Printf("curr sumRes: %+v", sumRes)
		log.Printf("curr addLater: %+v", addLater)
		sumRes += addLater
		newNodeData := sumRes % 10
		addLater = sumRes / 10
		log.Printf("curr newNodeData: %+v", newNodeData)
		result = append(result, newNodeData)

		if current1 != nil {
			if current1.Next != nil {
				current1 = current1.Next
			} else {
				current1 = nil
			}
		}

		if current2 != nil {
			if current2.Next != nil {
				current2 = current2.Next
			} else {
				current2 = nil
			}
		}
	}
	if addLater == 1 {
		result = append(result, addLater)
	}
	log.Printf("RESULT: %+v", result)
	return nil
}

func QuestionFive(head *dataStructures.LsNode) *dataStructures.LsNode {
	if head == nil || head.Next == nil {
		return nil
	}
	nextPointerMap := make(map[*dataStructures.LsNode]struct{})
	currentNode := head
	for currentNode.Next != nil {
		if _, k := nextPointerMap[currentNode]; k {
			return currentNode
		} else {
			nextPointerMap[currentNode] = struct{}{}
		}
		currentNode = currentNode.Next
	}
	return nil
}
