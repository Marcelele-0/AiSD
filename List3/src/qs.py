import sys
import random
import time
import utils.counters as counters
from sorts.quick_sort import quick_sort
from sorts.dual_pivot_qs import dual_pivot_quick_sort
from sorts.select_qs import quick_sort_select as quick_sort_with_select
from sorts.select_dpqs import select_dual_pivot_quick_sort as dual_pivot_quick_sort_with_select

# Funkcja do pomiaru czasu i statystyk dla różnych algorytmów
def measure_algorithm(algorithm, array: list, low: int, high: int):
    counters.reset_counters()  # Resetujemy liczniki
    start_time = time.time()  # Mierzymy czas przed
    algorithm(array, low, high)  # Wywołujemy algorytm
    end_time = time.time()  # Mierzymy czas po
    execution_time = end_time - start_time  # Obliczamy czas wykonania
    
    return execution_time, counters.comparison_count, counters.swap_count  # Zwracamy czas i statystyki

# Funkcja do uruchomienia testów
def run_tests():
    # Ustalamy parametry testu
    n_values = [100, 200, 500, 1000, 5000]  # Przykładowe rozmiary danych
    k = 50  # Przykładowy k-ty element
    
    for n in n_values:
        print(f"\nTesting for n = {n}:")

        # Generujemy dane wejściowe
        array = [random.randint(0, 1000) for _ in range(n)]
        
        # Testujemy oryginalny quick sort
        print("\n--- Original QuickSort ---")
        original_array = array[:]
        time_qs, comparisons_qs, swaps_qs = measure_algorithm(quick_sort, original_array, 0, n - 1)
        print(f"Time: {time_qs:.6f}s, Comparisons: {comparisons_qs}, Swaps: {swaps_qs}")
        print(f"Array after sorting (QuickSort): {original_array[:30]}...")  # Print first 30 elements

        # Testujemy quick sort z my_select (QuickSort with SELECT)
        print("\n--- QuickSort with my_select ---")
        array_with_select_qs = array[:]
        time_qs_select, comparisons_qs_select, swaps_qs_select = measure_algorithm(quick_sort_with_select, array_with_select_qs, 0, n - 1)
        print(f"Time: {time_qs_select:.6f}s, Comparisons: {comparisons_qs_select}, Swaps: {swaps_qs_select}")
        print(f"Array after sorting (QuickSort with SELECT): {array_with_select_qs[:30]}...")  # Print first 30 elements
        
        # Testujemy dual pivot quick sort (klasyczny)
        print("\n--- Dual Pivot QuickSort ---")
        array_with_dpqs = array[:]
        time_dpqs, comparisons_dpqs, swaps_dpqs = measure_algorithm(dual_pivot_quick_sort, array_with_dpqs, 0, n - 1)
        print(f"Time: {time_dpqs:.6f}s, Comparisons: {comparisons_dpqs}, Swaps: {swaps_dpqs}")
        print(f"Array after sorting (Dual Pivot QuickSort): {array_with_dpqs[:30]}...")  # Print first 30 elements
        
        # Testujemy dual pivot quick sort z my_select (Dual Pivot QuickSort with SELECT)
        print("\n--- Dual Pivot QuickSort with my_select ---")
        array_with_dpqs_select = array[:]
        time_dpqs_select, comparisons_dpqs_select, swaps_dpqs_select = measure_algorithm(dual_pivot_quick_sort_with_select, array_with_dpqs_select, 0, n - 1)
        print(f"Time: {time_dpqs_select:.6f}s, Comparisons: {comparisons_dpqs_select}, Swaps: {swaps_dpqs_select}")
        print(f"Array after sorting (Dual Pivot QuickSort with SELECT): {array_with_dpqs_select[:30]}...")  # Print first 30 elements

        # Sprawdzamy, czy oba algorytmy dają ten sam wynik
        if array_with_dpqs_select == sorted(array_with_dpqs_select):
            print("The array is correctly sorted for Dual Pivot QuickSort with SELECT.")
        else:
            print("Sorting error in Dual Pivot QuickSort with SELECT.")

if __name__ == "__main__":
    run_tests()
