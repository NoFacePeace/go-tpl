package queue

import (
	"log"
)

type Array struct {
	front int
	rear  int
	size  int
	queue []interface{}
}

func NewArray(size int) *Array {
	if size <= 0 {
		log.Fatal("the queue size cannot be lower than zero")
	}
	return &Array{
		front: 0,
		rear:  0,
		size:  size,
		queue: make([]interface{}, size),
	}
}

func (a *Array) Put(v interface{}) bool {
	if (a.rear+1)%a.size == a.front {
		return false
	}
	a.queue[a.rear] = v
	a.rear = (a.rear + 1) % a.size
	return true
}

func (a *Array) Poll() (interface{}, bool) {
	if a.rear == a.front {
		return nil, false
	}
	v := a.queue[a.front]
	a.front = (a.front + 1) % a.size
	return v, true
}
