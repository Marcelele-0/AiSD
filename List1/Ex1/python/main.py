from lifo import Stack
from fifo import Queue

def main():
    stack = Stack()
    queue = Queue()
    
    for i in range(50):
        stack.push(i)
        queue.enqueue(i)

    print("Stack:")
    while not stack.is_empty():
        print(stack.pop(), end=" ")
    print()

    print("Queue:")
    while not queue.is_empty():
        print(queue.dequeue(), end=" ")
    print()

if __name__ == "__main__":
    main()