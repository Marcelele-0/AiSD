import numpy as np
import random
import matplotlib.pyplot as plt
from insertion_sort import insertion_sort
from quick_sort import quick_sort
from hybrid_sort import hybrid_sort
import utils.counters as counters


def run_experiment(sort_function, n: int, k: int, threshold: int = 5) -> tuple:
    """
    Run sorting experiment for a given sorting function and return the average number of comparisons and swaps.
    
    Parameters:
        sort_function: Sorting function to be tested (insertion_sort, quick_sort, hybrid_sort)
        n (int): Number of elements to sort.
        k (int): Number of repetitions to average.
        threshold (int): Threshold value for hybrid_sort (if applicable).
    
    Returns:
        tuple: (average_comparisons, average_swaps)
    """
    comparisons_list = []
    swaps_list = []
    
    for _ in range(k):
        # Generate a random array for each repetition
        array = random.sample(range(1, 50_001), n)
        counters.reset_counters()
        
        if sort_function == hybrid_sort:
            sort_function(array, 0, n - 1, threshold)  # Hybrid sort with threshold
        else:
            sort_function(array, 0, n - 1)  # Insertion sort or quick sort or dual pivot quick sort
        
        # Collect comparisons and swaps
        comparisons_list.append(counters.comparison_count)
        swaps_list.append(counters.swap_count)

    # Calculate average comparisons and swaps
    avg_comparisons = np.mean(comparisons_list)
    avg_swaps = np.mean(swaps_list)
    
    return avg_comparisons, avg_swaps


def experiment_for_big_sizes():
    sizes = [i for i in range(1000, 50001, 1000)]
    k_values = [10]
    results = {'quick_sort': [], 'hybrid_sort': [], 'dual_pivot': []}
    
    for k in k_values:
        for size in sizes:
            print(f"\nRunning experiment for n={size}, k={k}...")
            quick_comparisons, quick_swaps = run_experiment(quick_sort, size, k)
            hybrid_comparisons, hybrid_swaps = run_experiment(hybrid_sort, size, k, threshold=10)
            dual_pivot_comparisons, dual_pivot_swaps = run_experiment(quick_sort, size, k)
            
            # Store the results in the dictionary
            results['quick_sort'].append((size, quick_comparisons, quick_swaps))
            results['hybrid_sort'].append((size, hybrid_comparisons, hybrid_swaps))
            results['dual_pivot'].append((size, dual_pivot_comparisons, dual_pivot_swaps))
    
    return results

def experiment_for_various_sizes():
    sizes = [10, 20, 30, 40, 50]
    k_values = [1000]
    results = {'insertion_sort': [], 'quick_sort': [], 'hybrid_sort': [], 'dual_pivot': []}
    
    for k in k_values:
        for size in sizes:
            print(f"\nRunning experiment for n={size}, k={k}...")
            insertion_comparisons, insertion_swaps = run_experiment(insertion_sort, size, k)
            quick_comparisons, quick_swaps = run_experiment(quick_sort, size, k)
            hybrid_comparisons, hybrid_swaps = run_experiment(hybrid_sort, size, k, threshold=10)
            dual_pivot_comparisons, dual_pivot_swaps = run_experiment(quick_sort, size, k)
            
            # Store the results in the dictionary
            results['insertion_sort'].append((size, insertion_comparisons, insertion_swaps))
            results['quick_sort'].append((size, quick_comparisons, quick_swaps))
            results['hybrid_sort'].append((size, hybrid_comparisons, hybrid_swaps))
            results['dual_pivot'].append((size, dual_pivot_comparisons, dual_pivot_swaps))
    
    return results


def plot_results(results):
    # Create the plots for comparisons and swaps
    fig, ax = plt.subplots(2, 2, figsize=(14, 10))
    
    for label, result in results.items():
        sizes = [r[0] for r in result]
        comparisons = [r[1] for r in result]
        swaps = [r[2] for r in result]
        
        ax[0, 0].plot(sizes, comparisons, label=label)
        ax[0, 1].plot(sizes, swaps, label=label)
        ax[1, 0].plot(sizes, np.array(comparisons) / np.array(sizes), label=label)
        ax[1, 1].plot(sizes, np.array(swaps) / np.array(sizes), label=label)
    
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

def calculate_C():
    """
    Calculate the constant C for comparisons vs n * log2(n) using experimental results.
    
    Parameters:
        results (dict): The dictionary containing the sorting experiment results.
        
    Returns:
        None: Displays the plot and prints the calculated C.
    """
    sizes = []
    comparisons = []
    
    results = experiment_for_big_sizes()

    # Collect data for quick_sort as an example
    for size, comp, _ in results['quick_sort']:
        sizes.append(size)
        comparisons.append(comp)

    # Calculate n * log2(n)
    n_log_n = [size * np.log2(size) for size in sizes]
    
    # Calculate the constant C (comparison / n * log2(n))
    C_values = [comp / n_log for comp, n_log in zip(comparisons, n_log_n)]
    
    # Calculate the average C
    C_average = np.mean(C_values)
    
    # Print the results
    print(f"Calculated constant C: {C_average:.4f}")
    
    # Plot comparisons vs n * log2(n)
    plt.figure(figsize=(8, 6))
    plt.scatter(n_log_n, comparisons, label="Experimental Data", color='blue')
    plt.plot(n_log_n, [C_average * n_log for n_log in n_log_n], label=f"Fitted Line (C = {C_average:.4f})", color='red')
    plt.xlabel('n * log2(n)')
    plt.ylabel('Comparisons')
    plt.title('Comparisons vs n * log2(n)')
    plt.legend()
    plt.show()

def analyze_results():
    # Run the experiment and plot the results
    results = experiment_for_various_sizes()

    # Print the averaged results at the end
    for sort_name, result in results.items():
        print(f"\n{sort_name} results:")
        for size, comparisons, swaps in result:
            print(f"n={size}, Comparisons: {comparisons}, Swaps: {swaps}")

    # Plot the results
    plot_results(results)

def analyze_results_big_sizes():
    # Run the experiment for larger sizes and plot the results
    results_big_sizes = experiment_for_big_sizes()

    # Print the averaged results at the end
    for sort_name, result in results_big_sizes.items():
        print(f"\n{sort_name} results:")
        for size, comparisons, swaps in result:
            print(f"n={size}, Comparisons: {comparisons}, Swaps: {swaps}")

    # Plot the results
    plot_results(results_big_sizes)

if __name__ == "__main__":
    # Uncomment the function you want to run
    #analyze_results_big_sizes()
    #analyze_results()
    calculate_C()
