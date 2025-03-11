pub struct Stack<T> {
    items: Vec<T>,
}

impl<T> Stack<T> {
    // Tworzenie nowego stosu
    pub fn new() -> Self {
        Stack { items: Vec::new() }
    }

    // Dodawanie elementu do stosu
    pub fn push(&mut self, item: T) {
        self.items.push(item);
    }

    // Usuwanie elementu ze stosu
    pub fn pop(&mut self) -> Option<T> {
        if self.is_empty() {
            print!("Error: Trying to pop from an empty stack!");
            None
        } else {
            self.items.pop()
        }
    }

    // Sprawdzanie, czy stos jest pusty
    pub fn is_empty(&self) -> bool {
        self.items.is_empty()
    }
}
