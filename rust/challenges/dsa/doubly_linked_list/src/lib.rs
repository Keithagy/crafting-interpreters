use std::cell::RefCell;
use std::rc::Rc;

#[derive(Default)]
pub struct DoublyLinkedList<T> {
    pub head: Link<T>,
    pub tail: Link<T>,
    pub size: usize,
}

impl<T> Drop for DoublyLinkedList<T> {
    fn drop(&mut self) {
        while let Some(node) = self.head.take() {
            let _ = node.borrow_mut().prev.take();
            self.head = node.borrow_mut().next.take();
        }
        self.tail.take();
    }
}

impl<T> IntoIterator for DoublyLinkedList<T> {
    type Item = <ListIterator<T> as Iterator>::Item;

    type IntoIter = ListIterator<T>;

    fn into_iter(self) -> Self::IntoIter {
        Self::IntoIter::new(self)
    }
}

pub struct ListIterator<T> {
    list: DoublyLinkedList<T>,
}

impl<T> ListIterator<T> {
    fn new(list: DoublyLinkedList<T>) -> Self {
        Self { list }
    }
}

impl<T> Iterator for ListIterator<T> {
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        self.list.pop_head()
    }
}

impl<T> DoubleEndedIterator for ListIterator<T> {
    fn next_back(&mut self) -> Option<Self::Item> {
        self.list.pop_tail()
    }
}


impl<T> DoublyLinkedList<T> {
    pub fn new() -> Self {
        Self {
            head: None,
            tail: None,
            size: 0,
        }
    }
    pub fn len(&self) -> usize {
        self.size
    }
    pub fn is_empty(&self) -> bool {
        self.size == 0
    }
    pub fn push_head(&mut self, item: T) {
        let node = Rc::new(RefCell::new(ListNode::new(item)));
        if let Some(prev_head) = self.head.take() {
            prev_head.borrow_mut().prev = Some(node.clone());
            node.borrow_mut().next = Some(prev_head.clone());
            self.head = Some(node);
            self.size += 1
        } else {
            self.tail = Some(node.clone());
            self.head = Some(node);
            self.size = 1;
        }
    }

    pub fn push_tail(&mut self, item: T) {
        let node = Rc::new(RefCell::new(ListNode::new(item)));
        if let Some(prev_tail) = self.tail.take() {
            prev_tail.borrow_mut().next = Some(node.clone());
            node.borrow_mut().prev = Some(prev_tail.clone());
            self.tail = Some(node);
            self.size += 1;
        } else {
            self.head = Some(node.clone());
            self.tail = Some(node);
            self.size = 1;
        }
    }
    pub fn pop_head(&mut self) -> Option<T> {
        self.head.take().map(|prev_head| {
            self.size -= 1;
            match prev_head.borrow_mut().next.take() {
                None => { self.tail.take(); } // this serves to switch self.nail to None (since self. tail would in this case be pointing to prev_head)
                Some(node) => {
                    node.borrow_mut().prev = None;
                    self.head = Some(node);
                }
            }
            Rc::try_unwrap(prev_head).ok().unwrap().into_inner().item
        })
    }
    pub fn pop_tail(&mut self) -> Option<T> {
        self.tail.take().map(|prev_tail| {
            self.size -= 1;
            match prev_tail.borrow_mut().prev.take() {
                None => { self.head.take(); } // this serves to switch self.nail to None (since self. tail would in this case be pointing to prev_head)
                Some(node) => {
                    node.borrow_mut().next = None;
                    self.tail = Some(node);
                }
            }
            Rc::try_unwrap(prev_tail).ok().unwrap().into_inner().item
        })
    }
}

pub struct ListNode<T> {
    item: T,
    next: Link<T>,
    prev: Link<T>,
}

impl<T> ListNode<T> {
    fn new(item: T) -> Self {
        Self {
            item,
            next: None,
            prev: None,
        }
    }
}

type Link<T> = Option<Rc<RefCell<ListNode<T>>>>;

#[cfg(test)]
mod tests {
    use super::DoublyLinkedList;

    #[test]
    fn it_works() {
        let mut list = DoublyLinkedList::new();

        list.push_tail(3);
        list.push_tail(4);
        assert_eq!(list.pop_head(), Some(3));
        assert_eq!(list.len(), 1);

        list.push_head(5);
        assert_eq!(list.pop_tail(), Some(4));
        assert_eq!(list.pop_tail(), Some(5));
        assert_eq!(list.pop_tail(), None);
        assert_eq!(list.pop_head(), None);
        assert_eq!(list.len(), 0);
    }

    #[test]
    fn can_push_back() {
        let mut list = DoublyLinkedList::new();
        assert_eq!(list.pop_tail(), None);

        list.push_tail(3);
        list.push_tail(4);
        list.push_tail(5);
        assert_eq!(list.pop_tail(), Some(5));

        list.push_tail(6);
        list.push_tail(7);
        assert_eq!(list.pop_tail(), Some(7));
        assert_eq!(list.pop_tail(), Some(6));
        assert_eq!(list.pop_tail(), Some(4));
        assert_eq!(list.pop_tail(), Some(3));

        list.push_tail(2);
        assert_eq!(list.pop_tail(), Some(2));
        assert_eq!(list.pop_tail(), None);
        assert_eq!(list.len(), 0);
    }

    #[test]
    fn it_drops_correctly() {
        let mut list = DoublyLinkedList::new();
        for i in 0..3 {
            list.push_tail(i);
        }

        drop(list);
    }

    #[test]
    fn can_iterate_forward() {
        let mut list = DoublyLinkedList::new();
        for i in 0..10 {
            list.push_head(i);
        }

        assert!(Iterator::eq(list.into_iter(), (0..10).rev()));
    }

    #[test]
    fn can_iterate_back() {
        let mut list = DoublyLinkedList::new();
        for i in 0..10 {
            list.push_head(i);
        }

        assert!(Iterator::eq(list.into_iter().rev(), 0..10));
    }
}
