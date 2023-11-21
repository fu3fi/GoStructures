package main

import "fmt"

type Queue struct {
	data []interface{}
} 

func (q *Queue) pop() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value
}

func (q *Queue) push(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) len() int {
	return len(q.data)
}

func main() {
	queue := Queue{}
	queue.push(1)
	queue.push(2)
	fmt.Println(queue.len())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
}