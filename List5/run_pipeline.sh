#!/bin/bash

# Kolory dla lepszej czytelności
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Funkcja do wyświetlania kolorowych komunikatów
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Sprawdź czy jesteśmy w odpowiednim katalogu
if [[ ! -d "ex1_2" ]]; then
    log_error "Nie znaleziono katalogu ex1_2. Uruchom skrypt z katalogu List5"
    exit 1
fi

echo "======================================"
echo "🚀 AUTOMATYCZNY PIPELINE MST + INFO SPREAD"
echo "======================================"

# KROK 1: Kompilacja i uruchomienie MST
log_info "Krok 1/5: Kompilacja programu MST..."
cd ex1_2
if ! g++ -std=c++17 -O2 mst.cpp -o mst; then
    log_error "Błąd kompilacji programu MST"
    exit 1
fi
log_success "Program MST skompilowany pomyślnie"

log_info "Krok 2/5: Uruchamianie analizy MST..."
echo "--------------------------------------"
if ! ./mst; then
    log_error "Błąd podczas wykonywania programu MST"
    exit 1
fi
echo "--------------------------------------"
log_success "Analiza MST zakończona pomyślnie"

# Sprawdź czy powstał plik wyników MST
if [[ ! -f "results.csv" ]]; then
    log_error "Nie znaleziono pliku wyników MST (results.csv)"
    exit 1
fi

# KROK 2: Kompilacja i uruchomienie Info Spread
log_info "Krok 3/5: Kompilacja programu Info Spread..."
if ! g++ -std=c++17 -O2 info_spread.cpp -o info_spread; then
    log_error "Błąd kompilacji programu Info Spread"
    exit 1
fi
log_success "Program Info Spread skompilowany pomyślnie"

log_info "Krok 4/5: Uruchamianie analizy rozprzestrzeniania informacji..."
echo "--------------------------------------"
if ! ./info_spread; then
    log_error "Błąd podczas wykonywania programu Info Spread"
    exit 1
fi
echo "--------------------------------------"
log_success "Analiza Info Spread zakończona pomyślnie"

# Sprawdź czy powstał plik wyników Info Spread
if [[ ! -f "info_spread_results.csv" ]]; then
    log_error "Nie znaleziono pliku wyników Info Spread (info_spread_results.csv)"
    exit 1
fi

# KROK 3: Aktywacja conda environment i uruchomienie skryptów plotowania
log_info "Krok 5/5: Aktywacja środowiska conda i generowanie wykresów..."

# Sprawdź czy conda jest dostępna
if ! command -v conda &> /dev/null; then
    log_error "Conda nie jest zainstalowana lub nie jest w PATH"
    exit 1
fi

# Aktywuj środowisko conda
log_info "Aktywacja środowiska conda 'counterfactuals'..."
if ! eval "$(conda shell.bash hook)" || ! conda activate counterfactuals; then
    log_error "Nie udało się aktywować środowiska conda 'counterfactuals'"
    log_warning "Spróbuję uruchomić skrypty bez aktywacji środowiska..."
fi

# Uruchom skrypt plotowania MST
log_info "Generowanie wykresów MST..."
if [[ -f "mst_plot_results.py" ]]; then
    if python mst_plot_results.py; then
        log_success "Wykres MST wygenerowany pomyślnie"
    else
        log_error "Błąd podczas generowania wykresu MST"
    fi
else
    log_warning "Nie znaleziono skryptu mst_plot_results.py"
fi

# Uruchom skrypt plotowania Info Spread
log_info "Generowanie wykresów Info Spread..."
if [[ -f "info_spread_plot.py" ]]; then
    if python info_spread_plot.py; then
        log_success "Wykres Info Spread wygenerowany pomyślnie"
    else
        log_error "Błąd podczas generowania wykresu Info Spread"
    fi
else
    log_warning "Nie znaleziono skryptu info_spread_plot.py"
fi

echo "======================================"
log_success "🎉 PIPELINE ZAKOŃCZONY POMYŚLNIE!"
echo "======================================"

# Podsumowanie plików wynikowych
echo ""
log_info "📊 Pliki wynikowe:"
if [[ -f "results.csv" ]]; then
    echo "  ✅ MST results: ex1_2/results.csv"
fi
if [[ -f "info_spread_results.csv" ]]; then
    echo "  ✅ Info Spread results: ex1_2/info_spread_results.csv"
fi

# Sprawdź pliki wykresów
echo ""
log_info "📈 Wygenerowane wykresy:"
if [[ -f "mst_comparison.png" ]]; then
    echo "  ✅ MST plot: ex1_2/mst_comparison.png"
fi
if [[ -f "info_spread_plot.png" ]]; then
    echo "  ✅ Info Spread plot: ex1_2/info_spread_plot.png"
fi

echo ""
log_info "Pipeline zakończony. Wszystkie wyniki są gotowe!"
