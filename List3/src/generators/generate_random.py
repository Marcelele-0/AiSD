import sys
import random

def generate_random(n: int, k: int, seed: int = None) -> None:
    """
    Generates a sequence of `n` random keys and prints them.

    Parameters:
        n (int): The number of elements to generate.
        k (int): The k-th statistical position.
        seed (int, optional): The seed for reproducibility.
    """
    if seed is not None:
        random.seed(seed)  # Ensure reproducibility

    array = [random.randint(1, 100) for _ in range(n)]  # Generate random numbers within a range

    # Print the output in the expected format for select.py
    print(n)  # Number of elements
    print(k)  # k-th statistic
    print(" ".join(map(str, array)))  # The generated array

if __name__ == "__main__":
    # Fetch command-line arguments or ask the user for inputs
    n = int(sys.argv[1]) if len(sys.argv) > 1 else int(input("Enter number of elements: "))
    k = int(sys.argv[2]) if len(sys.argv) > 2 else int(input("Enter the k-th statistic: "))
    seed = int(sys.argv[3]) if len(sys.argv) > 3 else None  # Optional seed for reproducibility
    
    generate_random(n, k, seed)
