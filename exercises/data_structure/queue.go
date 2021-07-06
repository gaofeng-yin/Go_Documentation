package data_structure

import "log"

type Queue struct {
	items []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

//Add in the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Remove from beginning
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		log.Panic("Empty Queue")
	}
	toRemove := q.items[0]

	q.items = q.items[1:]

	return toRemove
}
