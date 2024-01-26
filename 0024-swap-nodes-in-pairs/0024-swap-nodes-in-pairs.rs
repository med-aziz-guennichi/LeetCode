// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if let Some(mut first) = head {
            if let Some(mut second) = first.next.take() {
                first.next = Self::swap_pairs(second.next.take());
                second.next = Some(first);
                Some(second)
            } else {
                Some(first)
            }
        } else {
            None
        }
    }
}