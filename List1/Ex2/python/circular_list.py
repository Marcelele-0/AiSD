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
        if self.final_node is None:
            print("List is empty.")
            return
        head = self.final_node.next
        current = head
        for _ in range(self.node_count):
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
         - other_list: CircularList to be merged into this list
        """
    
        if self.final_node is None:
            self.final_node = other_list.final_node
            self.node_count = other_list.node_count
            return
        
        if other_list.final_node is None:
            return
        
        # Save the first nodes of both lists
        self_first_node = self.final_node.next
        other_first_node = other_list.final_node.next

        # Connect the last node of the first list to the first node of the second list
        self.final_node.next = other_first_node
        other_list.final_node.next = self_first_node

        # Update the last node of the first list
        self.final_node = other_list.final_node

        # Update size 
        self.node_count += other_list.node_count

    def list_search(self, value):
        """
        Searches for a node with the specified value in the list.

        params:
         - value: value to search for

        returns:
         - Tuple (found: bool, comparison_count: int) 
        """
        if self.final_node is None:
            return (False, 0)
        
        # initialize variables 
        comparison_count = 0
        head = self.final_node.next
        current = head  

        # search for the value
        for _ in range(self.node_count):
            comparison_count += 1
            if current.value == value:
                return (True, comparison_count)
            current = current.next
        
        return (False, comparison_count)
    
