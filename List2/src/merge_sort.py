import sys
import utils.counters as counters

def merge(arr, left, right):
    """
    Merge two sorted sub-arrays of arr[] into a single sorted sub-array.
    
    Parameters:
        arr (list): The list to be sorted.
        left (int): Starting index of the left sub-array.
        right (int): Ending index of the right sub-array.
    
    Returns:
        None: The list is sorted in place.
    """
    mid = (left + right) // 2
    left_half = arr[left:mid + 1]
    right_half = arr[mid + 1:right + 1]
    
    i = j = 0
    k = left

    # Merge the two halves
    while i < len(left_half) and j < len(right_half):
        counters.comparison_count += 1  # Count comparison
        if left_half[i] < right_half[j]:
            arr[k] = left_half[i]
            i += 1
        else:
            arr[k] = right_half[j]
            j += 1
        k += 1

    # If there are remaining elements in left_half
    while i < len(left_half):
        arr[k] = left_half[i]
        i += 1
        k += 1

    # If there are remaining elements in right_half
    while j < len(right_half):
        arr[k] = right_half[j]
        j += 1
        k += 1

def merge_sort(arr, left, right):
    """
    Sort the array using Merge Sort algorithm (recursive).
    
    Parameters:
        arr (list): The list to be sorted.
        left (int): Starting index of the array.
        right (int): Ending index of the array.
    
    Returns:
        None: The list is sorted in place.
    """
    if left < right:
        mid = (left + right) // 2
        
        # Recursively split the array and sort
        merge_sort(arr, left, mid)
        merge_sort(arr, mid + 1, right)
        
        # Merge the sorted sub-arrays
        merge(arr, left, right)

if __name__ == "__main__":
    input_data = sys.stdin.read().splitlines()
    n = int(input_data[0].strip())
    array = list(map(int, input_data[1].split()))

    if n < 40:
        print("Initial array:", " ".join(f"{x:02}" for x in array))

    counters.reset_counters()  # Reset counters before sorting
    merge_sort(array, 0, n - 1)  # Sort the array in place, with correct low and high indices

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