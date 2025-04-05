import sys
import random

def generate_random(n: int, seed: int = None) -> None:
    """
    Generates a sequence of `n` random keys and prints them.

    Parameters:
        n (int): The number of elements to generate.
        seed (int, optional): The seed for reproducibility.
    """
    if seed is not None:
        random.seed(seed)  # Ensure reproducibility

    array = [random.randint(1, 100) for _ in range(n)]  # Large range for variety
    print(n)
    print(" ".join(map(str, array)))

if __name__ == "__main__":
    n = int(sys.argv[1]) if len(sys.argv) > 1 else int(input("Enter number of elements: "))
    generate_random(n)
