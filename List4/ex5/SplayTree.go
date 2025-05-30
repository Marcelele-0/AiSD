package main

// SplayNode struktura reprezentująca węzeł w drzewie Splay
type SplayNode struct {
	key    int
	left   *SplayNode
	right  *SplayNode
	parent *SplayNode
}

// SplayTree struktura reprezentująca drzewo Splay
type SplayTree struct {
	root           *SplayNode
	comparisons    int
	pointerUpdates int
	rotations      int
	splays         int
	height         int
}

// NewSplayTree tworzy nowe puste drzewo Splay
func NewSplayTree() *SplayTree {
	return &SplayTree{
		root: nil,
	}
}

// leftRotate wykonuje rotację w lewo
func (st *SplayTree) leftRotate(x *SplayNode) {
	st.rotations++
	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		st.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
	st.pointerUpdates += 6
}

// rightRotate wykonuje rotację w prawo
func (st *SplayTree) rightRotate(x *SplayNode) {
	st.rotations++
	y := x.left
	x.left = y.right

	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		st.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y
	st.pointerUpdates += 6
}

// splay wykonuje operację splay na węźle
func (st *SplayTree) splay(node *SplayNode) {
	if node == nil {
		return
	}

	st.splays++
	for node.parent != nil {
		st.comparisons++
		if node.parent.parent == nil {
			// Zig case
			if node == node.parent.left {
				st.rightRotate(node.parent)
			} else {
				st.leftRotate(node.parent)
			}
		} else if node == node.parent.left && node.parent == node.parent.parent.left {
			// Zig-zig case (left)
			st.rightRotate(node.parent.parent)
			st.rightRotate(node.parent)
		} else if node == node.parent.right && node.parent == node.parent.parent.right {
			// Zig-zig case (right)
			st.leftRotate(node.parent.parent)
			st.leftRotate(node.parent)
		} else if node == node.parent.right && node.parent == node.parent.parent.left {
			// Zig-zag case (left-right)
			st.leftRotate(node.parent)
			st.rightRotate(node.parent)
		} else {
			// Zig-zag case (right-left)
			st.rightRotate(node.parent)
			st.leftRotate(node.parent)
		}
	}
}

// search wyszukuje węzeł o podanym kluczu i wykonuje splay
func (st *SplayTree) search(key int) *SplayNode {
	current := st.root
	var lastNode *SplayNode

	for current != nil {
		st.comparisons++
		lastNode = current
		if key == current.key {
			st.splay(current)
			return current
		} else if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	// Jeśli nie znaleziono klucza, splay ostatni odwiedzony węzeł
	if lastNode != nil {
		st.splay(lastNode)
	}
	return nil
}

// insert wstawia nowy węzeł do drzewa Splay
func (st *SplayTree) insert(key int) {
	if st.root == nil {
		st.root = &SplayNode{key: key}
		st.pointerUpdates++
		return
	}

	// Sprawdź czy klucz już istnieje
	found := st.search(key)
	if found != nil && found.key == key {
		return // Klucz już istnieje
	}

	// Wstaw nowy węzeł
	newNode := &SplayNode{key: key}
	current := st.root

	for {
		st.comparisons++
		if key < current.key {
			if current.left == nil {
				current.left = newNode
				newNode.parent = current
				st.pointerUpdates += 2
				break
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				newNode.parent = current
				st.pointerUpdates += 2
				break
			}
			current = current.right
		}
	}

	// Splay nowy węzeł do korzenia
	st.splay(newNode)
}

// deleteNode usuwa węzeł o podanym kluczu
func (st *SplayTree) deleteNode(key int) {
	node := st.search(key)
	if node == nil || node.key != key {
		return // Węzeł nie istnieje
	}

	st.pointerUpdates++

	if node.left == nil && node.right == nil {
		// Węzeł jest liściem
		st.root = nil
	} else if node.left == nil {
		// Węzeł ma tylko prawe dziecko
		st.root = node.right
		node.right.parent = nil
	} else if node.right == nil {
		// Węzeł ma tylko lewe dziecko
		st.root = node.left
		node.left.parent = nil
	} else {
		// Węzeł ma oba dzieci
		leftSubtree := node.left
		rightSubtree := node.right

		leftSubtree.parent = nil
		rightSubtree.parent = nil

		// Znajdź maksymalny element w lewym poddrzewie
		st.root = leftSubtree
		maxLeft := leftSubtree
		for maxLeft.right != nil {
			maxLeft = maxLeft.right
		}

		// Splay maksymalny element lewego poddrzewa
		st.splay(maxLeft)

		// Podłącz prawe poddrzewo
		st.root.right = rightSubtree
		rightSubtree.parent = st.root
		st.pointerUpdates += 2
	}
}

// resetStats resetuje wszystkie statystyki
func (st *SplayTree) resetStats() {
	st.comparisons = 0
	st.pointerUpdates = 0
	st.rotations = 0
	st.splays = 0
	st.height = 0
}

// calculateHeight oblicza wysokość drzewa
func (st *SplayTree) calculateHeight() int {
	return st.calculateHeightRecursive(st.root)
}

func (st *SplayTree) calculateHeightRecursive(node *SplayNode) int {
	if node == nil {
		return 0
	}
	leftHeight := st.calculateHeightRecursive(node.left)
	rightHeight := st.calculateHeightRecursive(node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// inorderTraversal przechodzi przez drzewo w porządku inorder
func (st *SplayTree) inorderTraversal() []int {
	var result []int
	st.inorderRecursive(st.root, &result)
	return result
}

func (st *SplayTree) inorderRecursive(node *SplayNode, result *[]int) {
	if node != nil {
		st.inorderRecursive(node.left, result)
		*result = append(*result, node.key)
		st.inorderRecursive(node.right, result)
	}
}
