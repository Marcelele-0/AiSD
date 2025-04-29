import sys
import utils.counters as counters
from quick_sort import quick_sort

def merge(arr, left, right):
    """
    Merge two sorted sub-arrays of arr[] into a single sorted sub-array.
    """
    mid = (left + right) // 2
    left_half = arr[left:mid + 1]
    right_half = arr[mid + 1:right + 1]
    
    i = j = 0
    k = left
    while i < len(left_half) and j < len(right_half):
        counters.comparison_count += 1
        if left_half[i] <= right_half[j]:
            arr[k] = left_half[i]; i += 1
        else:
            arr[k] = right_half[j]; j += 1
        k += 1
    while i < len(left_half):
        arr[k] = left_half[i]; i += 1; k += 1
    while j < len(right_half):
        arr[k] = right_half[j]; j += 1; k += 1

def find_runs(arr, left, right):
    """
    Find ascending runs in arr[left:right+1], return list of (start,end).
    """
    runs = []
    start = left
    for i in range(left+1, right+1):
        counters.comparison_count += 1
        if arr[i] < arr[i-1]:
            runs.append((start, i-1))
            start = i
    runs.append((start, right))
    return runs

def my_merge_sort(arr, left, right, threshold=15):
    """
    Sort the array using natural runs and merge them (kept same signature).
    """
    # znajdź rosnące podciągi w arr[left:right]
    runs = find_runs(arr, left, right)
    # scalaj dopóki jest więcej niż jeden run
    while len(runs) > 1:
        new_runs = []
        for i in range(0, len(runs)-1, 2):
            l, _ = runs[i]
            _, r = runs[i+1]
            merge(arr, l, r)
            new_runs.append((l, r))
        if len(runs) % 2 == 1:
            new_runs.append(runs[-1])
        runs = new_runs

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))

    if n < 40:
        print("Initial array:", " ".join(f"{x:02}" for x in array))

    original_array = array[:]          # zapamiętaj przed sortowaniem
    counters.reset_counters()
    my_merge_sort(array, 0, n-1)       # bez zmiany wywołania

    if n < 40:
        print("Sorted array:", " ".join(f"{x:02}" for x in array))

    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")

    if array == sorted(original_array):
        print("The array is correctly sorted.")
    else:
        print("The array is NOT correctly sorted.")
