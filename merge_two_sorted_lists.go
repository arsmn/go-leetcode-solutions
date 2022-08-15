package leetcodesolutions

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		output = new(ListNode)
		prev   = output
	)

	for list1 != nil || list2 != nil {
		if list1 == nil {
			prev.Next = list2
			break
		}

		if list2 == nil {
			prev.Next = list1
			break
		}

		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}

		prev = prev.Next
	}

	return output.Next
}
