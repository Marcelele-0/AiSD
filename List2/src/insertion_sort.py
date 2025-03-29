from typing import List

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
    print("Enter the number of elements followed by the elements themselves:")
    try:
        n = int(input().strip())  
        array = list(map(int, input().split()))  
        insertion_sort(n, array)  
        print("Sorted array:", " ".join(map(str, array)))
    except ValueError:
        print("Invalid input. Please enter integers only.")
    except IndexError:
        print("The number of elements does not match the provided array length.")
    except Exception as e:
        print(f"An error occurred: {e}")
