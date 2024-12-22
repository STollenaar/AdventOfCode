package internal

import "sort"

type Queue[T comparable] struct {
	Elements      []T
	SortFunction  func(i, j int) bool
	EqualFunction func(a, b T) bool
}

// Add element to the end of the queue
func (q *Queue[T]) Push(p ...T) {
	q.Elements = append(q.Elements, p...)
}

// Add element to the end of the queue if not there
func (q *Queue[T]) PushUnique(p ...T) {
	for _, t := range p {
		for _, e := range q.Elements {
			if !q.EqualFunction(e, t) {
				q.Elements = append(q.Elements, p...)
			}
		}
	}
}

// Add element to beginning of the queue
func (q *Queue[T]) Enqueue(p T) {
	q.Elements = append([]T{p}, q.Elements...)
}

func (q *Queue[T]) Delete(point T) {
	for i, iQ := range q.Elements {
		if iQ == point {
			q.Elements = append(q.Elements[:i], q.Elements[i+1:]...)
			break
		}
	}
}

func (q *Queue[T]) Find(p T) interface{} {
	for _, iQ := range q.Elements {
		if q.EqualFunction(iQ, p) {
			return iQ
		}
	}
	return nil
}

func (q *Queue[T]) FindIndex(p T) int {
	for i, iQ := range q.Elements {
		if q.EqualFunction(iQ, p) {
			return i
		}
	}
	return -1
}

// Get element from the start of the queue
func (q *Queue[T]) Shift() T {
	e := q.Elements[0]
	q.Elements = q.Elements[1:]
	return e
}

// Get element from the end of the queue
func (q *Queue[T]) Pop() T {
	e := q.Elements[len(q.Elements)-1]
	q.Elements = q.Elements[:len(q.Elements)-1]
	return e
}

// Sort the queue
func (q *Queue[T]) Sort() {
	sort.Slice(q.Elements, q.SortFunction)
}

func (q *Queue[T]) Empty() {
	q.Elements = []T{}
}
