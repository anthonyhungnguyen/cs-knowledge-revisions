package main

import "fmt"

type Array struct {
	size     int
	capacity int
	Data     []int
}

func (m *Array) get_size() int {
	return m.size
}

func (m *Array) get_capacity() int {
	return m.capacity
}

func (m *Array) is_empty() bool {
	return m.size == 0
}

func (m *Array) at(index int) int {
	if index <= 0 || index > m.size {
		panic("Index out of range")
	}
	return m.Data[index]
}

func (m *Array) insert(index int, item int) {
	if index < 0 || index >= m.capacity {
		panic("Index out of range")
	}
	m.resize()
	if index == m.size-1 {
		m.push(item)
	} else {
		copy(m.Data[index+1:], m.Data[index:])
		m.Data[index] = item
	}
	m.size++
}

func (m *Array) prepend(item int) {
	m.insert(0, item)
	m.size++
}

func (m *Array) pop() {
	if m.size < 0 {
		panic("Index out of range")
	} else {
		m.Data[m.size-1] = 0
		m.size--
	}
}

func (m *Array) delete(index int) {
	if index < 0 || index >= m.size {
		panic("Index out of range")
	}
	copy(m.Data[index:], m.Data[index+1:])
	m.pop()
	m.resize()
}

func (m *Array) find(item int) int {
	for i, s := range m.Data {
		if s == item {
			return i
		}
	}
	return -1
}

func (m *Array) remove(item int) {
	for {
		i := m.find(item)
		if i != -1 {
			m.delete(i)
		} else {
			break
		}
	}
	m.resize()
}

func (m *Array) resize() {
	needs_resize := false
	if m.capacity == 0 {
		m.capacity += 1
		needs_resize = true
	} else if m.size == m.capacity {
		m.capacity *= 2
		needs_resize = true
	} else if m.size < m.capacity/4 {
		m.capacity /= 4
		needs_resize = true
	}
	if needs_resize {
		new_data := make([]int, m.capacity)
		copy(new_data[:], m.Data[:])
		m.Data = new_data
	}
}

func (m *Array) push(val int) {
	m.resize()
	m.Data[m.size] = val
	m.size++
}

func main() {
	array := Array{
		capacity: 0,
		Data:     []int{},
	}

	fmt.Println("Array capacity is: ", array.get_capacity())
	fmt.Println("Array size is: ", array.get_size())
	fmt.Println("Array is empty: ", array.is_empty())
	array.push(1)

	fmt.Println("Array capacity is: ", array.get_capacity())
	fmt.Println("Array size is: ", array.get_size())
	fmt.Println("Array is empty: ", array.is_empty())
	fmt.Println("Array data is: ", array.Data)

	array.push(2)
	array.push(3)
	array.push(4)
	array.push(5)

	fmt.Println("Array capacity is: ", array.get_capacity())
	fmt.Println("Array size is: ", array.get_size())
	fmt.Println("Array is empty: ", array.is_empty())
	fmt.Println("Array data is: ", array.Data)
	fmt.Println("Array data at index 2 is: ", array.at(2))
	// fmt.Println("Array data at index 10 is: ", array.at(10))

	// array.insert(10, 10)
	// fmt.Println("Array data is: ", array.Data)

	array.prepend(10)
	fmt.Println("Array data is: ", array.Data)

	array.pop()
	fmt.Println("Array data is: ", array.Data)

	array.delete(2)
	array.delete(0)
	fmt.Println("Array data is: ", array.Data)

	fmt.Println("Array data of value 3 is at index: ", array.find(3))
	array.push(3)
	array.remove(3)
	array.remove(1)
	fmt.Println("Array data is: ", array.Data)

	array.remove(4)
	fmt.Println(array.size, array.capacity)
}
