package dataStructures

import "log"

func run() {
	queue := newTestArray(4)
	queue.insert(1)
	queue.insert(3)
	queue.insert(9)
	queue.insert(2)
	log.Print(queue.get())
	log.Print(queue.get())
	queue.insert(21)
	log.Print(queue.get())
	log.Print(queue.get())
	log.Print(queue.get())
}

type queueTest struct {
	queue       []int
	currentSize int
	maxSize     int
}

func newTestArray(size int) queueTest {
	toReturn := queueTest{
		queue:   make([]int, 0, size),
		maxSize: size,
	}
	return toReturn
}

func (qu *queueTest) insert(val int) {
	if qu.currentSize == qu.maxSize {
		return
	}
	qu.queue = append(qu.queue, val)
	qu.currentSize += 1
}

func (qu *queueTest) get() int {
	if qu.currentSize == 0 {
		return 0
	}
	qu.currentSize -= 1
	toReturn := qu.queue[0]
	qu.queue = qu.queue[1:]
	return toReturn
}
