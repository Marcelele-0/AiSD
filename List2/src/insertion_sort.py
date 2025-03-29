from typing import List


def insertion_sort(array_length: int, array_to_sort: List[int]) -> List[int]:
    """
    This function implements the insertion sort algorithm to sort an array in place.
    Intuition: we go over the array, and for each element, we check if it is smaller than the previous elements.

    Parameters:
        array_length (int): The length of the array.
        array_to_sort (List[int]): The array to be sorted.

    Returns:
        List[int]: The sorted array.
    """
    for i in range(1, array_length):
        key = array_to_sort[i]
        j = i - 1

        while j >= 0 and key < array_to_sort[j]:
            array_to_sort[j + 1] = array_to_sort[j]
            j -= 1
        array_to_sort[j + 1] = key

    return array_to_sort
