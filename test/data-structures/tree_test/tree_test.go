package treetest

import (
	"testing"
	"tree"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BinaryTreeSuite struct {
	suite.Suite
	data tree.BinarySearchTree
}

func (suite *BinaryTreeSuite) SetupSuite() {
	suite.data = *tree.NewBinarySearchTree()
}

func (suite *BinaryTreeSuite) TestTraversePreOrder() {
	bt := suite.data
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 6)
	bt.Insert(bt.Root, 2)
	assert.Equal(suite.T(), tree.TraversePreOrder(bt.Root), []int{5, 1, 2, 6})
}
func (suite *BinaryTreeSuite) TestTraverseInOrder() {
	bt := suite.data
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 6)
	bt.Insert(bt.Root, 2)
	assert.Equal(suite.T(), tree.TraverseInOrder(bt.Root), []int{1, 2, 5, 6})
}
func (suite *BinaryTreeSuite) TestTraversePostOrder() {
	bt := suite.data
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 6)
	bt.Insert(bt.Root, 2)
	assert.Equal(suite.T(), tree.TraversePostOrder(bt.Root), []int{2, 1, 6, 5})
}

func (suite *BinaryTreeSuite) TestHeightTree() {
	bt := suite.data
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 0)
	bt.Insert(bt.Root, -1)
	bt.Insert(bt.Root, -2)
	curHeight := bt.Height(bt.Root)
	assert.Equal(suite.T(), curHeight, 3, "Expected height is 3 but got %d", curHeight)
}

// -----1-----
// -------9--
// ------3----
// --------5--
func (suite *BinaryTreeSuite) TestBFS() {
	bt := suite.data
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 9)
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 4)
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 6)
	bt.Insert(bt.Root, 7)
	assert.Equal(suite.T(), bt.BFS(bt.Root), []int{2, 1, 9, 4, 5, 6, 7})
}

func (suite *BinaryTreeSuite) TestDFS() {
	bt := suite.data
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 6)
	bt.Insert(bt.Root, 0)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 8)
	bt.Insert(bt.Root, 4)
	bt.Insert(bt.Root, 3)
	assert.Equal(suite.T(), bt.DFS(bt.Root), []int{5, 6, 8, 1, 2, 4, 3, 0})
}

func (suite *BinaryTreeSuite) TestIsValid() {
	bt := suite.data
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 1)
	bt.Insert(bt.Root, 6)
	bt.Insert(bt.Root, 0)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 8)
	bt.Insert(bt.Root, 4)
	bt.Insert(bt.Root, 3)
	assert.Equal(suite.T(), tree.IsValid(bt.Root), true)
}

func (suite *BinaryTreeSuite) TestIsInValid() {
	root := &tree.TreeNode{
		Value: 2,
		Left: &tree.TreeNode{
			Value: 3,
			Left:  nil,
			Right: nil,
		},
		Right: &tree.TreeNode{
			Value: 4,
			Left:  nil,
			Right: nil,
		},
	}
	assert.Equal(suite.T(), tree.IsValid(root), false)
}

func (suite *BinaryTreeSuite) TestTreeDelete() {
	bt := suite.data
	bt.Insert(bt.Root, 10)
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 15)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 7)
	bt.Insert(bt.Root, 20)
	bt.DeleteNode(bt.Root, 2)
	assert.Equal(suite.T(), tree.TraverseInOrder(bt.Root), []int{5, 7, 10, 15, 20})
	bt.DeleteNode(bt.Root, 5)
	assert.Equal(suite.T(), tree.TraverseInOrder(bt.Root), []int{7, 10, 15, 20})
	bt.Insert(bt.Root, 9)
	bt.Insert(bt.Root, 8)
	bt.DeleteNode(bt.Root, 10)
	assert.Equal(suite.T(), tree.TraverseInOrder(bt.Root), []int{7, 8, 9, 15, 20})
}

func (suite *BinaryTreeSuite) TestNodeCount() {
	bt := suite.data
	bt.Insert(bt.Root, 10)
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 15)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 7)
	bt.Insert(bt.Root, 20)
	assert.Equal(suite.T(), 6, tree.GetNodeCount(bt.Root))
}

func (suite *BinaryTreeSuite) TestIsInTree() {
	bt := suite.data
	bt.Insert(bt.Root, 10)
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 15)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 7)
	bt.Insert(bt.Root, 20)
	assert.True(suite.T(), tree.IsInTree(bt.Root, 2))
	assert.False(suite.T(), tree.IsInTree(bt.Root, 100))
}

func (suite *BinaryTreeSuite) TestMaxMin() {
	bt := suite.data
	bt.Insert(bt.Root, 10)
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 15)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 7)
	bt.Insert(bt.Root, 20)
	min, _ := tree.GetMin(bt.Root)
	max, _ := tree.GetMax(bt.Root)
	assert.Equal(suite.T(), 2, min)
	assert.Equal(suite.T(), 20, max)
}
func (suite *BinaryTreeSuite) TestDeleteTree() {
	bt := suite.data
	bt.Insert(bt.Root, 10)
	bt.Insert(bt.Root, 5)
	bt.Insert(bt.Root, 15)
	bt.Insert(bt.Root, 2)
	bt.Insert(bt.Root, 7)
	bt.Insert(bt.Root, 20)
	tree.DeleteTree(&bt.Root)
	assert.Equal(suite.T(), 0, tree.GetNodeCount(bt.Root))
}

func TestBinaryTree(t *testing.T) {
	suite.Run(t, new(BinaryTreeSuite))
}
