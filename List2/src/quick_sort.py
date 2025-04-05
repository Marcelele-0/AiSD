from typing import List
import sys
import random
import utils.counters as counters


def quick_sort(array_to_sort: List[int], low: int, high: int) -> None:
    """
    Implements the quick sort algorithm to sort an array in place with a random pivot.
    Tracks comparisons and swaps during sorting.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        low (int): The starting index of the array segment to be sorted.
        high (int): The ending index of the array segment to be sorted.

    Returns:
        None - It sorts the array in place.
    """
    if low < high:
        pivot_index = partition(array_to_sort, low, high)
        quick_sort(array_to_sort, low, pivot_index - 1)  # Left sub-array
        quick_sort(array_to_sort, pivot_index + 1, high)  # Right sub-array

def partition(array_to_sort: List[int], low: int, high: int) -> int:
    """
    Partitions the array into two sub-arrays based on a pivot chosen randomly.
    Rearranges the elements in the array such that all elements less than the pivot
    are on the left, and all elements greater than the pivot are on the right.

    Parameters:
        array_to_sort (List[int]): The array to be partitioned.
        low (int): The starting index of the partition.
        high (int): The ending index of the partition.

    Returns:
        int: The index of the pivot after partitioning.
    """
   # Median of three: choose pivot as the median of low, high, and middle elements
    mid = (low + high) // 2
    pivot_index = sorted([low, mid, high], key=lambda x: array_to_sort[x])[1]
    
    # Swap pivot with the last element
    array_to_sort[pivot_index], array_to_sort[high] = array_to_sort[high], array_to_sort[pivot_index]
    pivot = array_to_sort[high]
    
    i = low - 1
    for j in range(low, high):
        if counters.compare(array_to_sort[j], pivot):  # If current element <= pivot
            i += 1
            counters.swap(array_to_sort, i, j)  # Swap the elements

    counters.swap(array_to_sort, i + 1, high)  # Place pivot in the correct position
    return i + 1

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    if n < 40:
        print("Initial array:", " ".join(str(x) for x in array))
    
    counters.reset_counters()  # Reset counters before sorting
    quick_sort(array, 0, n - 1)  # Sort the array using quick sort
    
    if n < 40:
        print("Sorted array:", " ".join(str(x) for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")
