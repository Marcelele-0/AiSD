package main

// Color typ reprezentujący kolor węzła w drzewie Red-Black
type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

// RBNode struktura reprezentująca węzeł w drzewie Red-Black
type RBNode struct {
	key    int
	color  Color
	left   *RBNode
	right  *RBNode
	parent *RBNode
}

// RB_BST struktura reprezentująca drzewo Red-Black Binary Search Tree
type RB_BST struct {
	root           *RBNode
	nil            *RBNode // sentinel node
	comparisons    int
	pointerUpdates int
	rotations      int
	height         int
}

// NewRB_BST tworzy nowe drzewo Red-Black
func NewRB_BST() *RB_BST {
	nil := &RBNode{color: BLACK}
	return &RB_BST{
		root: nil,
		nil:  nil,
	}
}

// isRed sprawdza czy węzeł jest czerwony
func (rb *RB_BST) isRed(node *RBNode) bool {
	if node == rb.nil || node == nil {
		return false
	}
	return node.color == RED
}

// setColor ustawia kolor węzła
func (rb *RB_BST) setColor(node *RBNode, color Color) {
	if node != rb.nil && node != nil {
		node.color = color
	}
}

// leftRotate wykonuje rotację w lewo
func (rb *RB_BST) leftRotate(x *RBNode) {
	rb.rotations++
	y := x.right
	x.right = y.left

	if y.left != rb.nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == rb.nil {
		rb.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
	rb.pointerUpdates += 3
}

// rightRotate wykonuje rotację w prawo
func (rb *RB_BST) rightRotate(x *RBNode) {
	rb.rotations++
	y := x.left
	x.left = y.right

	if y.right != rb.nil {
		y.right.parent = x
	}

	y.parent = x.parent

	if x.parent == rb.nil {
		rb.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y
	rb.pointerUpdates += 3
}

// insert wstawia nowy węzeł do drzewa Red-Black
func (rb *RB_BST) insert(key int) {
	newNode := &RBNode{
		key:    key,
		color:  RED,
		left:   rb.nil,
		right:  rb.nil,
		parent: rb.nil,
	}

	var y *RBNode = rb.nil
	x := rb.root
	// Znajdź miejsce do wstawienia
	for x != rb.nil {
		rb.comparisons++
		y = x
		rb.comparisons++
		if newNode.key < x.key {
			rb.pointerUpdates++
			x = x.left
		} else {
			rb.pointerUpdates++
			x = x.right
		}
	}
	newNode.parent = y
	rb.pointerUpdates++

	rb.comparisons++
	if y == rb.nil {
		rb.pointerUpdates++
		rb.root = newNode
	} else {
		rb.comparisons++
		if newNode.key < y.key {
			rb.pointerUpdates++
			y.left = newNode
		} else {
			rb.pointerUpdates++
			y.right = newNode
		}
	}

	rb.pointerUpdates++

	// Napraw właściwości Red-Black
	rb.insertFixup(newNode)
}

// insertFixup naprawia właściwości Red-Black po wstawieniu
func (rb *RB_BST) insertFixup(z *RBNode) {
	for rb.isRed(z.parent) {
		if z.parent.parent == rb.nil {
			break
		}
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if rb.isRed(y) {
				// Przypadek 1
				rb.setColor(z.parent, BLACK)
				rb.setColor(y, BLACK)
				rb.setColor(z.parent.parent, RED)
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					// Przypadek 2
					z = z.parent
					rb.leftRotate(z)
				}
				// Przypadek 3
				rb.setColor(z.parent, BLACK)
				rb.setColor(z.parent.parent, RED)
				rb.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if rb.isRed(y) {
				// Przypadek 1 (symetryczny)
				rb.setColor(z.parent, BLACK)
				rb.setColor(y, BLACK)
				rb.setColor(z.parent.parent, RED)
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					// Przypadek 2 (symetryczny)
					z = z.parent
					rb.rightRotate(z)
				}
				// Przypadek 3 (symetryczny)
				rb.setColor(z.parent, BLACK)
				rb.setColor(z.parent.parent, RED)
				rb.leftRotate(z.parent.parent)
			}
		}
	}
	rb.setColor(rb.root, BLACK)
}

// search szuka węzła o podanym kluczu
func (rb *RB_BST) search(key int) *RBNode {
	current := rb.root
	for current != rb.nil {
		rb.comparisons++
		if key == current.key {
			return current
		}
		rb.comparisons++
		if key < current.key {
			rb.pointerUpdates++
			current = current.left
		} else {
			rb.pointerUpdates++
			current = current.right
		}
	}
	return rb.nil
}

// minimum znajduje węzeł z minimalnym kluczem w poddrzewie
func (rb *RB_BST) minimum(node *RBNode) *RBNode {
	for node.left != rb.nil {
		rb.comparisons++
		rb.pointerUpdates++
		node = node.left
	}
	return node
}

// transplant zastępuje poddrzewo zakorzenione w u poddrzewem zakorzenionym w v
func (rb *RB_BST) transplant(u, v *RBNode) {
	rb.comparisons++
	if u.parent == rb.nil {
		rb.pointerUpdates++
		rb.root = v
	} else {
		rb.comparisons++
		if u == u.parent.left {
			rb.pointerUpdates++
			u.parent.left = v
		} else {
			rb.pointerUpdates++
			u.parent.right = v
		}
	}
	rb.pointerUpdates++
	v.parent = u.parent
}

// deleteNode usuwa węzeł o podanym kluczu
func (rb *RB_BST) deleteNode(key int) {
	z := rb.search(key)
	rb.comparisons++
	if z == rb.nil {
		return
	}

	y := z
	yOriginalColor := y.color
	var x *RBNode

	rb.comparisons++
	if z.left == rb.nil {
		rb.pointerUpdates++
		x = z.right
		rb.transplant(z, z.right)
	} else {
		rb.comparisons++
		if z.right == rb.nil {
			rb.pointerUpdates++
			x = z.left
			rb.transplant(z, z.left)
		} else {
			y = rb.minimum(z.right)
			yOriginalColor = y.color
			rb.pointerUpdates++
			x = y.right

			rb.comparisons++
			if y.parent == z {
				rb.pointerUpdates++
				if x != rb.nil {
					x.parent = y
				}
			} else {
				rb.transplant(y, y.right)
				rb.pointerUpdates += 2
				y.right = z.right
				y.right.parent = y
			}

			rb.transplant(z, y)
			rb.pointerUpdates += 3
			y.left = z.left
			y.left.parent = y
			y.color = z.color
		}
	}

	rb.comparisons++
	if yOriginalColor == BLACK {
		rb.deleteFixup(x)
	}
}

// deleteFixup naprawia właściwości Red-Black po usunięciu
func (rb *RB_BST) deleteFixup(x *RBNode) {
	// Sprawdź czy x nie jest nil lub sentinel
	if x == rb.nil || x == nil {
		return
	}
	for x != rb.root && !rb.isRed(x) {
		if x.parent == rb.nil || x.parent == nil {
			break
		}
		if x == x.parent.left {
			w := x.parent.right
			if w == nil || w == rb.nil {
				break
			}

			if rb.isRed(w) {
				rb.setColor(w, BLACK)
				rb.setColor(x.parent, RED)
				rb.leftRotate(x.parent)
				w = x.parent.right
				if w == nil || w == rb.nil {
					break
				}
			}

			if !rb.isRed(w.left) && !rb.isRed(w.right) {
				rb.setColor(w, RED)
				x = x.parent
			} else {
				if !rb.isRed(w.right) {
					rb.setColor(w.left, BLACK)
					rb.setColor(w, RED)
					rb.rightRotate(w)
					w = x.parent.right
					if w == nil || w == rb.nil {
						break
					}
				}

				w.color = x.parent.color
				rb.setColor(x.parent, BLACK)
				rb.setColor(w.right, BLACK)
				rb.leftRotate(x.parent)
				x = rb.root
			}
		} else {
			w := x.parent.left
			if w == nil || w == rb.nil {
				break
			}

			if rb.isRed(w) {
				rb.setColor(w, BLACK)
				rb.setColor(x.parent, RED)
				rb.rightRotate(x.parent)
				w = x.parent.left
				if w == nil || w == rb.nil {
					break
				}
			}

			if !rb.isRed(w.right) && !rb.isRed(w.left) {
				rb.setColor(w, RED)
				x = x.parent
			} else {
				if !rb.isRed(w.left) {
					rb.setColor(w.right, BLACK)
					rb.setColor(w, RED)
					rb.leftRotate(w)
					w = x.parent.left
					if w == nil || w == rb.nil {
						break
					}
				}

				w.color = x.parent.color
				rb.setColor(x.parent, BLACK)
				rb.setColor(w.left, BLACK)
				rb.rightRotate(x.parent)
				x = rb.root
			}
		}
	}
	rb.setColor(x, BLACK)
}

// calculateHeight oblicza wysokość drzewa
func (rb *RB_BST) calculateHeight() int {
	if rb.root == rb.nil {
		rb.height = 0
		return 0
	}
	rb.height = rb.calculateHeightRecursive(rb.root)
	return rb.height
}

func (rb *RB_BST) calculateHeightRecursive(node *RBNode) int {
	if node == rb.nil || node == nil {
		return 0
	}
	leftHeight := rb.calculateHeightRecursive(node.left)
	rightHeight := rb.calculateHeightRecursive(node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// inorderTraversal przechodzi przez drzewo w porządku inorder
func (rb *RB_BST) inorderTraversal() []int {
	var result []int
	rb.inorderRecursive(rb.root, &result)
	return result
}

func (rb *RB_BST) inorderRecursive(node *RBNode, result *[]int) {
	if node != rb.nil && node != nil {
		rb.inorderRecursive(node.left, result)
		*result = append(*result, node.key)
		rb.inorderRecursive(node.right, result)
	}
}

// resetStats resetuje wszystkie statystyki
func (rb *RB_BST) resetStats() {
	rb.comparisons = 0
	rb.pointerUpdates = 0
	rb.rotations = 0
	rb.height = 0
}
