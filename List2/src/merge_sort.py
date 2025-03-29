from typing import List

def merge_sort(length: int, arr: List[int]) -> List[int]:
    """
    This function implements the merge sort algorithm to sort an array in place.

    Parameters:
        length (int): The length of the array.
        arr (List[int]): The array to be sorted.

    Returns:
        array (List[int]): The sorted array.
    """
    if length <= 1:
        return arr
    
    mid = length // 2
    left = merge_sort(mid, arr[:mid])
    right = merge_sort(length - mid, arr[mid:])
    
    return merge(left, right)
    


def merge(left: List[int], right: List[int]) -> List[int]:
    """
    This function merges two sorted arrays into a single sorted array.

    Parameters:
        left (List[int]): The left sorted array.
        right (List[int]): The right sorted array.

    Returns:
        merged (List[int]): The merged sorted array.
    """
    if len(left) == 0:
        return right
    if len(right) == 0:
        return left
    
    if left[0] < right[0]:
        return [left[0]] + merge(left[1:], right)
    else:
        return [right[0]] + merge(left, right[1:])

