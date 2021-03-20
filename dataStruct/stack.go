package dataStruct

type Stact struct {
	ll *LinkedList
}

func NewStack() *Stact {
	return &Stact{ll: &LinkedList{}}
}

func (s *Stact) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stact) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}

func (s *Stact) Empty() bool {
	return s.ll.Empty()
}
