package memory

import (
	"container/list"
	"fmt"
)

type Queue struct {
	data *list.List
}

func New() *Queue {
	return &Queue{data: list.New()}
}

func (q *Queue) Enqueue(value int) {
	q.data.PushBack(value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}

	front := q.data.Front()
	q.data.Remove(front)

	return front.Value.(int), nil
}

func (q *Queue) Front() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}

	return q.data.Front().Value.(int), nil
}

func (q *Queue) isEmpty() bool {
	return q.data.Len() == 0
}
