package deque

type Deque struct {
	data []int
}

func New() *Deque {
	return &Deque{}
}

func (q *Deque) PopBack() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	lastItem := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return lastItem, true
}

func (q *Deque) PushBack(newItem int) {
	q.data = append(q.data, newItem)
}

func (q *Deque) PushFront(newItem int) {
	q.data = append(q.data, 0)
	copy(q.data[1:], q.data[0:])
	q.data[0] = newItem
}

func (q *Deque) PopFront() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	lastItem := q.data[0]
	q.data = q.data[1:]
	return lastItem, true
}

func (q *Deque) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Deque) Size() int {
	return len(q.data)
}

func (q *Deque) Clear() {
	q.data = []int{}
}
