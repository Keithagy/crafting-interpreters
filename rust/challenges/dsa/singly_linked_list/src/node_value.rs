use crate::node::Node;

pub struct NodeValue<T> {
    pub item: T,
    pub next: Box<Node<T>>,
}

impl<T> NodeValue<T> {
    pub fn new(value: T, next: Box<Node<T>>) -> Self {
        Self { item: value, next }
    }
}
