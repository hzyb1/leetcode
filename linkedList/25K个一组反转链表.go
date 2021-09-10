package linkedList

type ListNode struct {
 	Val int
 	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next:head}
	p:=head
	kPre := hair
	for p!=nil {
		pre:=p
		i:=0
		for pre!=nil && i<k{
			pre = pre.Next
			i++
		}
		if i!=k {
			break
		}
		cur := p
		i = 0
		for cur!=nil && i<k {
			i++
			tmpNext := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmpNext
		}
		kPre.Next = pre
		kPre = p
		p = cur
	}
	return hair.Next
}