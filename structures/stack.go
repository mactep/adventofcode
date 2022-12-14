package structures

type Stack []interface{}

// NewStack creates a new stack
func NewStack() Stack {
	s := make(Stack, 0)
	return s
}

// Push adds an element to the top of the stack
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

// PushBottom adds an element to the bottom of the stack
func (s *Stack) PushBottom(v interface{}) {
	*s = append([]interface{}{v}, *s...)
}

// Pop removes an element from the top of the stack
func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Peek returns the top element of the stack
func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}
