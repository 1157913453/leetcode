package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m1 := map[*ListNode]bool{}
    for i := headA;i != nil; i=i.Next {
        m1[i] = true
    }
    for i:= headB;i!= nil;i=i.Next {
        if m1[i] {
            return i
        }
    }
    return nil
}