use crate::node_value::NodeValue;

#[derive(Default)]
pub enum Node<T> {
    #[default]
    None,
    Some(NodeValue<T>),
}

impl<T> Node<T> {
    pub fn new(item: T, next: Box<Self>) -> Self {
        Self::Some(NodeValue { item, next })
    }
    pub fn take(&mut self) -> Self {
        std::mem::take(self)
    }
}
