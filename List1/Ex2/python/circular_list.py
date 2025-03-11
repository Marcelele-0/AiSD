from node import Node

class CircularList:
    """
    this class is a circular linked list

    implements:
     - insert - adds a new node to the list with the specified value
     - merge - merges two lists
     - print - prints the list
    """

    def __init__(self):
        """ 
        Initializes a new instance of the CircularList class.
        """
        self.final_node = None
        self.node_count = 0 

    def print_list(self):
        """
        prints the list
        """
        if self.tail is None:
            print("List is empty.")
            return
        head = self.tail.next
        current = head
        for _ in range(self.size):
            print(current.value, end=" -> ")
            current = current.next
        print("(back to head)")
    
    def insert(self, value):
        """
        Adds a new node to the list with the specified value.

         - if the list is empty, the new node will be the tail of the list
         - otherwise, the new node will be inserted after the tail and become the new tail

        params:
         - value: value to be stored in the node
        """
        new_node = Node(value)
        if self.final_node is None:
            self.final_node = new_node
            self.final_node.next = self.final_node  # point to itself
        else:
            new_node.next = self.final_node.next    # point to the first node
            self.final_node.next = new_node         # point to the new node
            self.final_node = new_node              # update the tail
        self.node_count += 1

    def merge(self, other_list) -> None: 
        """
        Merges another circular list into this list.

        params:
         - self: CircularList to merge into
         - other_list: CircularList to be merged into this list
        """
    pass

