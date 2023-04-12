package dataStructures

type Stack struct {
	StackArray   []int
	StackMaxSize int
	CurrentSize  int
}

func GetNewStack(size int) *Stack {
	return &Stack{
		StackArray:   make([]int, 0, size),
		StackMaxSize: size,
		CurrentSize:  0,
	}
}

func (sk *Stack) Push(val int) bool {
	if sk.CurrentSize == sk.StackMaxSize {
		return false
	}
	sk.StackArray = append(sk.StackArray, val)
	sk.CurrentSize += 1
	return true
}

func (sk *Stack) Pop() (bool, int) {
	if sk.CurrentSize == 0 {
		return false, 0
	}
	output := sk.StackArray[sk.CurrentSize-1]
	sk.StackArray = sk.StackArray[:sk.CurrentSize-1]
	sk.CurrentSize -= 1
	return true, output
}

func (sk *Stack) Peek() (bool, int) {
	if sk.CurrentSize == 0 {
		return false, 0
	}
	output := sk.StackArray[sk.CurrentSize-1]
	return true, output
}

func (sk *Stack) IsEmpty() bool {
	return sk.CurrentSize == 0
}
