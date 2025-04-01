import sys

def generate_sorted(n: int) -> None:
    """
    Generates a sorted sequence of `n` keys in increasing order and prints them.

    Parameters:
        n (int): The number of elements to generate.
    """
    array = list(range(1, n + 1))
    print(n)
    print(" ".join(map(str, array)))

if __name__ == "__main__":
    n = int(sys.argv[1]) if len(sys.argv) > 1 else int(input("Enter number of elements: "))
    generate_sorted(n)
