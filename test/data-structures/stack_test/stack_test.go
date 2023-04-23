package stacktest

import (
	"stack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
	data stack.StackGeneric[int]
}

func (suite *StackSuite) SetupSuite() {
	suite.data = *stack.New[int]()
}

func (s *StackSuite) TestPushPopStack() {
	stack := s.data
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(s.T(), stack.Data, []int{1, 2, 3})
	item, err := stack.Pop()
	assert.Empty(s.T(), err)
	assert.Equal(s.T(), *item, 3)
	assert.Equal(s.T(), stack.Data, []int{1, 2})
}

func TestStack(t *testing.T) {
	suite.Run(t, new(StackSuite))
}
