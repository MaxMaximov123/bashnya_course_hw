package stack

type Stack struct {
	data []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	lastItem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return lastItem, true
}

func (s *Stack) Push(newItem int) {
	s.data = append(s.data, newItem)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = []int{}
}
