class Stack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if self.is_empty():
            print("Error: Trying to pop from an empty stack!")
            return None
        return self.items.pop()

    def peek(self):
        if self.is_empty():
            print("Error: Stack is empty!")
            return None
        return self.items[-1]

    def is_empty(self):
        return len(self.items) == 0

