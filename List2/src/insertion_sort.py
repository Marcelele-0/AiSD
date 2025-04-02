from typing import List
import sys
import utils.counters as counters

def insertion_sort(array_length: int, array_to_sort: List[int]) -> None:
    """
    Implements the insertion sort algorithm to sort an array in place.
    Tracks comparisons and swaps during sorting.

    Parameters:
        array_length (int): The length of the array.
        array_to_sort (List[int]): The array to be sorted.
    
    Returns:
        None - the array is sorted in place.
    """
    for i in range(1, array_length):
        key = array_to_sort[i]
        j = i - 1

        while j >= 0 and counters.compare(array_to_sort[j], key):  # Use compare from counters
            counters.swap(array_to_sort, j + 1, j)  # Call swap from counters
            j -= 1
        array_to_sort[j + 1] = key

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    if n < 40:
        print("Initial array:", " ".join(f"{x:02}" for x in array))
    
    counters.reset_counters()  # Reset counters before sorting
    insertion_sort(n, array)  # Sort the array in place
    
    if n < 40:
        print("Sorted array:", " ".join(f"{x:02}" for x in array))
    
    print(f"Total comparisons: {counters.comparison_count}")
    print(f"Total swaps: {counters.swap_count}")
