package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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

// findMin znajduje węzeł o najmniejszej wartości
func (st *SplayTree) findMin(node *SplayNode) *SplayNode {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
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

// printTree wypisuje wizualizację drzewa z oznaczeniami Splay Tree
func (st *SplayTree) printTree(node *SplayNode, prefix string, isLast bool, isRoot bool) {
	if node == nil {
		return
	}

	// Kolor dla Splay Tree (niebieski dla wyróżnienia)
	colorCode := "\033[34m" // Niebieski tekst w terminalu
	resetCode := "\033[0m"  // Reset koloru

	// Wypisz obecny węzeł
	var connector string
	if isRoot {
		connector = "🌳 "
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
	hasLeft := node.left != nil
	hasRight := node.right != nil

	// Wypisz dzieci
	if hasLeft || hasRight {
		if hasLeft {
			st.printTree(node.left, childPrefix, !hasRight, false)
		}
		if hasRight {
			st.printTree(node.right, childPrefix, true, false)
		}
	}
}

// printCompactView drukuje kompaktowy widok drzewa
func (st *SplayTree) printCompactView() {
	values := st.inorderTraversal()
	if len(values) == 0 {
		return
	}

	fmt.Printf("🔷 Splay Tree (inorder): [")
	for i, val := range values {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d", val)
	}
	fmt.Printf("]\n")
}

// printTreeDetailed drukuje szczegółowy widok drzewa z wszystkimi informacjami
func (st *SplayTree) printTreeDetailed() {
	st.height = st.calculateHeight()

	fmt.Printf("┌── 🔷 Splay Tree State ──┐\n")
	fmt.Printf("│ Height: %-15d │\n", st.height)
	fmt.Printf("│ Comparisons: %-10d │\n", st.comparisons)
	fmt.Printf("│ Pointer Updates: %-6d │\n", st.pointerUpdates)
	fmt.Printf("│ Rotations: %-12d │\n", st.rotations)
	fmt.Printf("│ Splays: %-15d │\n", st.splays)
	fmt.Printf("└─────────────────────────┘\n")

	if st.root == nil {
		fmt.Println("🌿 Empty Splay Tree")
	} else {
		st.printTree(st.root, "", true, true)
	}

	st.printCompactView()
	fmt.Println()
}

// printOperation drukuje informacje o wykonywanej operacji
func printSplayOperation(operation string, key int, step int) {
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
func printSplaySeparator(title string) {
	line := strings.Repeat("█", 60)
	fmt.Printf("\n%s\n", line)
	fmt.Printf("█%s█\n", strings.Repeat(" ", 58))
	fmt.Printf("█  %-55s █\n", title)
	fmt.Printf("█%s█\n", strings.Repeat(" ", 58))
	fmt.Printf("%s\n\n", line)
}

func main() {
	st := NewSplayTree()
	n := 30 // Reduced for better readability in demo

	printSplaySeparator("CASE 1: INCREASING SEQUENCE INSERT + RANDOM DELETE")

	// Przypadek 1: Wstawianie rosnącego ciągu i usuwanie losowej permutacji
	fmt.Println("📈 Inserting increasing sequence (1 to", n, ") into Splay Tree:")

	for i := 1; i <= n; i++ {
		printSplayOperation("INSERT", i, i)
		st.insert(i)
		st.printTreeDetailed()
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
		printSplayOperation("DELETE", key, idx+1)
		st.deleteNode(key)
		st.printTreeDetailed()
	}

	printSplaySeparator("CASE 2: RANDOM INSERT + RANDOM DELETE")

	// Przypadek 2: Wstawianie losowej permutacji i usuwanie losowej permutacji
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	st = NewSplayTree() // Reset Splay Tree
	fmt.Println("🎲 Inserting random permutation into Splay Tree:")
	for idx, key := range keys {
		printSplayOperation("INSERT", key, idx+1)
		st.insert(key)
		st.printTreeDetailed()
	}

	// Kolejna losowa permutacja do usuwania
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	fmt.Println("\n🗑️  Deleting in random order:")
	for idx, key := range keys {
		printSplayOperation("DELETE", key, idx+1)
		st.deleteNode(key)
		st.printTreeDetailed()
	}

	printSplaySeparator("FINAL STATISTICS")
	fmt.Printf("🔷 Splay Tree Final Statistics:\n")
	fmt.Printf("Total comparisons: %d\n", st.comparisons)
	fmt.Printf("Total pointer updates: %d\n", st.pointerUpdates)
	fmt.Printf("Total rotations: %d\n", st.rotations)
	fmt.Printf("Total splays: %d\n", st.splays)
	fmt.Printf("Final height: %d\n", st.calculateHeight())
	fmt.Println("\n🎉 Splay Tree operations completed successfully!")
}
