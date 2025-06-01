package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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

func printMain() {
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
