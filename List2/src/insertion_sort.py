from typing import List
import sys
import utils.counters as counters


def insertion_sort(array_to_sort: List[int], low: int, high: int) -> None:
    """
    Implements the insertion sort algorithm to sort a subarray in place.
    Tracks comparisons and swaps during sorting.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        low (int): The starting index of the subarray to be sorted.
        high (int): The ending index of the subarray to be sorted.

    Returns:
        None - the subarray is sorted in place.
    """
    for i in range(low + 1, high + 1):
        key = array_to_sort[i]
        j = i - 1

        while j >= low and not counters.compare(array_to_sort[j], key):   # Use the compare function from counters
            counters.swap(array_to_sort, j + 1, j)                      # Use the swap function from counters
            j -= 1
        array_to_sort[j + 1] = key

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    if n < 40:
        print("Initial array:", " ".join(f"{x:02}" for x in array))
    
    counters.reset_counters()  # Reset counters before sorting
    insertion_sort(array, 0, n - 1)  # Sort the array in place, with correct low and high indices
    
    if n < 40:
        print("Sorted array:", " ".join(f"{x:02}" for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")
    original_array = array[:]  # Make a copy of the original array
    # Check if the result is a sorted version of the input
    if array == sorted(original_array):
        print("The array is correctly sorted.")
    else:
        print("The array is NOT correctly sorted.")