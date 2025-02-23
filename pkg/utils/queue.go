package pkg

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

// Enqueue adds an element to the back of the queue
func (q *Queue) Enqueue(value any) {
	q.list.PushBack(value)
}

// Dequeue removes and returns the front element
func (q *Queue) Dequeue() any {
	front := q.list.Front()
	if front != nil {
		q.list.Remove(front)
		return front.Value
	}
	return nil
}

// Peek returns the front element without removing it
func (q *Queue) Peek() any {
	front := q.list.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}
