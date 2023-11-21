package main

import "fmt"

type Set struct {
	data map[interface{}]bool
} 

// func (s *Set) pop() interface{} {
// 	if len(s.data) == 0 {
// 		return nil
// 	}
// 	value := s.data[len(s.data)-1]
// 	s.data = s.data[:len(s.data)-1]
// 	return value
// }

func (s *Set) add(value interface{}) {
	s.data[value] = true
}

func (s *Set) check(value interface{}) bool {
	_, ok := s.data[value]
	return ok
}

func (s *Set) len() int {
	return len(s.data)
}

func (s *Set) isEmpty() bool {
	return len(s.data)
}


func main() {
	set := Set{}
	set.add(1)
	set.add(2)
	fmt.Println(set.len())
	fmt.Println(set.pop())
	fmt.Println(set.pop())
	fmt.Println(set.pop())
}