import time
import matplotlib.pyplot as plt
from typing import List
import random
import utils.counters as counters
from insertion_sort import insertion_sort
from quick_sort import quick_sort
from hybrid_sort import hybrid_sort


def generate_random_array(n: int) -> List[int]:
    """
    Generates a random array of integers of length n.
    """
    return [random.randint(1, 1000) for _ in range(n)]


def run_experiment(n: int, k: int, algorithm: str) -> List[int]:
    """
    Runs a sorting experiment for a given algorithm and returns comparison and swap counts.
    """
    comparison_counts = []
    swap_counts = []
    times = []
    
    for _ in range(k):
        array = generate_random_array(n)
        counters.reset_counters()  # Reset counters for each experiment
        
        start_time = time.time()

        if algorithm == "quick_sort":
            quick_sort(array, 0, n - 1)
        elif algorithm == "insertion_sort":
            insertion_sort(n, array)
        elif algorithm == "hybrid_sort":
            hybrid_sort(array)
        else:
            raise ValueError(f"Unknown algorithm: {algorithm}")

        end_time = time.time()

        comparison_counts.append(counters.comparison_count)
        swap_counts.append(counters.swap_count)
        times.append(end_time - start_time)

    avg_comparisons = sum(comparison_counts) / k
    avg_swaps = sum(swap_counts) / k
    avg_time = sum(times) / k

    return avg_comparisons, avg_swaps, avg_time


def run_sorting_benchmarking(k: int, sizes: List[int]) -> None:
    """
    Runs the experiment for all algorithms and plots the results.
    """
    algorithms = ["quick_sort", "insertion_sort", "hybrid_sort"]
    
    avg_comparisons = {algo: [] for algo in algorithms}
    avg_swaps = {algo: [] for algo in algorithms}
    avg_times = {algo: [] for algo in algorithms}

    for n in sizes:
        print(f"Running experiments for n = {n}...")
        for algo in algorithms:
            avg_comparisons[algo].append(run_experiment(n, k, algo)[0])
            avg_swaps[algo].append(run_experiment(n, k, algo)[1])
            avg_times[algo].append(run_experiment(n, k, algo)[2])

    # Plotting
    plt.figure(figsize=(10, 6))

    # Plot comparisons
    plt.subplot(3, 1, 1)
    for algo in algorithms:
        plt.plot(sizes, avg_comparisons[algo], label=f"{algo} comparisons")
    plt.xlabel("Array size (n)")
    plt.ylabel("Average comparisons")
    plt.legend()

    # Plot swaps
    plt.subplot(3, 1, 2)
    for algo in algorithms:
        plt.plot(sizes, avg_swaps[algo], label=f"{algo} swaps")
    plt.xlabel("Array size (n)")
    plt.ylabel("Average swaps")
    plt.legend()

    # Plot times
    plt.subplot(3, 1, 3)
    for algo in algorithms:
        plt.plot(sizes, avg_times[algo], label=f"{algo} time (s)")
    plt.xlabel("Array size (n)")
    plt.ylabel("Average time (s)")
    plt.legend()

    plt.tight_layout()
    plt.show()


if __name__ == "__main__":
    sizes = [10, 20, 30, 40, 50]  # Sizes of arrays to experiment with
    k_values = [1, 10, 100]  # Number of repetitions for each experiment
    
    for k in k_values:
        print(f"\nRunning experiments for k = {k}...")
        run_sorting_benchmarking(k, sizes)
