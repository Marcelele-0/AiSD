import numpy as np
import random
import matplotlib.pyplot as plt
from insertion_sort import insertion_sort
from quick_sort import quick_sort
from dual_pivot_qs import dual_pivot_quick_sort
from my_sort import my_merge_sort
from merge_sort import merge_sort
from hybrid_sort import hybrid_sort
import utils.counters as counters
import concurrent.futures

def run_experiment(sort_function, n: int, k: int, threshold: int = 5) -> tuple:
    """
    Run sorting experiment for a given sorting function and return the average number of comparisons and swaps.
    
    Parameters:
        sort_function: Sorting function to be tested (insertion_sort, quick_sort, hybrid_sort, etc.)
        n (int): Number of elements to sort.
        k (int): Number of repetitions to average.
        threshold (int): Threshold value for hybrid_sort (if applicable).
    
    Returns:
        tuple: (average_comparisons, average_swaps)
    """
    comparisons_list = []
    swaps_list = []
    
    # Generate the initial array once (values from 1 to n)
    original_array = list(range(1, n + 1))
    
    for _ in range(k):
        # Copy and shuffle the array for each repetition
        array = original_array.copy()
        random.shuffle(array)
        counters.reset_counters()
        
        if sort_function == hybrid_sort:
            sort_function(array, 0, n - 1, threshold)
        else:
            sort_function(array, 0, n - 1)
        
        comparisons_list.append(counters.comparison_count)
        swaps_list.append(counters.swap_count)
    
    avg_comparisons = np.mean(comparisons_list)
    avg_swaps = np.mean(swaps_list)
    
    return avg_comparisons, avg_swaps


def run_experiment_for_size(size: int, k: int) -> dict:
    """
    Run experiments for a given size for all sorting algorithms.
    
    Returns a dictionary with keys for each algorithm.
    """
    print(f"Running experiment for size {size} with {k} repetitions...")
    results = {}
    results['quick_sort'] = run_experiment(quick_sort, size, k)
    results['hybrid_sort'] = run_experiment(hybrid_sort, size, k, threshold=6)
    results['dual_pivot'] = run_experiment(dual_pivot_quick_sort, size, k)
    results['merge_sort'] = run_experiment(merge_sort, size, k)
    results['my_merge_sort'] = run_experiment(my_merge_sort, size, k)
    return {size: results}


def experiment_for_big_sizes_parallel() -> dict:
    sizes = [i for i in range(1000, 50001, 1000)]
    k = 10 # Number of repetitions for each size
    combined_results = {
        'quick_sort': [],
        'hybrid_sort': [],
        'dual_pivot': [],
        'merge_sort': [],
        'my_merge_sort': []
    }
    
    with concurrent.futures.ProcessPoolExecutor(max_workers=20) as executor:
        # Submit a task for each size
        futures = {executor.submit(run_experiment_for_size, size, k): size for size in sizes}
        for future in concurrent.futures.as_completed(futures):
            size = futures[future]
            try:
                result_for_size = future.result()  # result_for_size is a dict: {size: {alg: (comp, swaps), ...}}
                for alg, (comp, swaps) in result_for_size[size].items():
                    combined_results[alg].append((size, comp, swaps))
            except Exception as exc:
                print(f"Size {size} generated an exception: {exc}")
    
    return combined_results


def experiment_for_various_sizes():
    sizes = [10, 20, 30, 40, 50]
    k = 1000
    results = {
        'insertion_sort': [],
        'quick_sort': [],
        'hybrid_sort': [],
        'dual_pivot': [],
        'merge_sort': [],
        'my_merge_sort': []
    }
    
    for size in sizes:
        print(f"\nRunning experiment for n={size}, k={k}...")
        results['insertion_sort'].append((size, *run_experiment(insertion_sort, size, k)))
        results['quick_sort'].append((size, *run_experiment(quick_sort, size, k)))
        results['hybrid_sort'].append((size, *run_experiment(hybrid_sort, size, k, threshold=10)))
        results['dual_pivot'].append((size, *run_experiment(dual_pivot_quick_sort, size, k)))
        results['merge_sort'].append((size, *run_experiment(merge_sort, size, k)))
        results['my_merge_sort'].append((size, *run_experiment(my_merge_sort, size, k)))
    
    return results


def plot_results(results):
    fig, ax = plt.subplots(2, 2, figsize=(14, 10))
    
    colors = {
        'insertion_sort': '#e6194b',  # red
        'quick_sort': '#3cb44b',      # green
        'hybrid_sort': '#ffe119',     # yellow
        'dual_pivot': '#4363d8',      # blue
        'merge_sort': '#f58231',      # orange
        'my_merge_sort': '#911eb4',   # purple
    }
    
    for label, result in results.items():
        sizes = [r[0] for r in result]
        comparisons = [r[1] for r in result]
        swaps = [r[2] for r in result]
        color = colors.get(label, None)
        
        ax[0, 0].plot(sizes, comparisons, label=label, color=color)
        ax[0, 1].plot(sizes, swaps, label=label, color=color)
        ax[1, 0].plot(sizes, np.array(comparisons) / np.array(sizes), label=label, color=color)
        ax[1, 1].plot(sizes, np.array(swaps) / np.array(sizes), label=label, color=color)
    
    ax[0, 0].set_title("Average Comparisons vs n")
    ax[0, 0].set_xlabel("n")
    ax[0, 0].set_ylabel("Comparisons")
    
    ax[0, 1].set_title("Average Swaps vs n")
    ax[0, 1].set_xlabel("n")
    ax[0, 1].set_ylabel("Swaps")
    
    ax[1, 0].set_title("Comparisons per n vs n")
    ax[1, 0].set_xlabel("n")
    ax[1, 0].set_ylabel("Comparisons / n")
    
    ax[1, 1].set_title("Swaps per n vs n")
    ax[1, 1].set_xlabel("n")
    ax[1, 1].set_ylabel("Swaps / n")
    
    for ax_ in ax.flat:
        ax_.legend()
    
    plt.tight_layout()
    plt.show()


def calculate_C(results=None):
    sizes = []
    comparisons = []
    
    if results is None:
        results = experiment_for_big_sizes_parallel()
    
    for size, comp, _ in results['quick_sort']:
        sizes.append(size)
        comparisons.append(comp)
    
    n_log_n = [size * np.log2(size) for size in sizes]
    C_values = [comp / n_log for comp, n_log in zip(comparisons, n_log_n)]
    C_average = np.mean(C_values)
    
    print(f"Calculated constant C: {C_average:.4f}")
    
    plt.figure(figsize=(8, 6))
    plt.scatter(n_log_n, comparisons, label="Experimental Data", color='blue')
    plt.plot(n_log_n, [C_average * n_log for n_log in n_log_n],
             label=f"Fitted Line (C = {C_average:.4f})", color='red')
    plt.xlabel('n * log2(n)')
    plt.ylabel('Comparisons')
    plt.title('Comparisons vs n * log2(n)')
    plt.legend()
    plt.show()


def analyze_results():
    results = experiment_for_various_sizes()
    
    for sort_name, result in results.items():
        print(f"\n{sort_name} results:")
        for size, comp, swaps in result:
            print(f"n={size}, Comparisons: {comp}, Swaps: {swaps}")
    
    plot_results(results)


def analyze_results_big_sizes():
    results_big_sizes = experiment_for_big_sizes_parallel()
    
    for sort_name, result in results_big_sizes.items():
        print(f"\n{sort_name} results:")
        for size, comp, swaps in result:
            print(f"n={size}, Comparisons: {comp}, Swaps: {swaps}")
    
    plot_results(results_big_sizes)
    

if __name__ == "__main__":
    # Uncomment the function(s) you want to run:
    #analyze_results()
    #results_big = analyze_results_big_sizes()
    calculate_C()
