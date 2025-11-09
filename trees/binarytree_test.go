package trees

import (
	"testing"
)

func TestTreeNode(t *testing.T) {
	t.Run("Node creation", func(t *testing.T) {
		node := binaryTreeNode[int]{}
		t.Log("Node: ", node)

		if node.Value != 0 && node.Left == nil && node.Right == nil {
			t.Errorf("expected node value to be 0, left and right to be nil, got %d, %v, %v", node.Value, node.Left, node.Right)
		}

		node.Value = 2

		node.insertChild(3)
		node.insertChild(1)

		if node.Left.Value != 1 && node.Right.Value != 3 {
			t.Errorf("expected left value to be 1, right value to be 3, got %d, %d", node.Left.Value, node.Right.Value)
		}
	})

	t.Run("Node insertion", func(t *testing.T) {
		node := binaryTreeNode[int]{Value: 3}
		node.insertChild(2)
		node.insertChild(6)
		node.insertChild(1)
		node.insertChild(5)

		if node.Left.Value != 2 {
			t.Errorf("expected left value to be 2, got %d", node.Left.Value)
		}

		if node.Right.Value != 6 {
			t.Errorf("expected right value to be 4, got %d", node.Right.Value)
		}

		if node.Left.Left.Value != 1 {
			t.Errorf("expected left left value to be 1, got %d", node.Left.Left.Value)
		}

		if node.Right.Left.Value != 5 {
			t.Errorf("expected right left value to be 5, got %d", node.Right.Left.Value)
		}

	})
}

func TestBinaryTreeInsertion(t *testing.T) {

	treeInsertionTestHelper := func(t *testing.T, tree *BinaryTree[int]) {
		t.Helper()
		if tree.Root.Value != 9 {
			t.Errorf("expected root value to be 9, got %d", tree.Root.Value)
		}

		if tree.Root.Left.Value != 4 {
			t.Errorf("expected left value to be 4, got %d", tree.Root.Left.Value)
		}

		if tree.Root.Left.Left.Value != 1 {
			t.Errorf("expected left left value to be 1, got %d", tree.Root.Left.Left.Value)
		}

		if tree.Root.Left.Right.Value != 6 {
			t.Errorf("expected left right value to be 6, got %d", tree.Root.Left.Right.Value)
		}

		if tree.Root.Right.Value != 20 {
			t.Errorf("expected right value to be 20, got %d", tree.Root.Right.Value)
		}

		if tree.Root.Right.Left.Value != 15 {
			t.Errorf("expected right left value to be 15, got %d", tree.Root.Right.Left.Value)
		}
	}

	tests := []struct {
		name       string
		insertFunc func(*BinaryTree[int], int)
	}{
		{
			name: "Test recursive insertion",
			insertFunc: func(tree *BinaryTree[int], val int) {
				tree.InsertRecursive(val)
			},
		},
		{
			name: "Test iterative insertion",
			insertFunc: func(tree *BinaryTree[int], val int) {
				tree.InsertIterative(val)
			},
		},
	}

	valuesToInsert := []int{9, 4, 6, 20, 170, 15, 1}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBinaryTree[int]()

			for _, val := range valuesToInsert {
				tt.insertFunc(tree, val)
			}

			treeInsertionTestHelper(t, tree)
		})
	}

}

func TestBinaryTreeLookup(t *testing.T) {
	tree := NewBinaryTree[int]()
	valuesToInsert := []int{9, 4, 6, 20, 170, 15, 1}

	for _, value := range valuesToInsert {
		tree.InsertRecursive(value)
	}

	for _, value := range valuesToInsert {
		resIterative := tree.LookUp(value)

		if resIterative.Value != value {
			t.Errorf("expected result to be %d, got %v when looking up %d iteratively", value, resIterative, value)
		}

		resRecursive := tree.LookUpRecursive(value)

		if resRecursive.Value != value {
			t.Errorf("expected result to be %d, got %v when looking up %d recursively", value, resRecursive, value)
		}

	}

	res := tree.LookUp(44)

	if res != nil {
		t.Errorf("expected result to be nil when looking up 44 iteratively, got %v", res)
	}

	resRecursive := tree.LookUpRecursive(44)

	if resRecursive != nil {
		t.Errorf("expected result to be nil when looking up 44 recursively, got %v", resRecursive)
	}

}

func TestBinaryTreeDeletion(t *testing.T) {
	tree := NewBinaryTree[int]()
	valuesToInsert := []int{9, 4, 6, 20, 170, 15, 1}

	for _, value := range valuesToInsert {
		tree.InsertRecursive(value)
	}

	//       9
	// 	   /   \
	//    4     20
	//   / \    / \
	//  1   6  15  170

	res := tree.Remove(9)

	if !res {
		t.Errorf("expected result to be true when removing 9, got %v", res)
	}

	if tree.Root.Value != 15 {
		t.Errorf("expected root value to be 15, got %d", tree.Root.Value)
	}

	if tree.Root.Right.Left != nil {
		t.Errorf("expected right left to be nil, got %v", tree.Root.Right.Left)
	}

	//       15
	// 	   /   \
	//    4     20
	//   / \      \
	//  1   6     170

	res = tree.Remove(20)

	if !res {
		t.Errorf("expected result to be true when removing 20, got %v", res)
	}

	if tree.Root.Right.Value != 170 {
		t.Errorf("expected right value to be 170, got %d", tree.Root.Right.Value)
	}

	if tree.Root.Right.Left != nil && tree.Root.Right.Right != nil {
		t.Errorf("expected right left and right right to be nil, got %v, %v", tree.Root.Right.Left, tree.Root.Right.Right)
	}

	//       15
	// 	   /   \
	//    4     170
	//   / \
	//  1   6

	res = tree.Remove(1)

	if !res {
		t.Errorf("expected result to be true when removing 1, got %v", res)
	}

	if tree.Root.Left.Left != nil {
		t.Errorf("expected left left to be nil, got %v", tree.Root.Left.Left)
	}

	//       15
	// 	   /   \
	//    4     170
	//     \
	//      6

	res = tree.Remove(4)

	if !res {
		t.Errorf("expected result to be true when removing 4, got %v", res)
	}

	if tree.Root.Left.Value != 6 {
		t.Errorf("expected left value to be 6, got %d", tree.Root.Left.Value)
	}

	tree.Remove(15)
	tree.Remove(170)

	if tree.Root.Value != 6 {
		t.Errorf("expected root value to be 6, got %d", tree.Root.Value)
	}

	if tree.Root.Left != nil {
		t.Errorf("expected left to be nil, got %v", tree.Root.Left)
	}

	res = tree.Remove(44)

	if res {
		t.Errorf("expected result to be false when removing 44, got %v", res)
	}

	tree.Remove(6)

	if tree.Root != nil {
		t.Errorf("expected root to be nil, got %v", tree.Root)
	}

}
