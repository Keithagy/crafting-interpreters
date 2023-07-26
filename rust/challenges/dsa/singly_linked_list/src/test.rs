#[cfg(test)]
mod tests {
    use crate::list::SinglyLinkedList;

    #[test]
    fn it_works() {
        let mut linked_list: SinglyLinkedList<usize> = SinglyLinkedList::new();
        for i in 1..=10 {
            linked_list.push(i);
        }

        for i in (1..=10).rev() {
            let cur = linked_list.pop();
            assert_eq!(Some(i), cur);
        }

        assert_eq!(None, linked_list.pop());
    }

    #[test]
    fn test_series_of_pops_and_inserts() {
        let mut list: SinglyLinkedList<usize> = SinglyLinkedList::new();
        assert_eq!(list.pop(), None);

        list.push(3);
        list.push(42);
        assert_eq!(list.pop(), Some(42));
        assert_eq!(list.len(), 1);

        list.push(93);
        assert_eq!(list.len(), 2);
        assert_eq!(list.pop(), Some(93));
        assert_eq!(list.pop(), Some(3));
        assert_eq!(list.pop(), None);
        assert_eq!(list.len(), 0);

        list.push(20);
        assert_eq!(list.pop(), Some(20));
        assert_eq!(list.pop(), None);
    }
}
