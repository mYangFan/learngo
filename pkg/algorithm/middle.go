package algorithm

import "fmt"

//单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//todo 解题思路
// 1.遍历两个链表，取出链表中的元素，对齐相加
// 2.算出进位
// 3.在结果链表尾部加上进位
func TwoAdd(l1, l2 *ListNode) *ListNode {
	//[7,4,3]  [8,5,9]  [5,0,3,1]
	var tail,head *ListNode
	carry := 0 //进位

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		//1.7+8+0  15
		//2.4+5+1  10
		//3.3+9+1  13
		sum := n1 + n2 + carry //这里加上进位的原因：因为链表中每个元素在相加的时候，都有可能产生进位，所以把进位加在sum里面，最后再算最高位的进位

		//1.5,1
		//2.0,1
		//3.3,1
		sum, carry = sum%10, sum/10
		//把头节点给尾节点
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		}else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val:carry}
	}

	return head
}

func ShowAll(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}


//func HeadInsert(num int, lastNode *ListNode) (listNode *ListNode) {
//	listNode = &ListNode{Val: num, Next: lastNode}
//	return
//}

//头插
func HeadInsert(nums []int) (nodeList *ListNode) {
	var lastNode *ListNode
	for k, v := range nums{
		if k == 0 {
			//头节点
			nodeList = &ListNode{Val: v, Next: nil}
			lastNode = nodeList
		}else {
			nodeList = &ListNode{Val: v, Next: lastNode}
			lastNode = nodeList
		}
	}
	return
}

//尾插
func TailInsert(nums []int) *ListNode {
	var head,tail *ListNode
	for _,v := range nums{
		if head == nil{
			head = &ListNode{Val:v}
			tail = head
		}else {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
	}
	return head
}
