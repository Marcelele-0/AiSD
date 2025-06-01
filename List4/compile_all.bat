@echo off
echo 🔨 Kompilowanie wszystkich programów drzew...

echo 📂 Kompilowanie BST (ex1)...
cd ex1
go mod init bst 2>nul
go build -o bst.exe
if exist bst.exe (echo ✅ BST skompilowany) else (echo ❌ Błąd kompilacji BST)
cd ..

echo 📂 Kompilowanie RB-BST (ex3)...
cd ex3
go mod init rbbst 2>nul
go build -o rbbst.exe
if exist rbbst.exe (echo ✅ RB-BST skompilowany) else (echo ❌ Błąd kompilacji RB-BST)
cd ..

echo 📂 Kompilowanie Splay Tree (ex5)...
cd ex5
go mod init splay 2>nul
go build -o splay.exe
if exist splay.exe (echo ✅ Splay Tree skompilowany) else (echo ❌ Błąd kompilacji Splay Tree)
cd ..

echo ✅ Kompilacja zakończona!
echo 🚀 Użyj 'run_all.bat' aby uruchomić wszystkie programy
pause
