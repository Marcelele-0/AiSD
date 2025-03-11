use std::collections::VecDeque;

pub struct Queue<T> {
    items: VecDeque<T>,
}

impl<T> Queue<T> {
    // Tworzenie nowej kolejki
    pub fn new() -> Self {
        Queue {
            items: VecDeque::new(),
        }
    }

    // Dodawanie elementu do kolejki
    pub fn enqueue(&mut self, item: T) {
        self.items.push_back(item);
    }

    // Usuwanie elementu z kolejki
    pub fn dequeue(&mut self) -> Option<T> {
        if self.is_empty() {
            println!("Error: Trying to dequeue from an empty queue!");
            None
        } else {
            self.items.pop_front()
        }
    }

    // Sprawdzanie, czy kolejka jest pusta
    pub fn is_empty(&self) -> bool {
        self.items.is_empty()
    }
}
