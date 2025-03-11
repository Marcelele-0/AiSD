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
    # Tworzymy dwie listy o długości 10, zawierające dwucyfrowe liczby nieujemne.
    list1 = CircularList()
    list2 = CircularList()
    
    for _ in range(10):
        list1.insert(random.randint(10, 99))
        list2.insert(random.randint(10, 99))
    
    print("List 1:")
    list1.print_list()
    print("List 2:")
    list2.print_list()
    
    # Łączymy listy (merge)
    list1.merge(list2)
    print("\nMerged List:")
    list1.print_list()

def test_search_cost():
    list = CircularList()
    values = [random.randint(0, 100_000) for _ in range(1000)]

    for value in values:
        list.insert(value)

    # TODO: test search cost
    

def main():
    test_insert()
    print("Insert tests passed.")

    test_merge()
    print("Merge tests passed.")


if __name__ == "__main__":
    main()