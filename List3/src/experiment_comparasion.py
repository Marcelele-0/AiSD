import random
import numpy as np
import csv
import os
from multiprocessing import Pool, cpu_count
from utils import counters
from my_select import my_select
from random_select import randomized_select
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Parametry eksperymentu
n_values = list(range(1000, 50001, 100))  # np. co 1000
k_values = [1, 10, 100]
m = 50  # liczba powtórzeń
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

        # Algorytm Randomized Select
        arr_rand = data.copy()
        counters.reset_counters()
        randomized_select(arr_rand, 0, n - 1, min(k, n))
        rand_comps.append(counters.comparison_count)
        rand_swaps.append(counters.swap_count)

        # Algorytm Deterministic Select (MySelect)
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


def run_experiments():
    # Przygotowanie zadań
    tasks = [(n, k) for k in k_values for n in n_values]

    # Uruchomienie eksperymentów równolegle
    with Pool(processes=cpu_count()) as pool:
        results = pool.map(run_experiment, tasks)

    # Spłaszczenie listy wyników
    flattened = [row for result in results for row in result]

    # Zapisz wyniki do pliku CSV
    with open(output_file, mode="w", newline="") as file:
        writer = csv.writer(file)
        writer.writerow(["n", "k", "algorithm", "avg_comparisons", "avg_swaps"])
        writer.writerows(flattened)

    print(f"Wyniki zapisane do {output_file}")

def plot_results(file_path):
    # Wczytaj dane
    df = pd.read_csv(file_path)

    # Przekształć dane do wykresów
    df_melted = df.melt(
        id_vars=["n", "k", "algorithm"],
        value_vars=["avg_comparisons", "avg_swaps"],
        var_name="metric_type",
        value_name="value"
    )

    # Uproszczenie nazw metryk
    df_melted["metric_type"] = df_melted["metric_type"].map({
        "avg_comparisons": "Comparisons",
        "avg_swaps": "Swaps"
    })

    # Unikalne wartości k
    k_values = df["k"].unique()

    # Tworzenie wykresów
    os.makedirs("results/ex2", exist_ok=True)

    for k_val in k_values:
        subset = df_melted[df_melted["k"] == k_val]
        g = sns.relplot(
            data=subset,
            x="n", y="value", hue="algorithm", col="metric_type", kind="line",
            facet_kws={'sharey': False, 'sharex': True}
        )
        g.fig.suptitle(f"Select Comparison Metrics (k = {k_val})", y=1.05)
        g.set_axis_labels("n", "Average Count")
        g.set_titles(col_template="{col_name}")
        plt.tight_layout()
        plt.savefig(f"results/ex2/select_comparison_k{k_val}.png")
        plt.close()

    print("Plots saved to 'results/ex2/' folder.")

if __name__ == "__main__":
    run_experiments()
    plot_results(output_file)
