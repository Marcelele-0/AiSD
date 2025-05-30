@echo off
echo 🚀 Uruchamianie benchmarków wszystkich drzew...

echo 📂 Uruchamianie benchmarku BST (ex1)...
start "BST Benchmark" cmd /k "cd ex1 && echo 2 | bst.exe"

echo 📂 Uruchamianie benchmarku RB-BST (ex3)...
start "RB-BST Benchmark" cmd /k "cd ex3 && echo 2 | rbbst.exe"

echo 📂 Uruchamianie benchmarku Splay Tree (ex5)...
start "Splay Tree Benchmark" cmd /k "cd ex5 && echo 2 | splay.exe"

echo ✅ Wszystkie benchmarki uruchomione w osobnych oknach!
echo 📊 Sprawdź pliki JSON z wynikami w folderach ex1/, ex3/, ex5/
pause
