package solutions

import "fmt"

// ListNode is defined in the LeetCode problem.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode converts a slice of integers into a ListNode.
func NewListNode(numbers []int) *ListNode {
	if len(numbers) == 0 {
		return nil
	} else {
		return &ListNode{
			Val:  numbers[0],
			Next: NewListNode(numbers[1:]),
		}
	}
}

// Equals checks whether two ListNodes are equal.
func (ln *ListNode) Equals(other *ListNode) bool {
	if ln == nil && other == nil {
		return true
	} else if ln == nil || other == nil {
		return false
	} else {
		return ln.Val == other.Val && ln.Next.Equals(other.Next)
	}
}

// AsString returns a string representation of a ListNode.
func (ln *ListNode) AsString() string {
	if ln == nil {
		return ""
	} else if ln.Next == nil {
		return fmt.Sprintf("%d", ln.Val)
	} else {
		return fmt.Sprintf("%d -> %s", ln.Val, ln.Next.AsString())
	}
}

// unpackNode is a helper function for NextNode.
// It returns the value and the next node of a ListNode.
func unpackNode(node *ListNode) (int, *ListNode) {
	if node == nil {
		return 0, nil
	} else {
		return node.Val, node.Next
	}
}

// NextNode is a helper function for AddTwoNumbers.
// It returns the next node of the sum of two nodes and a remainder, if any.
func NextNode(node1 *ListNode, node2 *ListNode, remainder int) *ListNode {
	n1Val, n1Next := unpackNode(node1)
	n2Val, n2Next := unpackNode(node2)
	sum := n1Val + n2Val + remainder
	value, newRemainder := sum%10, sum/10

	if n1Next == nil && n2Next == nil && newRemainder == 0 {
		return &ListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		return &ListNode{
			Val:  value,
			Next: NextNode(n1Next, n2Next, newRemainder),
		}
	}
}

// AddTwoNumbers is the second LeetCode problem:
// - https://leetcode.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return NextNode(l1, l2, 0)
}
