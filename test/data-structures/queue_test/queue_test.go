package queue_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"queue"
	"testing"
)

type QueueAsLinkedListSuite struct {
	suite.Suite
	data queue.QueueAsLinkedList
}

type QueueAsArraySuite struct {
	suite.Suite
	data queue.QueueAsArray
}

func (suite *QueueAsLinkedListSuite) SetupTest() {
	// Prepare data before each test
	initNode := queue.QueueNode{
		Value: 0,
		Next:  nil,
	}
	suite.data = queue.QueueAsLinkedList{
		Head: &initNode,
		Tail: &initNode,
	}
}

func (suite *QueueAsArraySuite) SetupTest() {
	// Prepare data before each test
	suite.data = queue.QueueAsArray{
		Data: make([]int, 5),
		Cur:  0,
	}
}

func (suite *QueueAsLinkedListSuite) TestQueueLinkedListEnqueue() {
	suite.data.Enqueue(1)
	res, _ := suite.data.ValueAt(1)
	assert.Equal(suite.T(), 1, res, "Expected 1 but got %d", res)
	suite.data.Enqueue(2)

	res, _ = suite.data.ValueAt(2)
	assert.Equal(suite.T(), 2, res, "Expected 2 but got %d", res)
}

func (suite *QueueAsLinkedListSuite) TestQueueLinkedListDequeue() {
	queue := suite.data
	queue.Dequeue()
	_, err := queue.ValueAt(0)
	assert.NotNil(suite.T(), err, "Error should not be nil")
}

func (suite *QueueAsLinkedListSuite) TestQueueLinkedListEmpty() {
	queue := suite.data
	queue.Dequeue()
	res := queue.IsEmpty()
	assert.Equal(suite.T(), true, res, "Queue should be empty")
}

func (suite *QueueAsArraySuite) TestQueueArrayEmpty() {
	assert.Equal(suite.T(), true, suite.data.IsEmpty(), "Queue should be empty")
}

func (suite *QueueAsArraySuite) TestQueueArrayEnqueue() {
	queue := suite.data
	queue.Enqueue(1)
	fmt.Println(queue)
	assert.Equal(suite.T(), 1, queue.Data[0], "Expected 1 but got %d", queue.Data[0])
}

func (suite *QueueAsArraySuite) TestQueueArrayDequeue() {
	queue := suite.data
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	assert.Equal(suite.T(), 2, queue.Data[0], "Expected 2 but got %d", queue.Data[0])
}

func (suite *QueueAsArraySuite) TestQueueArrayIsFull() {
	queue := suite.data
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	assert.Equal(suite.T(), true, queue.IsFull(), "Queue should be full")
}

func TestQueueAsLinkedListSuite(t *testing.T) {
	suite.Run(t, new(QueueAsLinkedListSuite))
}

func TestQueueAsArraySuite(t *testing.T) {
	suite.Run(t, new(QueueAsArraySuite))
}
