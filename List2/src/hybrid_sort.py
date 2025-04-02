from typing import List
import sys
from quick_sort import quick_sort
from insertion_sort import insertion_sort
from utils.counters import comparison_count, swap_count, reset_counters

def hybrid_sort(array_to_sort: List[int], threshold: int = 10) -> None:
    """
    Implements a hybrid sorting algorithm that combines Quick Sort and Insertion Sort.
    If the sub-array size is below a certain threshold, Insertion Sort is used.
    Otherwise, Quick Sort is applied.

    Parameters:
        array_to_sort (List[int]): The array to be sorted.
        threshold (int): The cutoff size below which Insertion Sort is used. Default is 10.

    Returns:
        None - It sorts the array in place.
    """
    if len(array_to_sort) <= threshold:
        insertion_sort(len(array_to_sort), array_to_sort)
    else:
        quick_sort(array_to_sort, 0, len(array_to_sort) - 1)  

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    reset_counters()
    if n < 40:
        print("Input array:", " ".join(f"{num:02d}" for num in array))
    
    hybrid_sort(array)
    
    if n < 40:
        print("Sorted array:", " ".join(f"{num:02d}" for num in array))
    
    print(f"Total comparisons: {comparison_count}")
    print(f"Total swaps: {swap_count}")