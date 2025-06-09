#!/bin/bash

# Skrypt testowy do sprawdzenia pipeline'u
echo "🧪 TEST PIPELINE - MST + INFO SPREAD"
echo "======================================"

# Przejdź do katalogu ex1_2 i sprawdź czy pliki istnieją
cd ex1_2

echo "📁 Sprawdzenie plików:"
echo -n "  mst.cpp: "
[[ -f "mst.cpp" ]] && echo "✅" || echo "❌"

echo -n "  info_spread.cpp: "
[[ -f "info_spread.cpp" ]] && echo "✅" || echo "❌"

echo -n "  mst_plot_results.py: "
[[ -f "mst_plot_results.py" ]] && echo "✅" || echo "❌"

echo -n "  info_spread_plot.py: "
[[ -f "info_spread_plot.py" ]] && echo "✅" || echo "❌"

echo ""
echo "🔧 Test kompilacji:"
echo -n "  Kompilacja mst.cpp: "
if g++ -std=c++17 -O2 mst.cpp -o mst_test 2>/dev/null; then
    echo "✅"
    rm -f mst_test
else
    echo "❌"
fi

echo -n "  Kompilacja info_spread.cpp: "
if g++ -std=c++17 -O2 info_spread.cpp -o info_spread_test 2>/dev/null; then
    echo "✅"
    rm -f info_spread_test
else
    echo "❌"
fi

echo ""
echo "🐍 Test Python:"
echo -n "  Python dostępny: "
if command -v python &> /dev/null; then
    echo "✅ ($(python --version))"
else
    echo "❌"
fi

echo -n "  Conda dostępna: "
if command -v conda &> /dev/null; then
    echo "✅"
else
    echo "❌"
fi

echo ""
echo "Wszystko gotowe do uruchomienia pipeline'u!"
echo "Uruchom: ./run_pipeline.sh"
