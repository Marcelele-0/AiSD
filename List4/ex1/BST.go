package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Node struktura reprezentująca węzeł w drzewie
type Node struct {
	key   int
	left  *Node
	right *Node
}

// BST struktura reprezentująca drzewo binarne
type BST struct {
	root           *Node
	comparisons    int
	pointerUpdates int
	height         int
}

// insert wstawia nowy węzeł do drzewa
func (bst *BST) insert(key int) {
	bst.root = bst.insertRecursive(bst.root, key)
}

// insertRecursive rekurencyjnie wstawia nowy węzeł do odpowiedniego miejsca
func (bst *BST) insertRecursive(node *Node, key int) *Node {
	bst.comparisons++
	if node == nil {
		bst.pointerUpdates++
		return &Node{key: key}
	}

	bst.comparisons++
	if key < node.key {
		bst.pointerUpdates++
		node.left = bst.insertRecursive(node.left, key)
	} else if key > node.key {
		bst.pointerUpdates++
		node.right = bst.insertRecursive(node.right, key)
	}
	// else key == node.key, nie dodajemy duplikatów

	return node
}

// deleteNode usuwa węzeł o podanym kluczu
func (bst *BST) deleteNode(key int) {
	bst.root = bst.deleteRecursive(bst.root, key)
}

// deleteRecursive rekurencyjnie usuwa węzeł z drzewa
func (bst *BST) deleteRecursive(node *Node, key int) *Node {
	bst.comparisons++
	if node == nil {
		return nil
	}

	bst.comparisons++
	if key < node.key {
		bst.pointerUpdates++
		node.left = bst.deleteRecursive(node.left, key)
	} else if key > node.key {
		bst.pointerUpdates++
		node.right = bst.deleteRecursive(node.right, key)
	} else {
		// Węzeł do usunięcia
		bst.pointerUpdates++
		bst.comparisons++
		if node.left == nil {
			bst.pointerUpdates++
			return node.right
		} else if node.right == nil {
			bst.pointerUpdates++
			return node.left
		}

		// Węzeł z dwoma dziećmi
		temp := bst.minValueNode(node.right)
		bst.pointerUpdates++
		node.key = temp.key
		bst.pointerUpdates++
		node.right = bst.deleteRecursive(node.right, temp.key)
	}

	return node
}

// resetStats resetuje wszystkie statystki
func (bst *BST) resetStats() {
	bst.comparisons = 0
	bst.pointerUpdates = 0
	bst.height = 0
}

// search szuka węzła o podanym kluczu i śledzi porównania
func (bst *BST) search(key int) *Node {
	return bst.searchRecursive(bst.root, key)
}

// searchRecursive rekurencyjnie szuka węzła z śledzeniem porównań
func (bst *BST) searchRecursive(node *Node, key int) *Node {
	bst.comparisons++
	if node == nil || node.key == key {
		return node
	}

	bst.comparisons++
	if key < node.key {
		return bst.searchRecursive(node.left, key)
	}
	return bst.searchRecursive(node.right, key)
}

// minValueNode znajduje najmniejszy węzeł w prawym poddrzewie
func (bst *BST) minValueNode(node *Node) *Node {
	for node != nil && node.left != nil {
		bst.comparisons++
		bst.pointerUpdates++
		node = node.left
	}
	return node
}

// heightOfTree zwraca wysokość drzewa (wersja iteracyjna, BFS)
func (bst *BST) heightOfTree() int {
	if bst.root == nil {
		bst.height = 0
		return 0
	}

	type nodeLevel struct {
		node  *Node
		level int
	}

	queue := []nodeLevel{{bst.root, 1}}
	maxHeight := 0

	for len(queue) > 0 {
		nl := queue[0]
		queue = queue[1:]

		if nl.level > maxHeight {
			maxHeight = nl.level
		}
		if nl.node.left != nil {
			queue = append(queue, nodeLevel{nl.node.left, nl.level + 1})
		}
		if nl.node.right != nil {
			queue = append(queue, nodeLevel{nl.node.right, nl.level + 1})
		}
	}

	bst.height = maxHeight
	return maxHeight
}

// printTree wypisuje drzewo w ulepszonej formie tekstowej z lepszą czytelnością
func (bst *BST) printTree(node *Node, prefix string, isLast bool, isRoot bool) {
	if node == nil {
		return
	}

	// Wypisz obecny węzeł
	var connector string
	if isRoot {
		connector = "🌳 "
	} else if isLast {
		connector = "└── "
	} else {
		connector = "├── "
	}

	fmt.Printf("%s%s[%d]\n", prefix, connector, node.key)

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
			bst.printTree(node.left, childPrefix, !hasRight, false)
		}
		if hasRight {
			bst.printTree(node.right, childPrefix, true, false)
		}
	}
}

// printCompact wypisuje drzewo w kompaktowej formie (in-order traversal)
func (bst *BST) printCompact() {
	if bst.root == nil {
		fmt.Println("Empty tree")
		return
	}

	fmt.Print("In-order: ")
	bst.inOrderTraversal(bst.root)
	fmt.Println()
}

// inOrderTraversal wykonuje przechodzenie in-order przez drzewo
func (bst *BST) inOrderTraversal(node *Node) {
	if node != nil {
		bst.inOrderTraversal(node.left)
		fmt.Printf("[%d] ", node.key)
		bst.inOrderTraversal(node.right)
	}
}

// printTreeDetailed wypisuje drzewo z dodatkowymi informacjami
func (bst *BST) printTreeDetailed() {
	if bst.root == nil {
		fmt.Println("🌳 Tree is empty")
		return
	}

	fmt.Println("┌─── Binary Search Tree ───")
	fmt.Printf("│ 🌱 Root: [%d]\n", bst.root.key)
	fmt.Printf("│ 📏 Height: %d\n", bst.heightOfTree())
	fmt.Printf("│ 🔍 Comparisons: %d\n", bst.comparisons)
	fmt.Printf("│ 🔗 Pointer updates: %d\n", bst.pointerUpdates)
	fmt.Println("└───────────────────────────")
	fmt.Println()
	bst.printTree(bst.root, "", true, true)
	bst.printCompact()
	fmt.Println()
}

// printOperation wypisuje informacje o operacji z ładnym formatowaniem
func printOperation(operation string, key int, step int) {
	fmt.Printf("\n" + strings.Repeat("═", 50) + "\n")
	fmt.Printf("Step %d: %s %d\n", step, operation, key)
	fmt.Printf(strings.Repeat("═", 50) + "\n")
}

// printSeparator wypisuje separator między sekcjami
func printSeparator(title string) {
	line := strings.Repeat("█", 60)
	fmt.Printf("\n%s\n", line)
	fmt.Printf("█%s█\n", strings.Repeat(" ", 58))
	titlePadding := (58 - len(title)) / 2
	fmt.Printf("█%s%s%s█\n",
		strings.Repeat(" ", titlePadding),
		title,
		strings.Repeat(" ", 58-titlePadding-len(title)))
	fmt.Printf("█%s█\n", strings.Repeat(" ", 58))
	fmt.Printf("%s\n\n", line)
}

func main() {
	bst := BST{}
	n := 15 // Reduced for better readability in demo

	printSeparator("CASE 1: INCREASING SEQUENCE INSERT + RANDOM DELETE")

	// Przypadek 1: Wstawianie rosnącego ciągu i usuwanie losowej permutacji
	fmt.Println("📈 Inserting increasing sequence (1 to", n, "):")

	for i := 1; i <= n; i++ {
		printOperation("INSERT", i, i)
		bst.insert(i)
		bst.printTreeDetailed()
	}

	fmt.Printf("\n📊 CASE 1 INSERT STATISTICS:\n")
	fmt.Printf("   • Comparisons: %d\n", bst.comparisons)
	fmt.Printf("   • Pointer updates: %d\n", bst.pointerUpdates)
	fmt.Printf("   • Final height: %d\n", bst.heightOfTree())

	// Przygotowanie losowej permutacji do usuwania
	keys := make([]int, n)
	for i := 1; i <= n; i++ {
		keys[i-1] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	// Reset stats for delete operations
	insertStats := struct {
		comparisons    int
		pointerUpdates int
	}{bst.comparisons, bst.pointerUpdates}

	bst.resetStats()

	fmt.Println("\n🗑️  Deleting in random order:")
	for idx, key := range keys {
		printOperation("DELETE", key, idx+1)
		bst.deleteNode(key)
		bst.printTreeDetailed()
	}

	fmt.Printf("\n📊 CASE 1 FINAL STATISTICS:\n")
	fmt.Printf("   • Insert operations - Comparisons: %d, Pointer updates: %d\n",
		insertStats.comparisons, insertStats.pointerUpdates)
	fmt.Printf("   • Delete operations - Comparisons: %d, Pointer updates: %d\n",
		bst.comparisons, bst.pointerUpdates)
	fmt.Printf("   • Total operations - Comparisons: %d, Pointer updates: %d\n",
		insertStats.comparisons+bst.comparisons, insertStats.pointerUpdates+bst.pointerUpdates)

	printSeparator("CASE 2: RANDOM INSERT + RANDOM DELETE")

	// Przypadek 2: Wstawianie losowej permutacji i usuwanie losowej permutacji
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	bst = BST{} // Reset BST
	fmt.Println("🎲 Inserting random permutation:")
	for idx, key := range keys {
		printOperation("INSERT", key, idx+1)
		bst.insert(key)
		bst.printTreeDetailed()
	}

	fmt.Printf("\n📊 CASE 2 INSERT STATISTICS:\n")
	fmt.Printf("   • Comparisons: %d\n", bst.comparisons)
	fmt.Printf("   • Pointer updates: %d\n", bst.pointerUpdates)
	fmt.Printf("   • Final height: %d\n", bst.heightOfTree())

	// Kolejna losowa permutacja do usuwania
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	// Reset stats for delete operations
	insertStats2 := struct {
		comparisons    int
		pointerUpdates int
	}{bst.comparisons, bst.pointerUpdates}

	bst.resetStats()

	fmt.Println("\n🗑️  Deleting in random order:")
	for idx, key := range keys {
		printOperation("DELETE", key, idx+1)
		bst.deleteNode(key)
		bst.printTreeDetailed()
	}

	fmt.Printf("\n📊 CASE 2 FINAL STATISTICS:\n")
	fmt.Printf("   • Insert operations - Comparisons: %d, Pointer updates: %d\n",
		insertStats2.comparisons, insertStats2.pointerUpdates)
	fmt.Printf("   • Delete operations - Comparisons: %d, Pointer updates: %d\n",
		bst.comparisons, bst.pointerUpdates)
	fmt.Printf("   • Total operations - Comparisons: %d, Pointer updates: %d\n",
		insertStats2.comparisons+bst.comparisons, insertStats2.pointerUpdates+bst.pointerUpdates)

	printSeparator("FINAL STATISTICS COMPARISON")
	fmt.Printf("📈 CASE 1 (Ordered Insert): %d comparisons, %d pointer updates\n",
		insertStats.comparisons+bst.comparisons, insertStats.pointerUpdates+bst.pointerUpdates)
	fmt.Printf("🎲 CASE 2 (Random Insert): %d comparisons, %d pointer updates\n",
		insertStats2.comparisons+bst.comparisons, insertStats2.pointerUpdates+bst.pointerUpdates)
	fmt.Println("\n🎉 BST operations completed successfully!")
}
