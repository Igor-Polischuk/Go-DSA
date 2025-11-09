package trees

type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

type binaryTreeNode[T Ordered] struct {
	Value T
	Left  *binaryTreeNode[T]
	Right *binaryTreeNode[T]
}

type BinaryTree[T Ordered] struct {
	Root *binaryTreeNode[T]
}

func NewBinaryTree[T Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (node *binaryTreeNode[T]) insertChild(value T) *binaryTreeNode[T] {
	if node == nil {
		return &binaryTreeNode[T]{Value: value}
	}

	if value < node.Value {
		node.Left = node.Left.insertChild(value)
	} else {
		node.Right = node.Right.insertChild(value)
	}

	return node
}

func (node *binaryTreeNode[T]) lookUp(value T) *binaryTreeNode[T] {
	if node == nil {
		return node
	}

	if node.Value == value {
		return node
	}

	if value < node.Value {
		return node.Left.lookUp(value)
	} else {
		return node.Right.lookUp(value)
	}

}

func (tree *BinaryTree[T]) InsertRecursive(value T) {
	tree.Root = tree.Root.insertChild(value)
}

func (tree *BinaryTree[T]) InsertIterative(value T) {
	node := binaryTreeNode[T]{Value: value}
	if tree.Root == nil {
		tree.Root = &node
		return
	}

	curr := tree.Root
	inserted := false

	for !inserted {
		if value < curr.Value {
			if curr.Left == nil {
				curr.Left = &node
				inserted = true
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = &node
				inserted = true
			}
			curr = curr.Right
		}
	}
}

func (tree *BinaryTree[T]) LookUp(value T) *binaryTreeNode[T] {
	curr := tree.Root

	for curr != nil {
		if curr.Value == value {
			return curr
		}

		if value < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}

func (tree *BinaryTree[T]) LookUpRecursive(value T) *binaryTreeNode[T] {
	return tree.Root.lookUp(value)
}

func (tree *BinaryTree[T]) Remove(value T) bool {
	if tree.Root == nil {
		return false
	}

	var parent *binaryTreeNode[T]
	targetNode := tree.Root
	for targetNode != nil {
		if value < targetNode.Value {
			parent = targetNode
			targetNode = targetNode.Left
			continue
		} else if value > targetNode.Value {
			parent = targetNode
			targetNode = targetNode.Right
			continue
		} else if targetNode.Value == value {
			// Target node has no children
			if targetNode.Left == nil && targetNode.Right == nil {
				// Delete root node without children (tree contains only root)
				if parent == nil {
					tree.Root = nil
				} else if targetNode.Value < parent.Value {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				return true
			}
			// Target node has only left child
			// No need to check if left is present because the previous condition ensures we have at least one child
			if targetNode.Right == nil {

				// Remove root node
				if parent == nil {
					tree.Root = targetNode.Left
				} else {
					if targetNode.Value < parent.Value {
						parent.Left = targetNode.Left

					} else {
						parent.Right = targetNode.Left
					}
				}
				return true
			}
			// Target node has only right child
			if targetNode.Left == nil {
				// Check if target is root
				if parent == nil {
					targetNode.Left.Left = tree.Root.Left
					targetNode.Left.Right = tree.Root.Right
					tree.Root = targetNode.Right
				} else {
					if targetNode.Value < parent.Value {
						parent.Left = targetNode.Right

					} else {
						parent.Right = targetNode.Right
					}

				}

				return true
			}

			// If none of the conditions above were true, it means we have both children present
			// Process the case when target node has both left and right child
			// Look for successor in right subtree (lowest value)

			// 1. When right node doesn't have a left child, that node is the successor
			if targetNode.Right.Left == nil {
				if parent == nil {
					targetNode.Right.Left = tree.Root.Left
					tree.Root = targetNode.Right
				} else if targetNode.Value < parent.Value {
					parent.Left = targetNode.Right
				} else {
					parent.Right = targetNode.Right
				}

				return true
			}

			// When right node has a left child, search for the lowest value (on the left) in this subtree
			successor := targetNode.Right.Left
			successorParent := targetNode.Right

			for successor.Left != nil {
				successorParent = successor
				successor = successor.Left
			}

			if parent == nil {
				successor.Left = tree.Root.Left
				successor.Right = tree.Root.Right
				tree.Root = successor

			} else if targetNode.Value < parent.Value {
				parent.Left = successor
			} else {
				parent.Right = successor
			}

			successorParent.Left = nil
			return true

		}

	}

	return false
}
