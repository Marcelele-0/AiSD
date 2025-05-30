@echo off
echo 🚀 Uruchamianie wszystkich programów drzew w osobnych terminalach...

echo 📂 Uruchamianie BST (ex1)...
start "BST - Binary Search Tree" cmd /k "cd ex1 && bst.exe"

echo 📂 Uruchamianie RB-BST (ex3)...
start "RB-BST - Red-Black Tree" cmd /k "cd ex3 && rbbst.exe"

echo 📂 Uruchamianie Splay Tree (ex5)...
start "Splay Tree" cmd /k "cd ex5 && splay.exe"

echo ✅ Wszystkie programy uruchomione w osobnych oknach!
pause
