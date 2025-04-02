comparison_count = 0
swap_count = 0

def compare(a: int, b: int) -> bool:
    """Compares two elements and increments the comparison counter."""
    global comparison_count
    comparison_count += 1
    return a <= b

def swap(arr: list, i: int, j: int) -> None:
    """Swaps two elements in the array and increments the swap counter."""
    global swap_count
    arr[i], arr[j] = arr[j], arr[i]
    swap_count += 1

def reset_counters() -> None:
    """Resets the comparison and swap counters."""
    global comparison_count, swap_count
    comparison_count = 0
    swap_count = 0
