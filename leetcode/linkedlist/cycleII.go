package linkedlist

// https://leetcode.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {

	
    if head == nil || head.Next == nil{
        return nil
    }

    single := head.Next
    double := head.Next.Next
    cycle := false
    
    for double != nil && double.Next != nil && double.Next.Next != nil{
        single = single.Next
        double = double.Next.Next
        
        if single == double{
          cycle = true
          break
        }
    }
    
    if cycle == true{
    
    single = head
    for single != double{
        single = single.Next
        double = double.Next
    }
    
    return single
    }
    
    return nil
}
