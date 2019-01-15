package gofbwriter

type stack struct {
	s []string
}

func (s *stack) push(v string) {
	s.s = append(s.s, v)
}

func (s *stack) pop() string {
	l := len(s.s)
	if l == 0 {
		return ""
	}
	v := s.s[l-1]
	s.s = s.s[:l-1]
	return v
}

func (s *stack) peek() string {
	l := len(s.s)
	if l == 0 {
		return ""
	}
	return s.s[l-1]
}

func newStack() *stack {
	return &stack{[]string{}}
}

func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) length() int {
	return len(s.s)
}
