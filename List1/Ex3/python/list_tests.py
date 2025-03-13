from circular_list import CircularList
import random

def test_insert():
    list = CircularList()
    list.insert(1)
    assert list.final_node.value == 1
    assert list.final_node.next == list.final_node
    list.insert(2)
    assert list.final_node.value == 2
    assert list.final_node.next.value == 1
    list.insert(3)
    assert list.final_node.value == 3
    assert list.final_node.next.value == 1

def test_merge():
    # Create two lists of length 10, containing non-negative two-digit numbers.
    list1 = CircularList()
    list2 = CircularList()
    
    for _ in range(10):
        list1.insert(random.randint(10, 99))
        list2.insert(random.randint(10, 99))
    
    print("List 1:")
    list1.print_list()
    print("List 2:")
    list2.print_list()
    
    # Merge the lists
    list1.merge(list2)
    print("\nMerged List:")
    list1.print_list()

def test_search_cost():
    # Create an array T with 10000 random numbers in the range [0, 100000]
    T = [random.randint(0, 100000) for _ in range(10000)]
    
    # Insert these numbers into the list L
    L = CircularList()
    for num in T:
        L.insert(num)
    
    # Number of search attempts for elements that are definitely in the list
    num_searches_in_list = 1000
    total_cost_found = 0
    for _ in range(num_searches_in_list):
        x = random.choice(T)  # Choose a random number from T (which is definitely in the list)
        _, cost = L.list_search(x)  # Search for it in the list and measure the cost
        total_cost_found += cost
    
    avg_cost_found = total_cost_found / num_searches_in_list

    # Number of search attempts for random numbers that may or may not be in the list
    num_searches_random = 1000
    total_cost_not_found = 0
    for _ in range(num_searches_random):
        x = random.randint(0, 100000)  # Choose a random number, which may or may not be in the list
        _, cost = L.list_search(x)  # Search for it in the list and measure the cost
        total_cost_not_found += cost
    
    avg_cost_not_found = total_cost_not_found / num_searches_random

    print(f"\nAverage comparisons for numbers that are on the list: {avg_cost_found}")
    print(f"Average comparisons for numbers that are not on the list: {avg_cost_not_found}")
