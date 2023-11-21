package main

import "fmt"

type Stack struct {
	data []interface{}
} 

func (s *Stack) pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *Stack) push(value interface{}) {
	s.data = append(s.data, value)
}

func (s *Stack) len() int {
	return len(s.data)
}

func main() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	fmt.Println(stack.len())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}