package questions

import (
	"log"
	"quickDsp/dataStructures"
)

func qOne() {
	//base := make([]int, 300)
	// 0 - 99
	// 100 - 199
	// 200 - 299

}

type stackSharedArr struct {
	maxIndex    int
	minIndex    int
	currentSize int
	sizeLimit   int
	sharedArray []int
}

func newStackSharedArr(arr []int, maxIndex, minIndex int) *stackSharedArr {
	stack := &stackSharedArr{
		sharedArray: arr,
		maxIndex:    maxIndex,
		minIndex:    minIndex,
		sizeLimit:   maxIndex - minIndex + 1,
	}
	return stack
}

func (sk *stackSharedArr) push(val int) {
	if sk.currentSize == sk.sizeLimit {
		return
	}
	sk.sharedArray[sk.minIndex+sk.currentSize] = val
	sk.currentSize += 1
}

func (sk *stackSharedArr) pull() int {
	if sk.currentSize == 0 {
		return 0
	}
	output := sk.sharedArray[sk.minIndex+sk.currentSize-1]
	sk.sharedArray[sk.minIndex+sk.currentSize-1] = 0
	sk.currentSize -= 1
	return output
}

//_____________________________________________________________________________________________

func qTwo() {

}

type stackQ2 struct {
	currentSize  int
	stackMaxSize int
	stackArray   []int
	minVal       *int
}

func (sk *stackQ2) push(val int) {
	if sk.currentSize == sk.stackMaxSize {
		return
	}
	sk.stackArray = append(sk.stackArray, val)
	sk.currentSize += 1

	if sk.minVal == nil {
		sk.minVal = new(int)
		*sk.minVal = val
	} else {
		if val < *sk.minVal {
			*sk.minVal = val
		}
	}
}

//__________________________________________________________________________________________________
func qThree() {}

type SetOfStacks struct {
	StackMaxSize int

	ArraysStack     []*dataStructures.Stack
	ArraysStackSize int
}

func NewSetOfStacks(maxSize int) *SetOfStacks {
	output := &SetOfStacks{
		StackMaxSize: maxSize,
		ArraysStack:  make([]*dataStructures.Stack, 0, 50),
	}
	return output
}

func (ss *SetOfStacks) pushStack(val *dataStructures.Stack) {
	ss.ArraysStack = append(ss.ArraysStack, val)
	ss.ArraysStackSize += 1
}

func (ss *SetOfStacks) pullStack() *dataStructures.Stack {
	output := ss.ArraysStack[ss.ArraysStackSize-1]
	ss.ArraysStack = ss.ArraysStack[:ss.ArraysStackSize-1]
	ss.ArraysStackSize -= 1
	return output
}

func (ss *SetOfStacks) PushVal(val int) {
	if ss.ArraysStackSize == 0 {
		ss.pushStack(dataStructures.GetNewStack(ss.StackMaxSize))
	}

	currStack := ss.pullStack()
	if !currStack.Push(val) {
		ss.pushStack(currStack)
		newStack := dataStructures.GetNewStack(ss.StackMaxSize)
		newStack.Push(val)
		ss.pushStack(newStack)
	} else {
		ss.pushStack(currStack)
	}
}

func (ss *SetOfStacks) GetVal() int {
	if ss.ArraysStackSize == 0 {
		log.Print("no stacks")
		return 0
	}
	currStack := ss.pullStack()
	ok, res := currStack.Pop()
	if !ok {
		log.Print("no values")
		return 0
	}
	if currStack.CurrentSize > 0 {
		ss.pushStack(currStack)
	}
	return res
}

//_________________________________________________________________

func qFive() {
}

type QueueOfStacks struct {
	stack1 *dataStructures.Stack
	stack2 *dataStructures.Stack
}

func GetNewQueueOfStacks() *QueueOfStacks {
	size := 20
	return &QueueOfStacks{
		stack1: dataStructures.GetNewStack(size / 2),
		stack2: dataStructures.GetNewStack(size / 2),
	}
}

func (qs *QueueOfStacks) PushVal(val int) {
	if !qs.stack1.Push(val) {
		log.Print("queue is full")
	}
}

func (qs *QueueOfStacks) GetVal() int {
	if qs.stack2.CurrentSize > 0 {
		ok, val := qs.stack2.Pop()
		if ok {
			return val
		}
	}
	if qs.stack1.CurrentSize == 0 {
		return 00 // no values in queue
	}
	for i := qs.stack1.CurrentSize; i > 0; i-- {
		_, val := qs.stack1.Pop()
		if ok := qs.stack2.Push(val); !ok {
			break
		}
	}
	ok, val := qs.stack2.Pop()
	if ok {
		return val
	}
	return 00 // no values in queue
}

//_________________________________________________________________

func QSix(inStack *dataStructures.Stack) {
	if inStack.IsEmpty() || inStack.CurrentSize == 1 {
		return
	}
	resStack := dataStructures.GetNewStack(inStack.CurrentSize)
	storageStack := dataStructures.GetNewStack(inStack.CurrentSize)
	for !inStack.IsEmpty() {
		_, val := inStack.Pop()
		log.Printf("curr val: %+v", val)
		if resStack.IsEmpty() {
			resStack.Push(val)
			continue
		}
		_, topOrderedVal := resStack.Peek()
		log.Printf("topOrderedVal: %+v", topOrderedVal)
		for val > topOrderedVal && !resStack.IsEmpty() {
			_, orderedVal := resStack.Pop()
			storageStack.Push(orderedVal)
			_, topOrderedVal = resStack.Peek()
			log.Printf("topOrderedVal: %+v", topOrderedVal)
		}
		resStack.Push(val)
		for !storageStack.IsEmpty() {
			_, storedVal := storageStack.Pop()
			resStack.Push(storedVal)
		}
	}
	log.Printf("HERE RES: %+v", resStack.StackArray)
}
