package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (m *LinkedList) back() *Node {
	if !m.is_empty() {
		tmp := m.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		return tmp
	}
	return nil
}

func (m *LinkedList) print() {
	tmp := m.head
	for tmp != nil {
		fmt.Print(tmp.value, " ")
		tmp = tmp.next
	}
	fmt.Println()
}

func (m *LinkedList) get_size() int {
	return m.size
}

func (m *LinkedList) is_empty() bool {
	return m.size == 0
}

func (m *LinkedList) Value_at(index int) (int, error) {
	idx_cnt := 0
	tmp := m.head
	for tmp != nil {
		if index == idx_cnt {
			return tmp.value, nil
		} else {
			idx_cnt++
			tmp = tmp.next
		}
	}
	return 0, fmt.Errorf("Index out of range: %d", index)
}

func (m *LinkedList) Pushed_front(value int) {
	new_node := Node{
		value: value,
		next:  m.head,
	}
	m.head = &new_node
	m.size++
}

func (m *LinkedList) pop_front() {
	if !m.is_empty() {
		new_head := m.head.next
		m.head = new_head
		m.size--
	}
}

func (m *LinkedList) push_back(value int) {
	new_node := Node{
		value: value,
		next:  nil,
	}
	tail_node := m.back()
	if tail_node != nil {
		tail_node.next = &new_node
	} else {
		m.head = &new_node
	}
	m.size++
}

func (m *LinkedList) pop_back() {
	tmp := m.head
	for tmp != nil {
		if tmp.next != nil && tmp.next.next == nil {
			tmp.next = nil
			m.size--
			break
		} else {
			tmp = tmp.next
		}
	}
}

func (m *LinkedList) front() *Node {
	if !m.is_empty() {
		return m.head
	}
	return nil
}

func (m *LinkedList) insert(index int, value int) error {
	if index < 0 || index > m.size {
		return fmt.Errorf("Index out of range: %d", index)
	} else if index == 0 {
		new_node := Node{
			value: value,
			next:  m.head,
		}
		m.head = &new_node
		m.size++
	} else {
		cur_idx := 0
		tmp := m.head
		for tmp != nil {
			if cur_idx < index-1 && tmp.next != nil {
				cur_idx++
				tmp = tmp.next
			} else if cur_idx == index-1 {
				new_node := Node{
					value: value,
					next:  tmp.next,
				}
				tmp.next = &new_node
				m.size++
				break
			}
		}
	}
	return nil
}

func (m *LinkedList) erase(index int) error {
	if index < 0 || index >= m.size {
		return fmt.Errorf("Index out of range: %d", index)
	} else if index == 0 {
		m.head = m.head.next
		m.size--
	} else {
		tmp := m.head
		cur_idx := 0
		for tmp != nil {
			if cur_idx < index-1 && tmp.next != nil {
				cur_idx++
				tmp = tmp.next
			} else if cur_idx == index-1 && tmp.next != nil {
				tmp.next = tmp.next.next
				m.size--
				break
			}
		}
	}
	return nil
}

func (m *LinkedList) value_n_from_end(n int) (int, error) {
	if !m.is_empty() {
		size := m.get_size()
		if n == 0 {
			return m.back().value, nil
		} else if n == size-1 {
			return m.front().value, nil
		} else if n > 0 && n < size-1 {
			cur_idx := 0
			tmp := m.head
			for tmp != nil {
				if size-n-1 == cur_idx {
					return tmp.value, nil
				} else {
					cur_idx++
					tmp = tmp.next
				}
			}
		}
	}
	return 0, fmt.Errorf("Invalid n value: %d", n)
}

// 0(t1)->1(t1)->2(t2)->3->4
// t0 = 0, t1 = 1, t2 = 2
// t1.next = t0
// t0 = t1, t1 = t2, t2 = t2.next
// 0<-1(t0) 2(t1)->3(t2)->4
// t0 = 1, t1 = 2, t2 = 3
// t1.next = t0
// t0 = t1, t1 = t2, t2 = t2.next
// 0<-1<-2(t0) 3(t1)->4(t2)
// 0<-1<-2<-3(t0) 4(t1) nil(t2)
// 0<-1<-2<-3<-4(t0)
// head.next = nil
// m.head = t0

func (m *LinkedList) reverse() {
	// Reverse inner list
	if m.head != nil && m.head.next != nil {
		head := m.head
		t0 := m.head
		t1 := m.head.next
		t2 := m.head.next.next
		for t2 != nil {
			t1.next = t0
			t0 = t1
			t1 = t2
			t2 = t2.next
		}
		t1.next = t0
		m.head = t1
		head.next = nil
	}
}

func (m *LinkedList) remove_value(val int) {
	if !m.is_empty() {
		t := m.front()
		if t.value == val {
			m.head = m.head.next
			m.size--
			return
		}
		for t.next != nil {
			if t.next.value == val {
				t.next = t.next.next
				m.size--
				break
			} else {
				t = t.next
			}
		}
	}
}

func main() {
	ll := LinkedList{
		head: &Node{
			value: 0,
			next:  nil,
		},
		size: 1,
	}
	ll.push_back(1)
	ll.push_back(2)

	fmt.Println("Size of linked list:", ll.get_size())
	fmt.Println("Is empty:", ll.is_empty())
	Value_at_0, _ := ll.Value_at(0)
	Value_at_1, _ := ll.Value_at(1)
	Value_at_2, _ := ll.Value_at(2)
	fmt.Println("Value at index 0: ", Value_at_0)
	fmt.Println("Value at index 1: ", Value_at_1)
	fmt.Println("Value at index 2: ", Value_at_2)
	ll.print()
	ll.Pushed_front(-1)
	ll.print()
	ll.Pushed_front(-2)
	ll.print()
	ll.pop_front()
	ll.print()
	ll.push_back(3)
	ll.print()
	ll.pop_back()
	ll.print()
	fmt.Println("Front of linked list: ", ll.front())
	fmt.Println("Back of linked list: ", ll.back())
	ll.insert(0, -3)
	ll.print()
	ll.insert(1, -4)
	ll.print()
	ll.insert(6, 6)
	ll.print()
	// ll.insert(8, 7) // Raise exception
	// ll.print()
	ll.erase(0)
	ll.print()
	ll.erase(2)
	ll.print()
	ll.erase(4)
	ll.print()
	// ll.erase(5) # Raise exception
	// ll.print()
	nValue_at_0, _ := ll.value_n_from_end(0)
	nValue_at_size, _ := ll.value_n_from_end(ll.get_size() - 1)
	nValue_at_1, _ := ll.value_n_from_end(1)
	_, error_ := ll.value_n_from_end(4)
	fmt.Println("Value at 0th from the end: ", nValue_at_0)
	fmt.Println("Value at size th from the end: ", nValue_at_size)
	fmt.Println("Value at 1th from the end: ", nValue_at_1)
	fmt.Println("Value at 4th from the end, caught error: ", error_)
	fmt.Println("Reverse list: ")
	ll.reverse()
	ll.print()
	ll.reverse()
	ll.pop_front()
	ll.pop_front()
	ll.print()
	ll.reverse()
	ll.print()
	ll.remove_value(2)
	ll.print()
	ll.remove_value(1)
	ll.print()
}
