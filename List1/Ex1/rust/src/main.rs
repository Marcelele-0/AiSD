mod lifo; // Dodanie modułu lifo
mod fifo; // Dodanie modułu fifo


fn main() {
    let mut stack = lifo::Stack::new();
    let mut queue = fifo::Queue::new();
    
    // Dodawanie 50 elementów do stosu i kolejki
    for i in 1..=50 {
        stack.push(i);
        queue.enqueue(i);
    }

    // Wypisanie elementów ze stosu (LIFO)
    println!("Stack:");
    while !stack.is_empty() {
        print!("{}, ", stack.pop().unwrap()); // Usuwanie i wypisywanie elementów
    }
    println!(); // Dodanie nowej linii po wypisaniu elementów ze stosu

    // Wypisanie elementów z kolejki (FIFO)
    println!("Queue:");
    while !queue.is_empty() {
        print!("{}, ", queue.dequeue().unwrap()); // Usuwanie i wypisywanie elementów
    }
    println!(); // Dodanie nowej linii po wypisaniu elementów z kolejki
}
