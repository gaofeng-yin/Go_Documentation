package data_structure

import "log"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	//parameter: index of last inserted
	h.maxHeapifyUp(len(h.array) - 1)

}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		log.Panic("Empty array")
	}
	extracted := h.array[0]
	last := len(h.array) - 1

	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

func (h MaxHeap) IsEmpty() bool {
	return len(h.array) == 0
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := getLeftChildIndex(index), getRightChildIndex(index)
	childToCompare := 0

	for left <= lastIndex {
		if left == lastIndex {
			childToCompare = left
		} else if h.array[left] > h.array[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			left, right = getLeftChildIndex(index), getRightChildIndex(index)
		} else {
			return
		}
	}
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

func getLeftChildIndex(i int) int {
	return 2*i + 1
}

func getRightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(ind1, ind2 int) {
	h.array[ind1], h.array[ind2] = h.array[ind2], h.array[ind1]
}
