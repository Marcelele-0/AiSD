import random
import utils.counters as counters
import sys


def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1
    for j in range(low, high):
        if counters.compare(arr[j], pivot):
            i += 1
            counters.swap(arr, i, j)
    counters.swap(arr, i + 1, high)
    return i + 1


def randomized_partition(arr, low, high):
    pivot_index = random.randint(low, high)
    counters.swap(arr, pivot_index, high)
    return partition(arr, low, high)


def randomized_select(arr, low, high, k):
    if low == high:
        return arr[low]
    q = randomized_partition(arr, low, high)
    left_size = q - low + 1
    if k == left_size:
        return arr[q]
    elif k < left_size:
        return randomized_select(arr, low, q - 1, k)
    else:
        return randomized_select(arr, q + 1, high, k - left_size)


if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    k = int(input_data[1].strip())
    array = list(map(int, input_data[2].split()))

    if n < 30:
        print("Initial array:", array)

    # RANDOMIZED SELECT
    counters.reset_counters()
    arr_copy = array.copy()
    result_rand = randomized_select(arr_copy, 0, n - 1, k)
    print("\n--- RANDOMIZED SELECT ---")
    if n < 30:
        print("Array after:", arr_copy)
    print(f"{k}-th statistic: {result_rand}")
    print("Sorted array:", sorted(array))
    print(f"Comparisons: {counters.comparison_count}, Swaps: {counters.swap_count}")