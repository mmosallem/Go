package main

//StackString Stack of strings
type StackString []string

//IsEmpty check if stack is empty
func (s *StackString) IsEmpty() bool {
	return len(*s) == 0
}

//Push a new integer onto the stack
func (s *StackString) Push(x string) {
	*s = append(*s, x)
}

//Pop remove and return top element of stack, return false if stack is empty
func (s *StackString) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

//Peek return top element of stack, return false if stack is empty
func (s *StackString) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(*s) - 1
	x := (*s)[i]

	return x, true
}
