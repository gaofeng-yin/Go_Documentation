package data_structure

import "log"

type Deque struct {
	items []int
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) AddFirst(i int) {
	//Prepend
	d.items = append([]int{i}, d.items...)
}

func (d *Deque) AddLast(i int) {
	d.items = append(d.items, i)
}

func (d *Deque) RemoveFirst() int {
	if d.IsEmpty() {
		log.Panic("Empty Deque")
	}

	toRemove := d.items[0]
	d.items = d.items[1:]

	return toRemove
}

func (d *Deque) RemoveLast() int {
	if d.IsEmpty() {
		log.Panic("Empty Deque")
	}

	last := len(d.items) - 1

	toRemove := d.items[last]

	d.items = d.items[:last]

	return toRemove
}
