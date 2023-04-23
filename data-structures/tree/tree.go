package tree

import (
	"fmt"
	"queue"
	"stack"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Root: nil,
	}
}

// left, self, right

func TraverseInOrder(root *TreeNode) []int {
	res := make([]int, 0)
	traverse_in_order(root, &res)
	return res
}

func traverse_in_order(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traverse_in_order(root.Left, res)
	}
	*res = append(*res, root.Value)
	if root.Right != nil {
		traverse_in_order(root.Right, res)
	}
}

// self, left, right
func TraversePreOrder(root *TreeNode) []int {
	res := make([]int, 0)
	traverse_pre_order(root, &res)
	return res
}

func traverse_pre_order(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Value)
	if root.Left != nil {
		traverse_pre_order(root.Left, res)
	}
	if root.Right != nil {
		traverse_pre_order(root.Right, res)
	}
}

// left, right, self
func TraversePostOrder(root *TreeNode) []int {
	res := make([]int, 0)
	traverse_post_order(root, &res)
	return res
}

func traverse_post_order(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traverse_post_order(root.Left, res)
	}
	if root.Right != nil {
		traverse_post_order(root.Right, res)
	}
	*res = append(*res, root.Value)
}

func (t *BinarySearchTree) Insert(root *TreeNode, val int) *TreeNode {
	if t.Root == nil {
		t.Root = &TreeNode{
			Value: val,
			Left:  nil,
			Right: nil,
		}
		return t.Root
	}

	if root == nil {
		root = &TreeNode{
			Value: val,
			Left:  nil,
			Right: nil,
		}
		return root
	}
	if root.Value <= val {
		root.Right = t.Insert(root.Right, val)
	} else {
		root.Left = t.Insert(root.Left, val)
	}
	return root
}

// length from root to leaf node
// -------- root -------
// --left--------right--

func (t *BinarySearchTree) Height(root *TreeNode) int {
	if root.Left != nil && root.Right != nil {
		return max(t.Height(root.Left), t.Height(root.Right)) + 1
	}
	if root.Left != nil {
		return t.Height(root.Left) + 1
	}
	if root.Right != nil {
		return t.Height(root.Right) + 1
	}
	return 0
}

// Breath-first search
// ------1-------
// ----2----3----
// --4--5--6--7--

// Queue: 1 2 4 5 3 6 7
func (t *BinarySearchTree) BFS(root *TreeNode) []int {
	res := make([]int, 0)
	queue := queue.QueueGeneric[TreeNode]{
		Data: make([]TreeNode, 0),
		Cur:  0,
	}
	queue.Enqueue(*root)
	t.bfs(queue, &res)
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	return res
}

func (t *BinarySearchTree) bfs(queue queue.QueueGeneric[TreeNode], res *[]int) {
	item, err := queue.Dequeue()

	if err != nil {
		return
	}
	*res = append(*res, item.Value)
	if item.Left != nil {
		queue.Enqueue(*item.Left)
	}
	if item.Right != nil {
		queue.Enqueue(*item.Right)
	}
	t.bfs(queue, res)
}

// Use Stack
func (t *BinarySearchTree) DFS(root *TreeNode) []int {
	res := make([]int, 0)
	stack := stack.New[TreeNode]()
	stack.Push(*root)
	t.dfs(stack, &res)
	return res
}

func (t *BinarySearchTree) dfs(stack *stack.StackGeneric[TreeNode], res *[]int) {
	item, err := stack.Pop()
	if err != nil {
		return
	}
	*res = append(*res, item.Value)
	if item.Left != nil {
		stack.Push(*item.Left)
	}
	if item.Right != nil {
		stack.Push(*item.Right)
	}
	t.dfs(stack, res)
}

// in-order-traversal and check if the res is sorted
func IsValid(root *TreeNode) bool {
	res := TraverseInOrder(root)
	if len(res) == 0 || len(res) == 1 {
		return true
	}
	prevVal := res[0]
	for _, v := range res[1:] {
		if v >= prevVal {
			prevVal = v
			continue
		}
		return false
	}
	return true
}

// leaf node -> set nil
// internal node -> move leftmost of right branch or rightmost of left branch
// examples
func (t *BinarySearchTree) DeleteNode(root *TreeNode, valToDelete int) *TreeNode {
	if root == nil {
		return nil
	}

	if valToDelete > root.Value {
		root.Right = t.DeleteNode(root.Right, valToDelete)
	}

	if valToDelete < root.Value {
		root.Left = t.DeleteNode(root.Left, valToDelete)
	}

	if valToDelete == root.Value {
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			toReplace := FindPreOrderSuccessor(root.Left)
			root.Value = toReplace.Value
			root.Left = t.DeleteNode(root.Left, toReplace.Value)
		}
	}
	return root
}

func FindPreOrderSuccessor(root *TreeNode) *TreeNode {
	res := root
	if res.Right != nil {
		res = res.Right
	}
	return res
}

func GetNodeCount(root *TreeNode) int {
	return len(TraverseInOrder(root))
}

// val < root -> Find in left
// val > root -> Find in right
func IsInTree(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if val < root.Value {
		return IsInTree(root.Left, val)
	} else if val > root.Value {
		return IsInTree(root.Right, val)
	} else {
		return true
	}
}

func GetMin(root *TreeNode) (int, error) {
	if root == nil {
		return 0, fmt.Errorf("empty tree")
	}

	res := root
	for res.Left != nil {
		res = res.Left
	}
	return res.Value, nil
}

func GetMax(root *TreeNode) (int, error) {
	if root == nil {
		return 0, fmt.Errorf("empty tree")
	}

	res := root
	for res.Right != nil {
		res = res.Right
	}
	return res.Value, nil
}

// if no children -> remove
// if has left -> recur to left
// if has right -> recur to right

// DeleteTree deletes all nodes of a tree
func DeleteTree(root **TreeNode) {
	// If root is nil, return
	if root == nil {
		return
	}

	// If root is a leaf node, delete it and return
	if (*root).Left == nil && (*root).Right == nil {
		*root = nil
		return
	}

	// If root has two children, delete both children and root, and return
	if (*root).Right != nil && (*root).Left != nil {
		DeleteTree(&(*root).Left)
		DeleteTree(&(*root).Right)
		*root = nil
		return
	}
}
