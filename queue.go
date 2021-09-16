//Package queue provides a fast, simpler queue struct.
//Not thread-safe.
package queue

const defaultQueueLen = 8 //Default queue capacity

type T interface{}

type Queue struct {
	datas     []T //Data list.
	count     int //Length of the queue.
	popIndex  int //Head index.
	pushIndex int //Tail index.
}

//NewQueue return a Queue.
func NewQueue() *Queue {
	return &Queue{
		datas:     make([]T, defaultQueueLen),
		count:     0,
		popIndex:  0,
		pushIndex: 0,
	}
}

//Push the data to the end of the queue
func (m *Queue) Push(data T) {
	if m.count == len(m.datas) {
		m.expand()
	}
	m.datas[m.pushIndex] = data
	m.count++
	m.pushIndex = (m.pushIndex + 1) % len(m.datas)
}

//Pop the header data of the pop queue.
func (m *Queue) Pop() (data T) {
	if m.count == 0 {
		return nil
	}
	data = m.datas[m.popIndex]
	m.datas[m.popIndex] = nil
	m.count--
	m.popIndex = (m.popIndex + 1) % len(m.datas)
	return data
}

//Return queue length.
func (m *Queue) Len() int {
	return m.count
}

//Auto expand capacity.
//When the capacity is less than 1024, double the expansion,
//otherwise expand 1 / 4 of the current capacity.
func (m *Queue) expand() {
	curCap := len(m.datas)
	var newCap int
	if curCap == 0 {
		newCap = defaultQueueLen
	} else if curCap < 1024 {
		newCap = curCap * 2
	} else {
		newCap = curCap + (curCap / 4)
	}
	datas := make([]T, newCap)

	if m.popIndex == 0 {
		copy(datas, m.datas)
		m.pushIndex = curCap
	} else {
		newPopIndex := newCap - (curCap - m.popIndex)
		copy(datas, m.datas[:m.popIndex])
		copy(datas[newPopIndex:], m.datas[m.popIndex:])
		m.popIndex = newPopIndex
	}
	for i := range m.datas {
		m.datas[i] = nil //Help GC.
	}
	m.datas = datas
}
