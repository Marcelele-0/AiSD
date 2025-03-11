class Queue:
    def __init__(self):
        self.items = []

    def enqueue(self, item):
        self.items.append(item)

    def dequeue(self):
        if self.is_empty():
            print("Error: Trying to dequeue from an empty queue!")
            return None
        return self.items.pop(0)

    def is_empty(self):
        return len(self.items) == 0


