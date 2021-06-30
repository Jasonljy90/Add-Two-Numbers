package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head *ListNode
	size int
}

func (p *linkedList) addNode(Val int) error {
	newNode := &ListNode{
		Val:  Val,
		Next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
	p.size++
	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	myList := &linkedList{nil, 0}
	sum := l1.Val + l2.Val
	if sum >= 10 {
		carry = 1
		sum = sum - 10
	}
	myList.addNode(sum)
	for l1.Next != nil && l2.Next != nil {
		sum = l1.Next.Val + l2.Next.Val + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			sum = sum - 10
		}
		myList.addNode(sum)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1.Next != nil {
		sum = l1.Next.Val + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			sum = sum - 10
		}
		myList.addNode(sum)
		l1 = l1.Next
	}
	for l2.Next != nil {
		sum = l2.Next.Val + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			sum = sum - 10
		}
		myList.addNode(sum)
		l2 = l2.Next
	}
	if carry == 1 {
		myList.addNode(1)
	}
	return myList.head
}
