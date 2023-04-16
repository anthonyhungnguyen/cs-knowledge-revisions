package hashtable

import "fmt"

const TABLE_SIZE = 10

type HashTable struct {
	Keys   []int
	Values []int
}

func NewHashTable() *HashTable {
	return &HashTable{
		Keys:   make([]int, TABLE_SIZE),
		Values: make([]int, TABLE_SIZE),
	}
}

func hash(k int, m int) int {
	return k % m
}

func (m *HashTable) Print() {
	for i, v := range m.Keys {
		fmt.Printf("Key: %d, Value: %d\n", v, m.Values[i])
	}
}

func (m *HashTable) Add(key int, value int) {
	pos := hash(key, TABLE_SIZE)
	for m.Keys[pos] != 0 && m.Keys[pos] != key {
		pos = hash(pos+1, TABLE_SIZE)
	}
	m.Keys[pos] = key
	m.Values[pos] = value
}

func (m *HashTable) Exists(key int) bool {
	pos := hash(key, TABLE_SIZE)
	for m.Keys[pos] != 0 {
		if m.Keys[pos] == key {
			return true
		}
		pos = hash(pos+1, TABLE_SIZE)
	}
	return false
}

func (m *HashTable) Get(key int) (int, error) {
	pos := hash(key, TABLE_SIZE)
	for m.Keys[pos] != 0 {
		if m.Keys[pos] == key {
			return m.Values[pos], nil
		}
		pos = hash(pos+1, TABLE_SIZE)
	}
	return 0, fmt.Errorf("key not found")
}

func (m *HashTable) Remove(key int) error {
	pos := hash(key, TABLE_SIZE)
	for m.Keys[pos] != 0 {
		if m.Keys[pos] == key {
			m.Keys[pos] = 0
			m.Values[pos] = 0
			return nil
		}
		pos = hash(pos+1, TABLE_SIZE)
	}
	return fmt.Errorf("key not found")
}
