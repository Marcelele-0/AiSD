from list_tests import test_insert, test_merge, test_search_cost
    

def main():
    test_insert()
    print("Insert tests passed.")

    test_merge()
    print("Merge tests passed.")

    test_search_cost()
    print("Search cost tests passed.")


if __name__ == "__main__":
    main()