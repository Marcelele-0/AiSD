from typing import List
import sys
import utils.counters as counters
from my_select import my_select


def select_dual_pivot_quick_sort(array_to_sort: List[int], low: int, high: int) -> None:
    """
    Dual-pivot QuickSort using deterministic SELECT for pivot selection.
    """
    if low < high:
        p1, p2 = dual_partition(array_to_sort, low, high)
        select_dual_pivot_quick_sort(array_to_sort, low, p1 - 1)
        select_dual_pivot_quick_sort(array_to_sort, p1 + 1, p2 - 1)
        select_dual_pivot_quick_sort(array_to_sort, p2 + 1, high)


def dual_partition(array: List[int], low: int, high: int) -> tuple:
    """
    Partition array into three parts using two pivots selected by SELECT.
    """
    size = high - low + 1
    # positions for pivots: 1/3 and 2/3
    k1 = size // 3
    k2 = 2 * size // 3

    # select pivot values
    pivot_low = my_select(array, low, high, k1)
    pivot_high = my_select(array, low, high, k2)

    # find indices of these pivot values within [low, high]
    low_indices = [i for i in range(low, high+1) if array[i] == pivot_low]
    high_indices = [i for i in range(low, high+1) if array[i] == pivot_high and i not in low_indices]
    # fallback if not found
    pl = low_indices[0] if low_indices else low
    ph = high_indices[0] if high_indices else high

    counters.swap(array, low, pl)
    counters.swap(array, high, ph)

    # refresh pivot values
    pivot_low = array[low]
    pivot_high = array[high]
    if counters.compare(pivot_high, pivot_low):
        counters.swap(array, low, high)
        pivot_low, pivot_high = pivot_high, pivot_low

    left = low + 1
    right = high - 1
    i = left
    while i <= right:
        if counters.compare(array[i], pivot_low):
            counters.swap(array, i, left)
            left += 1
            i += 1
        elif counters.compare(pivot_high, array[i]):
            counters.swap(array, i, right)
            right -= 1
        else:
            i += 1

    counters.swap(array, low, left - 1)
    counters.swap(array, high, right + 1)
    return left - 1, right + 1


if __name__ == "__main__":
    data = sys.stdin.read().splitlines()
    n = int(data[0])
    arr = list(map(int, data[1].split()))
    if n < 40:
        print("Initial array:", arr)
    counters.reset_counters()
    select_dual_pivot_quick_sort(arr, 0, n-1)
    if n < 40:
        print("Sorted array:", arr)
    print(f"Comparisons: {counters.comparison_count}, Swaps: {counters.swap_count}")
    print("Correctly sorted?", arr == sorted(arr))
