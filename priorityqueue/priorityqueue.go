/* Custom implementation of Priority Queue.
   Lower number indicates higher priority
*/

package priorityqueue

type PriorityQueue[T any] struct {
	Values []PQValue[T]
}

type PQValue[T any] struct {
	Value    T
	Priority int
}

func (pq *PriorityQueue[T]) Insert(value T, priority int) {
	newValue := PQValue[T]{Value: value, Priority: priority}
	pq.Values = append(pq.Values, newValue)
	if len(pq.Values) == 1 {
		return
	}
	pq.bubbleUp()
}

func (pq *PriorityQueue[T]) bubbleUp() {
	var currIdx = len(pq.Values) - 1
	for {
		parentIdx := (currIdx - 1) / 2
		parentElem := pq.Values[parentIdx]
		currElem := pq.Values[currIdx]
		if currElem.Priority >= parentElem.Priority {
			break
		}
		pq.Values[currIdx], pq.Values[parentIdx] = pq.Values[parentIdx], pq.Values[currIdx]
		currIdx = parentIdx
	}
}

func (pq *PriorityQueue[T]) ExtractMin() (T, int, bool) {
	var value T
	if len(pq.Values) == 0 {
		return value, -1, false
	}
	var min = pq.Values[0]
	pq.Values[0], pq.Values[len(pq.Values)-1] = pq.Values[len(pq.Values)-1], pq.Values[0]
	pq.Values = pq.Values[:len(pq.Values)-1]
	if len(pq.Values) > 1 {
		pq.sinkDown()
	}
	return min.Value, min.Priority, true
}

func (pq *PriorityQueue[T]) sinkDown() {
	if len(pq.Values) <= 1 {
		return
	}
	var currIdx int
	for {
		var leftChildIdx = 2*currIdx + 1
		var rightChildIdx = 2*currIdx + 2
		var swapIdx = -1
		var leftChild PQValue[T]
		var rightChild PQValue[T]
		if leftChildIdx < len(pq.Values) {
			leftChild = pq.Values[leftChildIdx]
			if leftChild.Priority < pq.Values[currIdx].Priority {
				swapIdx = leftChildIdx
			}
		}

		if rightChildIdx < len(pq.Values) {
			rightChild = pq.Values[rightChildIdx]
			if (swapIdx == -1 && rightChild.Priority < pq.Values[currIdx].Priority) || (swapIdx != -1 && rightChild.Priority < leftChild.Priority) {
				swapIdx = rightChildIdx
			}
		}
		if swapIdx == -1 {
			break
		}

		pq.Values[currIdx], pq.Values[swapIdx] = pq.Values[swapIdx], pq.Values[currIdx]
		currIdx = swapIdx
	}
}

func (pq *PriorityQueue[T]) Peek() (T, int, bool) {
	var value T
	if len(pq.Values) == 0 {
		return value, -1, false
	}
	return pq.Values[0].Value, pq.Values[0].Priority, true
}
