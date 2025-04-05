from typing import List
import sys
from quick_sort import quick_sort
from insertion_sort import insertion_sort
import utils.counters as counters

def hybrid_sort(array_to_sort: List[int], threshold: int = 15) -> None:
    """
    Implements a hybrid sorting algorithm that combines Quick Sort and Insertion Sort.
    If the sub-array size is below a certain threshold, Insertion Sort is used.
    Otherwise, Quick Sort is applied recursively.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        threshold (int): The cutoff size below which Insertion Sort is used. Default is 7.

    Returns:
        None - It sorts the array in place.
    """
    # Start the sorting process from the full array
    _hybrid_sort_recursive(array_to_sort, 0, len(array_to_sort) - 1, threshold)

def _hybrid_sort_recursive(array_to_sort: List[int], low: int, high: int, threshold: int) -> None:
    """
    A helper recursive function that applies either Insertion Sort or Quick Sort
    depending on the size of the sub-array.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        low (int): The starting index of the array segment to be sorted.
        high (int): The ending index of the array segment to be sorted.
        threshold (int): The cutoff size below which Insertion Sort is used.

    Returns:
        None - It sorts the array in place.
    """
    if high - low + 1 <= threshold:
        # If the subarray size is small, use insertion sort
        insertion_sort(high - low + 1, array_to_sort[low:high+1])
    else:
        # Otherwise, apply quick sort
        if low < high:
            pivot_index = partition(array_to_sort, low, high)
            _hybrid_sort_recursive(array_to_sort, low, pivot_index - 1, threshold)
            _hybrid_sort_recursive(array_to_sort, pivot_index + 1, high, threshold)

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
    
    counters.reset_counters()  # Reset counters before sorting
    if n < 40:
        print("Initial array:", " ".join(f"{x:02}" for x in array))
    
    hybrid_sort(array)  # Call the hybrid_sort function
    
    if n < 40:
        print("Sorted array:", " ".join(f"{x:02}" for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")
