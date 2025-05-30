package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
	fmt.Printf("%s", "\n"+strings.Repeat("═", 50)+"\n")
	fmt.Printf("Step %d: %s %d\n", step, operation, key)
	fmt.Printf("%s", strings.Repeat("═", 50)+"\n")
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

func printMain() {
	bst := BST{}
	n := 30 // Reduced for better readability in demo

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
