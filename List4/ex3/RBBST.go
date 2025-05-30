package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
		if newNode.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = y
	rb.pointerUpdates++

	if y == rb.nil {
		rb.root = newNode
	} else if newNode.key < y.key {
		y.left = newNode
	} else {
		y.right = newNode
	}

	rb.pointerUpdates++

	// Napraw właściwości Red-Black
	rb.insertFixup(newNode)
}

// insertFixup naprawia właściwości Red-Black po wstawieniu
func (rb *RB_BST) insertFixup(z *RBNode) {
	for rb.isRed(z.parent) {
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
		} else if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	return rb.nil
}

// minimum znajduje węzeł z minimalnym kluczem w poddrzewie
func (rb *RB_BST) minimum(node *RBNode) *RBNode {
	for node.left != rb.nil {
		node = node.left
	}
	return node
}

// transplant zastępuje poddrzewo zakorzenione w u poddrzewem zakorzenionym w v
func (rb *RB_BST) transplant(u, v *RBNode) {
	if u.parent == rb.nil {
		rb.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
	rb.pointerUpdates++
}

// deleteNode usuwa węzeł o podanym kluczu
func (rb *RB_BST) deleteNode(key int) {
	z := rb.search(key)
	if z == rb.nil {
		return
	}

	y := z
	yOriginalColor := y.color
	var x *RBNode

	if z.left == rb.nil {
		x = z.right
		rb.transplant(z, z.right)
	} else if z.right == rb.nil {
		x = z.left
		rb.transplant(z, z.left)
	} else {
		y = rb.minimum(z.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent == z {
			x.parent = y
		} else {
			rb.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		rb.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		rb.deleteFixup(x)
	}
}

// deleteFixup naprawia właściwości Red-Black po usunięciu
func (rb *RB_BST) deleteFixup(x *RBNode) {
	for x != rb.root && !rb.isRed(x) {
		if x == x.parent.left {
			w := x.parent.right
			if rb.isRed(w) {
				rb.setColor(w, BLACK)
				rb.setColor(x.parent, RED)
				rb.leftRotate(x.parent)
				w = x.parent.right
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
				}

				w.color = x.parent.color
				rb.setColor(x.parent, BLACK)
				rb.setColor(w.right, BLACK)
				rb.leftRotate(x.parent)
				x = rb.root
			}
		} else {
			w := x.parent.left
			if rb.isRed(w) {
				rb.setColor(w, BLACK)
				rb.setColor(x.parent, RED)
				rb.rightRotate(x.parent)
				w = x.parent.left
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
	return rb.calculateHeightRecursive(rb.root)
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

// printTree wypisuje wizualizację drzewa z kolorowymi oznaczeniami
func (rb *RB_BST) printTree(node *RBNode, prefix string, isLast bool, isRoot bool) {
	if node == rb.nil || node == nil {
		return
	}

	// Wybierz kolor terminala na podstawie koloru węzła
	var colorCode, resetCode string
	if node.color == RED {
		colorCode = "\033[31m" // Czerwony tekst w terminalu
	} else {
		colorCode = "\033[90m" // Szary tekst w terminalu
	}
	resetCode = "\033[0m" // Reset koloru

	// Wypisz obecny węzeł
	var connector string
	if isRoot {
		connector = ""
	} else if isLast {
		connector = "└── "
	} else {
		connector = "├── "
	}

	fmt.Printf("%s%s%s[%d]%s\n", prefix, connector, colorCode, node.key, resetCode)

	// Przygotuj prefix dla dzieci
	var childPrefix string
	if isRoot {
		childPrefix = prefix
	} else if isLast {
		childPrefix = prefix + "    "
	} else {
		childPrefix = prefix + "│   "
	}

	// Sprawdź czy są dzieci
	hasLeft := node.left != rb.nil
	hasRight := node.right != rb.nil

	// Wypisz dzieci
	if hasLeft || hasRight {
		if hasLeft {
			rb.printTree(node.left, childPrefix, !hasRight, false)
		}
		if hasRight {
			rb.printTree(node.right, childPrefix, true, false)
		}
	}
}

// printCompactView drukuje kompaktowy widok drzewa
func (rb *RB_BST) printCompactView() {
	values := rb.inorderTraversal()
	if len(values) == 0 {
		return
	}

	fmt.Printf("🔴⚫ Red-Black Tree (inorder): [")
	for i, val := range values {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d", val)
	}
	fmt.Printf("]\n")
}

// printTreeDetailed drukuje szczegółowy widok drzewa z wszystkimi informacjami
func (rb *RB_BST) printTreeDetailed() {
	rb.height = rb.calculateHeight()

	fmt.Printf("┌─ 🔴⚫ Red-Black Tree State ─┐\n")
	fmt.Printf("│ Height: %-19d │\n", rb.height)
	fmt.Printf("│ Comparisons: %-14d │\n", rb.comparisons)
	fmt.Printf("│ Pointer Updates: %-10d │\n", rb.pointerUpdates)
	fmt.Printf("│ Rotations: %-16d │\n", rb.rotations)
	fmt.Printf("└─────────────────────────────┘\n")
	if rb.root == rb.nil {
		fmt.Println("🌿 Empty Red-Black Tree")
	} else {
		rb.printTree(rb.root, "", true, true)
	}

	rb.printCompactView()
	fmt.Println()
}

// printOperation drukuje informacje o wykonywanej operacji
func printRBOperation(operation string, key int, step int) {
	var emoji string
	switch operation {
	case "INSERT":
		emoji = "➕"
	case "DELETE":
		emoji = "➖"
	case "SEARCH":
		emoji = "🔍"
	default:
		emoji = "🔧"
	}
	fmt.Printf("%s %s %d (Step %d)\n", emoji, operation, key, step)
	fmt.Println(strings.Repeat("─", 40))
}

// printSeparator drukuje separator z tytułem
func printRBSeparator(title string) {
	line := strings.Repeat("█", 60)
	fmt.Printf("\n%s\n", line)
	fmt.Printf("█%s█\n", strings.Repeat(" ", 58))
	fmt.Printf("█  %-55s █\n", title)
	fmt.Printf("█%s█\n", strings.Repeat(" ", 58))
	fmt.Printf("%s\n\n", line)
}

// Przykładowe funkcje do testowania - można je przenieść do osobnego pliku main
// lub użyć w testach jednostkowych

// RunRBDemo uruchamia demonstrację Red-Black Tree
func main() {
	rb := NewRB_BST()
	n := 30 // Reduced for better readability in demo

	printRBSeparator("CASE 1: INCREASING SEQUENCE INSERT + RANDOM DELETE")

	// Przypadek 1: Wstawianie rosnącego ciągu i usuwanie losowej permutacji
	fmt.Println("📈 Inserting increasing sequence (1 to", n, ") into Red-Black Tree:")

	for i := 1; i <= n; i++ {
		printRBOperation("INSERT", i, i)
		rb.insert(i)
		rb.printTreeDetailed()
	}

	// Przygotowanie losowej permutacji do usuwania
	keys := make([]int, n)
	for i := 1; i <= n; i++ {
		keys[i-1] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	fmt.Println("\n🗑️  Deleting in random order:")
	for idx, key := range keys {
		printRBOperation("DELETE", key, idx+1)
		rb.deleteNode(key)
		rb.printTreeDetailed()
	}

	printRBSeparator("CASE 2: RANDOM INSERT + RANDOM DELETE")

	// Przypadek 2: Wstawianie losowej permutacji i usuwanie losowej permutacji
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	rb = NewRB_BST() // Reset RB Tree
	fmt.Println("🎲 Inserting random permutation into Red-Black Tree:")
	for idx, key := range keys {
		printRBOperation("INSERT", key, idx+1)
		rb.insert(key)
		rb.printTreeDetailed()
	}

	// Kolejna losowa permutacja do usywania
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	fmt.Println("\n🗑️  Deleting in random order:")
	for idx, key := range keys {
		printRBOperation("DELETE", key, idx+1)
		rb.deleteNode(key)
		rb.printTreeDetailed()
	}

	printRBSeparator("FINAL STATISTICS")
	fmt.Printf("🔴⚫ Red-Black Tree Final Statistics:\n")
	fmt.Printf("Total comparisons: %d\n", rb.comparisons)
	fmt.Printf("Total pointer updates: %d\n", rb.pointerUpdates)
	fmt.Printf("Total rotations: %d\n", rb.rotations)
	fmt.Printf("Final height: %d\n", rb.calculateHeight())
	fmt.Println("\n🎉 Red-Black BST operations completed successfully!")
}
