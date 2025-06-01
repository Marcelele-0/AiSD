# Algorithms and Data Structures - List 4
## Tree Structures Implementation

## 🌳 Zawartość
- **ex1/** - Binary Search Tree (BST)
- **ex3/** - Red-Black Tree (RB-BST) 
- **ex5/** - Splay Tree

## 🚀 Kompilacja i uruchomienie

### BST (ex1)
```bash
cd ex1
go mod init bst
go build -o bst.exe
bst.exe
```

### Red-Black Tree (ex3)
```bash
cd ex3
go mod init rbbst
go build -o rbbst.exe
rbbst.exe
```

### Splay Tree (ex5)
```bash
cd ex5
go mod init splay
go build -o splay.exe
splay.exe
```

## 📋 Menu opcji
Każdy program oferuje:
1. **Demo** - wizualizacja z małymi danymi
2. **Benchmark** - testy wydajności
3. **Wyjście**

## � Pliki wynikowe
- `*_benchmark_full.json` - pełne wyniki testów
- `*_benchmark_short.json` - uśrednione wyniki

## 🔧 Wymagania
- Go 1.18+
- Windows/Linux/macOS

## 🚀 Skrypty automatyczne (Windows)

### Kompilacja wszystkich programów
```bash
compile_all.bat
```

### Uruchomienie w 3 terminalach
```bash
run_all.bat
```

**Efekt**: Otworzy się 3 osobne okna terminali:
- 🌲 BST - Binary Search Tree  
- 🔴 RB-BST - Red-Black Tree
- 🔄 Splay Tree

## 📊 Generowanie wykresów porównawczych

Po uruchomieniu benchmarków możesz wygenerować wykresy porównawcze:

### Ręczne uruchomienie (jeśli masz Python)
```bash
python generate_charts.py
```

**Wymagania**: Python z bibliotekami `matplotlib` i `numpy`
**Efekt**: Generuje wykresy PNG porównujące wydajność wszystkich trzech drzew:
- 🔍 Porównania liczby porównań
- 🔗 Porównania aktualizacji wskaźników  
- 🌳 Porównania wysokości drzew
- ⏱️ Porównania czasów wykonania
