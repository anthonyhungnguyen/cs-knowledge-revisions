package linkedlist

import (
	"testing"
)

func TestLinkedListValueAt(t *testing.T) {
	// create a new linked list with values 1, 2, 3
	l := LinkedList{}
	l.pushed_front(1)
	l.pushed_front(2)
	l.pushed_front(3)

	// test getting the value at index 0, which should be 3
	val, err := l.value_at(0)
	if val != 3 || err != nil {
		t.Errorf("value_at(0) = (%d, %v), expected (3, nil)", val, err)
	}

	// test getting the value at index 1, which should be 2
	val, err = l.value_at(1)
	if val != 2 || err != nil {
		t.Errorf("value_at(1) = (%d, %v), expected (2, nil)", val, err)
	}

	// test getting the value at index 2, which should be 1
	val, err = l.value_at(2)
	if val != 1 || err != nil {
		t.Errorf("value_at(2) = (%d, %v), expected (1, nil)", val, err)
	}

	// test getting the value at index 3, which should return an error
	_, err = l.value_at(3)
	if err == nil {
		t.Errorf("value_at(3) did not return an error")
	}
}
