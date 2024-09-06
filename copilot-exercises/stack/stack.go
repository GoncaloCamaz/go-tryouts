// implementing a stack with the functions push, pop and peek
// push function adds an element to the top of the stack
// pop function removes the top element from the stack
// peek function returns the top element of the stack

package main

import "fmt"

type Stack struct {
	stack []int
}

func (s *Stack) push(value int) {
	s.stack = append(s.stack, value)
}

// since push is adding to the top of the stack, pop should remove the last element from the stack
func (s *Stack) pop() int {
	item := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return item
}

func (s *Stack) peek() int {
	return s.stack[len(s.stack)-1]
}

func main() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
