from typing import List
import sys
import random
import utils.counters as counters


def dual_pivot_quick_sort(array_to_sort: List[int], low: int, high: int) -> None:
    """
    Implements the dual-pivot quick sort algorithm to sort an array in place.
    This version uses the median of three to choose the two pivots.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        low (int): The starting index of the array segment to be sorted.
        high (int): The ending index of the array segment to be sorted.

    Returns:
        None - It sorts the array in place.
    """
    if low < high:
        pivot_1, pivot_2 = dual_partition(array_to_sort, low, high)
        dual_pivot_quick_sort(array_to_sort, low, pivot_1 - 1)
        dual_pivot_quick_sort(array_to_sort, pivot_1 + 1, pivot_2 - 1)
        dual_pivot_quick_sort(array_to_sort, pivot_2 + 1, high)

def dual_partition(array_to_sort: List[int], low: int, high: int) -> tuple:
    """
    Partitions the array into three sub-arrays based on two pivots 

    parameters:
        array_to_sort (List[int]): The array to be partitioned.
        low (int): The starting index of the partition.
        high (int): The ending index of the partition.
    """
    # Choose two pivots 
    lower_pivot_index, high_pivot_index = choose_two_pivots_without_sorting(array_to_sort, low, high)
 
    # Swap the pivots with the first and last elements
    counters.swap(array_to_sort, low, lower_pivot_index)
    counters.swap(array_to_sort, high, high_pivot_index)

    # Ensure the lower pivot is less than the higher pivot
    pivot_low = array_to_sort[low]
    pivot_high = array_to_sort[high]

    # Ensure the pivots are in the correct order
    if counters.compare(pivot_high, pivot_low):
        counters.swap(array_to_sort, low, high)  
        pivot_low, pivot_high = pivot_high, pivot_low


    # Initialize pointers for partitioning
    left_pointer = low + 1
    right_pointer = high - 1
    current_pointer = left_pointer

    # Loop through the array and partition it into three sections
    while current_pointer <= right_pointer:
        if counters.compare(array_to_sort[current_pointer], pivot_low):
            counters.swap(array_to_sort, current_pointer, left_pointer)
            left_pointer += 1
            current_pointer += 1
        elif counters.compare(pivot_high, array_to_sort[current_pointer]):
            counters.swap(array_to_sort, current_pointer, right_pointer)
            right_pointer -= 1
        else:
            current_pointer += 1
    
    # Swap the pivots to their correct positions
    counters.swap(array_to_sort, low, left_pointer - 1)
    counters.swap(array_to_sort, high, right_pointer + 1)

    # Return the indices of the pivots
    return left_pointer - 1, right_pointer + 1  

def choose_two_pivots_without_sorting(array_to_sort: List[int], low: int, high: int):
    # Select 5 elements (low, mid, high, and their neighbors)
    mid = (low + high) // 2
    candidates = [array_to_sort[low], array_to_sort[mid], array_to_sort[high], 
                  array_to_sort[low+1] if low+1 <= high else None, 
                  array_to_sort[high-1] if high-1 >= low else None]
    
    # Remove None if low+1 or high-1 is out of bounds
    candidates = [x for x in candidates if x is not None]
    
    # Now select 2 elements, e.g., the median of the first 3
    candidates_sorted = sorted(candidates)
    pivot_low = candidates_sorted[1]  # Second smallest
    pivot_high = candidates_sorted[3]  # Fourth smallest

    # Find the original indices
    pivot_low_index = array_to_sort.index(pivot_low)
    pivot_high_index = array_to_sort.index(pivot_high)
    
    return pivot_low_index, pivot_high_index


if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    if n < 40:
        print("Initial array:", " ".join(str(x) for x in array))
    
    counters.reset_counters()  # Reset counters before sorting
    dual_pivot_quick_sort(array, 0, n - 1)  # Sort the array using quick sort
    
    if n < 40:
        print("Sorted array:", " ".join(str(x) for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")

    original_array = array[:]  # Make a copy of the original array
    # Check if the result is a sorted version of the input
    if array == sorted(original_array):
        print("The array is correctly sorted.")
    else:
        print("The array is NOT correctly sorted.")