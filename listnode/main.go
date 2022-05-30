package main

type ListNode struct {
	Value int
	Next *ListNode
}
func main()  {

}

//判断链表中是否有环
func hasCycle(head *ListNode) bool {
	nodeHash := make(map[*ListNode]int)
	var isCycle func(head *ListNode) bool
	isCycle = func(head *ListNode) bool{
		if head != nil {
			if _, ok := nodeHash[head]; ok {
				return true
			}
			nodeHash[head] = 1
			return isCycle(head.Next)
		}

		return false
	}

	return isCycle(head)
}

