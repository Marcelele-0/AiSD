from circular_list import CircularList

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

def main():
    test_insert()
    print("All tests pass")

if __name__ == "__main__":
    main()