import random
import numpy as np
import time
import csv
import os
from utils import counters
from my_select import my_select
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt
from multiprocessing import Pool, cpu_count


# Funkcja generująca wykresy
def plot_results(file_path):
    # Wczytaj dane
    df = pd.read_csv(file_path)

    # Przekształć dane do wykresów
    df_melted = df.melt(id_vars=["n", "group_size", "k"], 
                        value_vars=["avg_comparisons", "avg_swaps", "avg_time_sec"], 
                        var_name="metric", value_name="value")

    # Dodaj kolumnę typu metryki
    df_melted["type"] = df_melted["metric"].apply(lambda x: "Comparisons" if "comparisons" in x else
                                                  ("Swaps" if "swaps" in x else "Time"))

    # Wykresy: Comparisons, Swaps, Time
    metrics = ["Comparisons", "Swaps", "Time"]

    for metric in metrics:
        subset = df_melted[df_melted["type"] == metric]

        # Osobny wykres dla każdej wartości k
        for k_val in sorted(subset["k"].unique()):
            k_subset = subset[subset["k"] == k_val]
            plt.figure(figsize=(10, 6))

            for group in sorted(k_subset["group_size"].unique()):
                group_subset = k_subset[k_subset["group_size"] == group]
                sns.lineplot(data=group_subset, x="n", y="value", label=f'group size = {group}', errorbar=None)

            plt.title(f'{metric} for Different Group Sizes (k={k_val})')
            plt.xlabel("n (Input Size)")
            plt.ylabel(f'Average {metric}')
            plt.legend(title="Group Size")
            plt.tight_layout()
            plt.savefig(f"results/ex3/{metric.lower()}_group_size_k{k_val}.png")
            plt.close()

    print("Wykresy zapisane w folderze 'results/ex3'")


# Parametry eksperymentu
n_values = list(range(1000, 50001, 1000))  # np. co 1000
group_sizes = [3, 5, 7, 9]
k_values = [15, 20, 30]  # Różne wartości k
m = 30  # liczba powtórzeń

# Plik wynikowy
output_file = "results/select_group_size_experiment.csv"
os.makedirs("results", exist_ok=True)
os.makedirs("results/ex3", exist_ok=True)

# Nagłówek
header = ["n", "group_size", "k", "avg_comparisons", "avg_swaps", "avg_time_sec"]

# Funkcja wykonująca eksperyment dla danej kombinacji n, group_size, k
def run_single_experiment(args):
    n, group_size, k = args
    comps, swaps, times = [], [], []

    for _ in range(m):
        data = [random.randint(1, 10**6) for _ in range(n)]
        arr = data.copy()

        counters.reset_counters()
        start = time.time()
        my_select(arr, 0, n - 1, k, group_size=group_size)
        end = time.time()

        comps.append(counters.comparison_count)
        swaps.append(counters.swap_count)
        times.append(end - start)

    return [n, group_size, k, np.mean(comps), np.mean(swaps), np.mean(times)]


# Wykonaj eksperyment i zapisz wyniki
def run_experiment():
    with open(output_file, mode="w", newline="") as file:
        writer = csv.writer(file)
        writer.writerow(header)

        # Przygotowanie zadań
        tasks = [(n, group_size, k) for group_size in group_sizes for n in n_values for k in k_values]

        # Uruchom eksperymenty równolegle
        with Pool(processes=cpu_count()) as pool:
            results = pool.map(run_single_experiment, tasks)

        # Zapisz wyniki do pliku
        for result in results:
            writer.writerow(result)

    print(f"Zapisano wyniki do {output_file}")


# Uruchom eksperyment lub generowanie wykresów na podstawie wyników
if __name__ == "__main__":
    #run_experiment()
    plot_results(output_file)
