
class Node:
    def __init__(self, value):
        """
        this class is a node of a singly linked list

        params
         - value: value to be stored in the node

        variables
         - value: value stored in the node
         - next: pointer to the next node in the list
         - prev: pointer to the previous node in the list
        """
        self.value = value  
        self.next = None
        self.prev = None
