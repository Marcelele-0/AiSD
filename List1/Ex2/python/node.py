
class Node:
    def __init__(self, value):
        """
        this class is a node of a singly linked list

        params
         - value: value to be stored in the node

        variables
         - value: value stored in the node
         - next: pointer to the next node in the list
        """
        self.value = value  # Wartość przechowywana przez element
        self.next = None    # Wskaźnik do następnego elementu w cyklu
