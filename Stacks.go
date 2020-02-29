package main

//Stack of strings
type Stack []string

//IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Push a new integer onto the stack
func (s *Stack) Push(x string) {
	*s = append(*s, x)
}

//Pop remove and return top element of stack, return false if stack is empty
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

//Peek return top element of stack, return false if stack is empty
func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(*s) - 1
	x := (*s)[i]

	return x, true
}

//isBalanced checks if the string is balanced strings meaning [],{},() brackets are balanced
func isBalanced(str string) bool {
	var s Stack
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		var curr = string(r[i])
		if curr == "[" || curr == "{" || curr == "(" {
			s.Push(curr)
		} else if curr == "]" || curr == "}" || curr == ")" {
			var prev, b = s.Pop()
			if !b {
				return false
			}
			if (prev == "[" && curr != "]") || (prev == "{" && curr != "}") || (prev == "(" && curr != ")") {
				return false
			}

		}

	}
	if s.IsEmpty() {
		return true
	}
	return false
}
