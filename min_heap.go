package meetings

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

type MinHeap[T Ordered] struct {
	length int
	data   []T
}

func NewMinHeap[T Ordered]() MinHeap[T] {
	return MinHeap[T]{data: make([]T, 0)}
}

func (mh *MinHeap[T]) Insert(value T) {
	mh.data[mh.length] = value
	mh.heapifyUp(mh.length)
	mh.length++
}

func (mh *MinHeap[T]) Delete() (T, error) {
	if mh.length == 0 {
		var zero T
		return zero, fmt.Errorf("heap is empty")
	}

	min := mh.data[0]
	mh.length--

	if mh.length == 0 {
		mh.data = make([]T, 0)
		return min, nil
	}

	mh.data[0] = mh.data[mh.length]
	mh.heapifyDown(0)

	return min, nil
}

func (mh *MinHeap[T]) heapifyDown(idx int) {
	lIdx := mh.leftChild(idx)
	rIdx := mh.rightChild(idx)

	if idx >= mh.length || lIdx >= mh.length {
		return
	}

	lV := mh.data[lIdx]
	rV := mh.data[rIdx]
	v := mh.data[idx]

	if lV > rV && v > rV {
		swap(mh.data, idx, rIdx)
		mh.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		swap(mh.data, idx, lIdx)
		mh.heapifyDown(lIdx)
	}
}

func (mh *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := mh.parent(idx)
	parentV := mh.data[p]
	v := mh.data[idx]

	if parentV > v {
		swap(mh.data, p, idx)
		mh.heapifyUp(p)
	}
}

func swap[T any](arr []T, i, j int) {
	if i >= 0 && j >= 0 && i < len(arr) && j < len(arr) {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (mh *MinHeap[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (mh *MinHeap[T]) leftChild(idx int) int {
	return idx*2 + 1
}

func (mh *MinHeap[T]) rightChild(idx int) int {
	return idx*2 + 2
}
