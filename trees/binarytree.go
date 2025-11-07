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
