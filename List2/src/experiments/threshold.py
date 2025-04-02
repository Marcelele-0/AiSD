import sys
import os
import random
import matplotlib.pyplot as plt

# Add the parent directory of `src` to the Python path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "../..")))

import src.utils.counters as counters
from hybrid_sort import hybrid_sort

def run_experiment(threshold_values, n_values, k=10):
    """
    Conducts experiments for different thresholds and data sizes n,
    performing k repetitions for each case. Collects data on the number of comparisons
    and swaps.
    """
    avg_comparisons = {threshold: [] for threshold in threshold_values}
    avg_swaps = {threshold: [] for threshold in threshold_values}

    for n in n_values:
        comparisons_per_threshold = {threshold: 0 for threshold in threshold_values}
        swaps_per_threshold = {threshold: 0 for threshold in threshold_values}

        print(f"\n--- Experiment for n = {n} ---")
        for i in range(k):
            # Generate a random array of size n
            array = [random.randint(1, 1000) for _ in range(n)]

            for threshold in threshold_values:
                array_copy = array.copy()
                counters.reset_counters()  # Reset counters
                hybrid_sort(array_copy, threshold)  # Sort with the given threshold
                comparisons = counters.comparison_count
                swaps = counters.swap_count
                comparisons_per_threshold[threshold] += comparisons
                swaps_per_threshold[threshold] += swaps

        # Calculate the average number of comparisons and swaps for the given n
        for threshold in threshold_values:
            avg_comp = comparisons_per_threshold[threshold] / k
            avg_swap = swaps_per_threshold[threshold] / k
            avg_comparisons[threshold].append(avg_comp)
            avg_swaps[threshold].append(avg_swap)
            print(f"For n = {n}, threshold {threshold}: average comparisons = {avg_comp}, average swaps = {avg_swap}")

    return avg_comparisons, avg_swaps

def plot_results(threshold_values, n_values, avg_comparisons, avg_swaps):
    """
    Creates plots of the average number of comparisons and swaps as a function of n.
    """
    print("\n--- Creating plots ---")
    plt.figure(figsize=(10, 6))

    # Plot for comparisons
    for threshold in threshold_values:
        plt.plot(n_values, avg_comparisons[threshold], marker='o', label=f'Threshold = {threshold}')
    plt.xlabel('Data size (n)')
    plt.ylabel('Average number of comparisons')
    plt.title('Comparisons as a function of threshold and data size')
    plt.legend()
    plt.grid(True)
    plt.savefig("comparison_plot.png")
    print("Comparison plot saved as comparison_plot.png")
    plt.show()

    # Plot for swaps
    plt.figure(figsize=(10, 6))
    for threshold in threshold_values:
        plt.plot(n_values, avg_swaps[threshold], marker='o', label=f'Threshold = {threshold}')
    plt.xlabel('Data size (n)')
    plt.ylabel('Average number of swaps')
    plt.title('Swaps as a function of threshold and data size')
    plt.legend()
    plt.grid(True)
    plt.savefig("swap_plot.png")
    print("Swap plot saved as swap_plot.png")
    plt.show()

def run_threshold_analysis():
    threshold_values = [6, 8, 10]
    n_values = [5, 6, 7, 8, 9, 10]
    k = 100

    avg_comparisons, avg_swaps = run_experiment(threshold_values, n_values, k)
    print("\n--- Average experiment results ---")
    print("Comparisons:", avg_comparisons)
    print("Swaps:", avg_swaps)

    plot_results(threshold_values, n_values, avg_comparisons, avg_swaps)

if __name__ == "__main__":
    run_threshold_analysis()
