from insertion_sort import insertion_sort
from quick_sort import partition
from typing import List
import sys
import utils.counters as counters


def hybrid_sort(array_to_sort: List[int], low: int, high: int, threshold:int) -> None:
    """
    Implements the hybrid sort algorithm, which combines quick sort and insertion sort.
    It uses quick sort for larger arrays and switches to insertion sort for smaller sub-arrays.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        low (int): The starting index of the array segment to be sorted.
        high (int): The ending index of the array segment to be sorted.
        threshold (int): The threshold size for switching to insertion sort.

    Returns:
        None - It sorts the array in place.
    """
    
    if high - low <= threshold:

        insertion_sort(array_to_sort, low, high)  # Use insertion sort for small sub-arrays
    else:
        pivot_index = partition(array_to_sort, low, high)
        hybrid_sort(array_to_sort, low, pivot_index - 1, threshold)  # Left sub-array
        hybrid_sort(array_to_sort, pivot_index + 1, high, threshold) # Right sub-array
    

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    if n < 40:
        print("Initial array:", " ".join(f"{x:02}" for x in array))
    
    counters.reset_counters()  # Reset counters before sorting
    hybrid_sort(array, 0, n - 1)  # Sort the array in place, with correct low and high indices
    
    if n < 40:
        print("Sorted array:", " ".join(f"{x:02}" for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")
