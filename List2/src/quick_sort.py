from typing import List
import sys
import src.utils.counters as counters


def quick_sort(array_to_sort: List[int], low: int, high: int) -> None:
    """
    This function implements the quick sort algorithm to sort an array in place.
    
    Intuition: We choose a pivot and partition the array into two sub-arrays:
    one with elements smaller than the pivot and one with elements larger than the pivot.
    We then recursively sort the sub-arrays. The function also tracks comparisons and swaps.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        low (int): The starting index of the array segment to be sorted.
        high (int): The ending index of the array segment to be sorted.

    Returns:
        None - It sorts the array in place.
    """
    if low < high:
        pivot_index = partition(array_to_sort, low, high)
        quick_sort(array_to_sort, low, pivot_index - 1)
        quick_sort(array_to_sort, pivot_index + 1, high)

def partition(array_to_sort: List[int], low: int, high: int) -> int:
    """
    This function partitions the array into two sub-arrays based on a pivot.
    It rearranges the elements in the array such that all elements less than the pivot
    are on the left side, and all elements greater than the pivot are on the right side.
    Additionally, it tracks comparisons and swaps during partitioning.

    Parameters:
        array_to_sort (List[int]): The array to be partitioned.
        low (int): The starting index of the partition.
        high (int): The ending index of the partition.

    Returns:
        int: The index of the pivot after partitioning.
    """
    pivot = array_to_sort[high]
    i = low - 1

    for j in range(low, high):
        if counters.compare(array_to_sort[j], pivot):  # Use the compare function from counters
            i += 1
            counters.swap(array_to_sort, i, j)         # Use the swap function from counters

    counters.swap(array_to_sort, i + 1, high)           # Use the swap function from counters
    return i + 1

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    if n < 40:
        print("Initial array:", " ".join(str(x) for x in array))
    
    counters.reset_counters()  # Reset counters before sorting
    quick_sort(array, 0, n - 1) 
    
    if n < 40:
        print("Sorted array:", " ".join(str(x) for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")
