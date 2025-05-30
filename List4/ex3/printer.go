package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
func printMain() {
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
