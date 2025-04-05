from experiments.threshold import run_threshold_analysis
from experiments.sorting_benchmarking import run_sorting_benchmarking

if __name__ == "__main__":
    run_threshold_analysis()

    sizes = [10, 20, 30, 40, 50]  # Sizes of arrays to experiment with
    k_values = [10, 100, 1000]  # Number of repetitions for each experiment
    
    for k in k_values:
        print(f"\nRunning experiments for k = {k}...")
        run_sorting_benchmarking(k=k, sizes=sizes) 