package queue

import (
	"fmt"
	"testing"
)

var testQueue *Queue

func TestQueue(t *testing.T) {
	//test normal
	testQueue = NewQueue()
	testQueue.Push(1)
	testQueue.Push(2)
	fmt.Printf("testQueue Push (1,2) len:%d \n", testQueue.Len())

	testQueue.Push(3)
	testQueue.Push(4)
	testQueue.Push(5)
	testQueue.Push(6)
	testQueue.Push(7)
	testQueue.Push(8)
	testQueue.Push(9)
	fmt.Printf("testQueue Push (3,4,5,6,7,8,9) len:%d \n", testQueue.Len())

	v1 := testQueue.Pop()
	fmt.Printf("testQueue len:%d Pop v1:%v \n", testQueue.Len(), v1)
	v2 := testQueue.Pop()
	fmt.Printf("testQueue len:%d Pop v2:%v \n", testQueue.Len(), v2)

}
