import sys
import time
import random
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

# Recursive binary search with comparison counting
def binary_search(arr, v, low, high, comps=0):
    if low > high:
        return False, comps + 1  # final comparison
    mid = (low + high) // 2
    comps += 1
    if arr[mid] == v:
        return True, comps
    elif v < arr[mid]:
        return binary_search(arr, v, low, mid - 1, comps)
    else:
        return binary_search(arr, v, mid + 1, high, comps)

# Demonstration on small data
print("=== Demonstracja dla małych danych ===")
small_arr = [1, 3, 5, 7, 9]
for v in [1, 5, 10]:
    found, c = binary_search(small_arr, v, 0, len(small_arr) - 1)
    print(f"Search {v:2d}:", "Found" if found else "Not found", f"Comparisons: {c}")

# Experimental setup
results = []
ns = list(range(1000, 100001, 1000))
for n in ns:
    arr = list(range(n))
    # scenarios: first, middle, last, absent
    scenarios = {
        'first': 0,
        'middle': n // 2,
        'last': n - 1,
        'absent': -1
    }
    for scen, v in scenarios.items():
        t0 = time.perf_counter()
        found, comps = binary_search(arr, v, 0, n - 1)
        t1 = time.perf_counter()
        results.append({'n': n, 'scenario': scen, 'comps': comps, 'time': t1 - t0})
    
    # random-average scenario
    comps_list = []
    times = []
    for _ in range(10):
        v = random.choice(arr)
        t0 = time.perf_counter()
        _, c = binary_search(arr, v, 0, n - 1)
        t1 = time.perf_counter()
        comps_list.append(c)
        times.append(t1 - t0)
    results.append({
        'n': n,
        'scenario': 'random_avg',
        'comps': sum(comps_list) / len(comps_list),
        'time': sum(times) / len(times)
    })

df = pd.DataFrame(results)
df['log2n'] = np.log2(df['n'])

# Estimate constant factor for each scenario: average of comps/log2(n)
factors = df.copy()
factors['factor_comps'] = factors['comps'] / factors['log2n']
factors['factor_time'] = factors['time'] / factors['log2n']
estimates = factors.groupby('scenario').agg({
    'factor_comps': 'mean',
    'factor_time': 'mean'
}).reset_index()

# Plot comparisons vs n
plt.figure()
for scen in df['scenario'].unique():
    subset = df[df['scenario'] == scen]
    plt.plot(subset['n'], subset['comps'], label=scen)
plt.xlabel('n')
plt.ylabel('Comparisons')
plt.title('Liczba porównań w zależności od n')
plt.legend()
plt.show()

# Plot time vs n
plt.figure()
for scen in df['scenario'].unique():
    subset = df[df['scenario'] == scen]
    plt.plot(subset['n'], subset['time'], label=scen)
plt.xlabel('n')
plt.ylabel('Czas wykonania [s]')
plt.title('Czas wykonania w zależności od n')
plt.legend()
plt.show()

# Print estimated factors
print("=== Oszacowane czynniki O(1) dla porównań i czasu ===")
print(estimates)
