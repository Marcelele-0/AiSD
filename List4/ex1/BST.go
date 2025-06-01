package main

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

// minValueNode znajduje najmniejszy węzeł w prawym poddrzewie
func (bst *BST) minValueNode(node *Node) *Node {
	for node != nil && node.left != nil {
		bst.comparisons++
		bst.pointerUpdates++
		node = node.left
	}
	return node
}

// heightOfTree zwraca wysokość drzewa (zoptymalizowana wersja rekurencyjna)
func (bst *BST) heightOfTree() int {
	if bst.root == nil {
		bst.height = 0
		return 0
	}

	bst.height = bst.heightRecursive(bst.root)
	return bst.height
}

// heightRecursive oblicza wysokość rekurencyjnie
func (bst *BST) heightRecursive(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := bst.heightRecursive(node.left)
	rightHeight := bst.heightRecursive(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
