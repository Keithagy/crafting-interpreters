use crate::node::Node;

pub struct SinglyLinkedList<T> {
    head: Box<Node<T>>,
    size: usize,
}

impl <T> Default for SinglyLinkedList<T> {
    fn default() -> Self {
        Self {
            head: Box::<Node<T>>::default(),
            size: 0
        }
    }
}

impl<T> SinglyLinkedList<T> {
    pub fn new() -> Self {
        Default::default()
    }
    pub fn len(&self) -> usize {
        self.size
    }
    pub fn is_empty(&self) -> bool {
        self.size == 0
    }
    pub fn push(&mut self, value: T) {
        let cur_head = self.head.take();
        let new_head = Box::new(Node::new(value, Box::new(cur_head)));
        self.head = new_head;
        self.size += 1;
    }
    pub fn pop(&mut self) -> Option<T> {
        if let Node::Some(value) = self.head.take() {
            self.head = value.next;
            self.size -= 1;
            Some(value.item)
        } else {
            None
        }
    }
}
