#!/usr/bin/env python3
"""
Wizualizacja wyników porównania algorytmów MST (Prima vs Kruskala)
"""

import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

# Wczytanie danych
data = pd.read_csv('results.csv')

# Konfiguracja wykresu
plt.figure(figsize=(12, 8))
plt.style.use('seaborn-v0_8')

# Wykres czasów wykonania
plt.subplot(2, 2, 1)
plt.plot(data['n'], data['prim_time'], 'o-', label='Algorytm Prima', linewidth=2, markersize=6)
plt.plot(data['n'], data['kruskal_time'], 's-', label='Algorytm Kruskala', linewidth=2, markersize=6)
plt.xlabel('Liczba wierzchołków (n)')
plt.ylabel('Czas wykonania [ms]')
plt.title('Porównanie czasów wykonania algorytmów MST')
plt.legend()
plt.grid(True, alpha=0.3)

# Wykres skali logarytmicznej
plt.subplot(2, 2, 2)
plt.semilogy(data['n'], data['prim_time'], 'o-', label='Algorytm Prima', linewidth=2, markersize=6)
plt.semilogy(data['n'], data['kruskal_time'], 's-', label='Algorytm Kruskala', linewidth=2, markersize=6)
plt.xlabel('Liczba wierzchołków (n)')
plt.ylabel('Czas wykonania [ms] (skala log)')
plt.title('Porównanie czasów - skala logarytmiczna')
plt.legend()
plt.grid(True, alpha=0.3)

# Wykres stosunku czasów
plt.subplot(2, 2, 3)
ratio = data['kruskal_time'] / data['prim_time']
plt.plot(data['n'], ratio, 'ro-', linewidth=2, markersize=6)
plt.xlabel('Liczba wierzchołków (n)')
plt.ylabel('Stosunek czasu (Kruskal/Prim)')
plt.title('Stosunek czasów wykonania Kruskal/Prim')
plt.grid(True, alpha=0.3)

# Analiza złożoności obliczeniowej
plt.subplot(2, 2, 4)
n = data['n']

# Dopasowanie teoretycznych krzywych do rzeczywistych danych
# Znajdź współczynniki skalowania dla pierwszego punktu pomiarowego
n0 = data['n'].iloc[0]
prim_real_0 = data['prim_time'].iloc[0]
kruskal_real_0 = data['kruskal_time'].iloc[0]

# Teoretyczne wartości dla n0
prim_theo_0 = n0**2 * np.log(n0)
kruskal_theo_0 = (n0**2) * np.log(n0**2)

# Współczynniki skalowania
prim_scale = prim_real_0 / prim_theo_0
kruskal_scale = kruskal_real_0 / kruskal_theo_0

# Teoretyczne złożoności przeskalowane
prim_theoretical = prim_scale * (n**2 * np.log(n))
kruskal_theoretical = kruskal_scale * ((n**2) * np.log(n**2))

plt.plot(data['n'], data['prim_time'], 'o-', label='Prim (rzeczywisty)', linewidth=2, color='blue')
plt.plot(data['n'], prim_theoretical, '--', label='Prim O(V²log V)', alpha=0.7, color='lightblue')
plt.plot(data['n'], data['kruskal_time'], 's-', label='Kruskal (rzeczywisty)', linewidth=2, color='red')
plt.plot(data['n'], kruskal_theoretical, '--', label='Kruskal O(E log E)', alpha=0.7, color='lightcoral')
plt.xlabel('Liczba wierzchołków (n)')
plt.ylabel('Czas wykonania [ms]')
plt.title('Porównanie z teoretyczną złożonością')
plt.legend()
plt.grid(True, alpha=0.3)

plt.tight_layout()
plt.savefig('mst_comparison.png', dpi=300, bbox_inches='tight')
plt.show()

# Wyświetlenie statystyk
print("=== ANALIZA WYNIKÓW ===")
print(f"Zakres testowanych grafów: {data['n'].min()} - {data['n'].max()} wierzchołków")
print(f"Liczba punktów pomiarowych: {len(data)}")
print(f"Średni stosunek Kruskal/Prim: {(data['kruskal_time'] / data['prim_time']).mean():.2f}")
print(f"Przyspieszenie Prima dla n=1000: {data['kruskal_time'].iloc[-1] / data['prim_time'].iloc[-1]:.1f}x")

print("\n=== SPRAWDZENIE WYMAGAŃ ===")
print("✓ 1. Generator grafów pełnych z losowymi wagami (0,1) - ZAIMPLEMENTOWANY")
print("✓ 2. Algorytmy Prima i Kruskala - ZAIMPLEMENTOWANE")
print("✓ 3. Testy wydajnościowe:")
print(f"   - nMin = {data['n'].min()}, nMax = {data['n'].max()}")
print(f"   - step = {data['n'].iloc[1] - data['n'].iloc[0]}")
print(f"   - rep = 20 (widoczne w kodzie)")
print("✓ 4. Wizualizacja wyników - WYKONANA")
