package medium

func oddEvenList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    even := head.Next
    pe := even
    p := head
    for  p.Next != nil {
        p.Next = p.Next.Next
        if p.Next != nil {
            p = p.Next
        }
        if pe != nil {
            if pe.Next != nil{
                pe.Next = p.Next
            }
            pe = pe.Next
        }

    }

    p.Next = even
    return head
}
