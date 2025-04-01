from typing import List
from quick_sort import quick_sort
from insertion_sort import insertion_sort
import sys

def hybrid_sort(lenght: int, array_to_sort: List[int], threshold: int = 10) -> None:
    """
    This function implements a hybrid sorting algorithm that combines Quick Sort and Insertion Sort.
    If the sub-array size is below a certain threshold, Insertion Sort is used.
    Otherwise, Quick Sort is applied.

    Parameters:
        lenght (int): The length of the array.
        array_to_sort (List[int]): The array to be sorted.
        threshold (int): The cutoff size below which Insertion Sort is used. Default is 10.

    Returns:
        None - It sorts the array in place.
    """
    if lenght <= threshold:
        insertion_sort(lenght, array_to_sort)
    else:
        quick_sort(array_to_sort, 0, lenght - 1)  


if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    hybrid_sort(n, array)  
    print(" ".join(map(str, array))) 