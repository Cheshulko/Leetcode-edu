package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	fakeHead := &ListNode{
		Val:  -5001,
		Next: head,
	}

	for prev := fakeHead; prev.Next != nil; {
		cur := prev.Next

		if prev.Val <= cur.Val {
			prev = prev.Next
			continue
		}

		prev.Next = prev.Next.Next

		for curHead := fakeHead; curHead != nil; curHead = curHead.Next {
			if curHead.Next == nil || curHead.Next.Val >= cur.Val {
				curHead.Next, cur.Next = cur, curHead.Next
				break
			}
		}

	}

	return fakeHead.Next
}
