import random
import numpy as np
import csv
import os
from multiprocessing import Pool, cpu_count
from utils import counters
from my_select import my_select
from random_select import randomized_select

# Parametry
n_values = list(range(1000, 50001, 100))
k_values = [1, 10, 100]
m = 50
output_file = "results/select_comparison.csv"
os.makedirs("results", exist_ok=True)

def run_experiment(args):
    n, k = args
    rand_comps = []
    rand_swaps = []
    det_comps = []
    det_swaps = []

    for _ in range(m):
        data = [random.randint(1, 10**6) for _ in range(n)]

        arr_rand = data.copy()
        counters.reset_counters()
        randomized_select(arr_rand, 0, n - 1, min(k, n))
        rand_comps.append(counters.comparison_count)
        rand_swaps.append(counters.swap_count)

        arr_det = data.copy()
        counters.reset_counters()
        my_select(arr_det, 0, n - 1, min(k, n))
        det_comps.append(counters.comparison_count)
        det_swaps.append(counters.swap_count)

    print(f"Done: n={n}, k={k}")
    return [
        [n, k, "randomized", np.mean(rand_comps), np.mean(rand_swaps)],
        [n, k, "deterministic", np.mean(det_comps), np.mean(det_swaps)]
    ]


if __name__ == "__main__":
    tasks = [(n, k) for k in k_values for n in n_values]

    with Pool(processes=cpu_count()) as pool:
        results = pool.map(run_experiment, tasks)

    # Flatten list of results
    flattened = [row for result in results for row in result]

    # Zapisz do pliku CSV
    with open(output_file, mode="w", newline="") as file:
        writer = csv.writer(file)
        writer.writerow(["n", "k", "algorithm", "avg_comparisons", "avg_swaps"])
        writer.writerows(flattened)

    print(f"Wyniki zapisane do {output_file}")
