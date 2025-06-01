import sys

def generate_reversed(n: int, k: int) -> None:
    """
    Generates a sorted sequence of `n` keys in decreasing order and prints them,
    along with the k-th statistical position.

    Parameters:
        n (int): The number of elements to generate.
        k (int): The k-th statistical position.
    """
    array = list(range(n, 0, -1))  # Sequence in decreasing order
    print(n)  # Number of elements
    print(k)  # k-th statistic
    print(" ".join(map(str, array)))  # The generated array

if __name__ == "__main__":
    n = int(sys.argv[1]) if len(sys.argv) > 1 else int(input("Enter number of elements: "))
    k = int(sys.argv[2]) if len(sys.argv) > 2 else int(input("Enter the k-th statistic: "))
    generate_reversed(n, k)
