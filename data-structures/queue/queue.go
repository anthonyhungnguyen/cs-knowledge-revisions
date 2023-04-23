// Implement using linked-list, with tail pointer:
// enqueue(value) - adds value at position at tail
// dequeue() - returns value and removes least recently added element (front)
// empty()
//  Implement using fixed-sized array:
// enqueue(value) - adds item at end of available storage
// dequeue() - returns value and removes least recently added element
// empty()
// full()
//  Cost:
// a bad implementation using linked list where you enqueue at head and dequeue at tail would be O(n) because you'd need the next to last element, causing a full traversal each dequeue
// enqueue: O(1) (amortized, linked list and array [probing])
// dequeue: O(1) (linked list and array)
// empty: O(1) (linked list and array)

package queue

import "fmt"

type QueueNode struct {
	Value int
	Next  *QueueNode
}

// 1(Head)->2->3(Tail)
// Enqueue(0): 0(Head) -> 1 -> 2 -> 3 (Tail)
// Dequeue: 0 (Head) -> 1 -> 2 (Tail)

type QueueAsLinkedList struct {
	Head *QueueNode
	Tail *QueueNode
}

type QueueAsArray struct {
	Data []int
	Cur  int
}

type QueueGeneric[T any] struct {
	Data []T
	Cur  int
}

func (m *QueueAsLinkedList) Print() {
	if m == nil {
		return
	}

	tmp := m.Head
	for tmp != nil {
		fmt.Print(tmp.Value, " ")
		tmp = tmp.Next
	}
	fmt.Println()
}

func (m *QueueAsLinkedList) ValueAt(pos int) (int, error) {
	if m.Head == nil {
		return 0, fmt.Errorf("index out of range: %d", pos)
	}
	tmp := m.Head
	counter := 0
	for tmp != nil {
		if counter == pos {
			return tmp.Value, nil
		}
		counter++
		tmp = tmp.Next
	}
	return 0, fmt.Errorf("index out of range: %d", pos)
}

func (m *QueueAsLinkedList) Enqueue(value int) {

	if m.Head == nil {
		newNode := QueueNode{
			Value: value,
			Next:  nil,
		}
		m.Head = &newNode
		m.Tail = &newNode
	}

	// Alter tail by new value
	// 1 -> 2 (Tail)
	// 1 -> 2 -> 3 (Tail)

	newTail := QueueNode{value, nil}
	m.Tail.Next = &newTail
	m.Tail = &newTail
}

// dequeue() - returns value and removes least recently added element (front)
func (m *QueueAsLinkedList) Dequeue() error {
	if m.Head == nil {
		return fmt.Errorf("empty queue")
	}

	// Move head to head.next
	m.Head = m.Head.Next
	return nil
}

func (m *QueueAsLinkedList) IsEmpty() bool {
	return m.Head == nil
}

func (m *QueueAsArray) IsEmpty() bool {
	return m.Cur == 0
}
func (m *QueueAsArray) IsFull() bool {
	return m.Cur == 5
}

func (m *QueueAsArray) Enqueue(value int) {
	m.Data[m.Cur] = value
	m.Cur++
}

func (m *QueueAsArray) Dequeue() (int, error) {
	if m.Cur == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	newQueue := make([]int, 5)
	for i, v := range m.Data {
		if i >= 1 {
			newQueue[i-1] = v
		}
	}
	itemToDequeue := m.Data[0]
	m.Data = newQueue
	m.Cur--
	return itemToDequeue, nil
}

func (m *QueueGeneric[T]) Enqueue(item T) {
	m.Data = append(m.Data, item)
}

func (m *QueueGeneric[T]) Dequeue() (*T, error) {
	if len(m.Data) == 0 {
		return nil, fmt.Errorf("cannot dequeue empty queue")
	}
	itemToDequeue := m.Data[0]
	m.Data = m.Data[1:]
	return &itemToDequeue, nil
}
