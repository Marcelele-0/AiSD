from typing import List
import sys

def insertion_sort(array_length: int, array_to_sort: List[int]) -> None:
    """
    Implements the insertion sort algorithm to sort an array in place.

    Parameters:
        array_length (int): The length of the array.
        array_to_sort (List[int]): The array to be sorted.
    
    Returns:
        None - the array is sorted in place.
    """
    for i in range(1, array_length):
        key = array_to_sort[i]
        j = i - 1

        while j >= 0 and key < array_to_sort[j]:
            array_to_sort[j + 1] = array_to_sort[j]
            j -= 1
        array_to_sort[j + 1] = key

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))
    
    insertion_sort(n, array)  
    print(" ".join(map(str, array))) 