import utils.counters as counters
import sys


def insertion_sort(arr, low, high):
    for i in range(low + 1, high + 1):
        key = arr[i]
        j = i - 1
        while j >= low and not counters.compare(arr[j], key):
            counters.swap(arr, j + 1, j)
            j -= 1
        arr[j + 1] = key


def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1
    for j in range(low, high):
        if counters.compare(arr[j], pivot):
            i += 1
            counters.swap(arr, i, j)
    counters.swap(arr, i + 1, high)
    return i + 1


def partition_around_pivot(arr, low, high, pivot_value):
    for i in range(low, high + 1):
        if arr[i] == pivot_value:
            counters.swap(arr, i, high)
            break
    return partition(arr, low, high)


def my_select(arr, low, high, k, group_size=5):
    n = high - low + 1
    if n <= group_size:
        insertion_sort(arr, low, high)
        return arr[low + k - 1]

    medians = []
    for i in range(low, high + 1, group_size):
        group_end = min(i + group_size - 1, high)
        insertion_sort(arr, i, group_end)
        median = arr[i + (group_end - i) // 2]
        medians.append(median)

    median_of_medians = my_select(medians, 0, len(medians) - 1, len(medians) // 2 + 1)
    pivot_index = partition_around_pivot(arr, low, high, median_of_medians)

    left_size = pivot_index - low + 1
    if k == left_size:
        return arr[pivot_index]
    elif k < left_size:
        return my_select(arr, low, pivot_index - 1, k)
    else:
        return my_select(arr, pivot_index + 1, high, k - left_size)


if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    k = int(input_data[1].strip())
    array = list(map(int, input_data[2].split()))

    # Jeżeli n < 30, wyświetlamy początkową tablicę
    if n < 30:
        print("Initial array:", array)

    counters.reset_counters()
    arr_copy = array.copy()
    result_select = my_select(arr_copy, 0, n - 1, k)

    print("\n--- DETERMINISTIC SELECT ---")
    
    if n < 30:
        print("Array after:", arr_copy)
    
    print(f"{k}-th statistic: {result_select}")
    print("Sorted array:", sorted(array))
    print(f"Comparisons: {counters.comparison_count}, Swaps: {counters.swap_count}")
