from typing import List

def quick_sort(array_to_sort: List[int], low: int, high: int) -> None:
    """
    This function implements the quick sort algorithm to sort an array in place.
    
    Intuition: We choose a pivot and partition the array into two sub-arrays:
    one with elements smaller than the pivot and one with elements larger than the pivot.
    We then recursively sort the sub-arrays.

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
        if array_to_sort[j] <= pivot:
            i += 1
            array_to_sort[i], array_to_sort[j] = array_to_sort[j], array_to_sort[i]

    array_to_sort[i + 1], array_to_sort[high] = array_to_sort[high], array_to_sort[i + 1]
    return i + 1

if __name__ == "__main__":
    n = int(input())
    array = list(map(int, input().split()))
    quick_sort(array, 0, n - 1)
    print(" ".join(map(str, array)))
