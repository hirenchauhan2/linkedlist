// Package doublycircular provides implementation of Doubly Circular Linkedlist
package doublycircular

import (
	"errors"
	"fmt"
	"strconv"
)

// A Node is the type of each element in the LinkedList
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

// InsertFront adds a node into the list at the end
func InsertFront(head *Node, data int) {
	// fmt.Println("Doubly LinkedList -> InsertFront ", data)
	// create a new node
	node := Node{
		Data: data,
		Next: nil,
		Prev: head,
	}

	// fmt.Println("Node : ", node)

	// check if the head has next node (list is not empty)
	if head.Next != nil {
		// Get the node next to head
		headNext := head.Next
		headNext.Prev = &node
		// make node's next node as head's next node
		node.Next = headNext
	} else {
		head.Prev = &node
		node.Next = head
	}
	// set the new node as head's next node
	head.Next = &node
}

// InsertEnd adds the node at end of list
func InsertEnd(head *Node, data int) {
	// create a new node
	node := Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
	// if head is the only node in the list
	if head.Next == nil {
		head.Next = &node
		node.Next = head
		node.Prev = head
	} else {
		// we have more nodes in the list
		// get the last node from the head (shortest way!-)
		last := head.Prev

		node.Next = last.Next
		last.Next = &node
		node.Prev = last
		head.Prev = &node
		fmt.Println("The last node's data", last.Data)

	}
}

// InsertBefore adds a node before specified node's data
func InsertBefore(head *Node, data int, before int) (int, error) {
	// fmt.Println("Inside InsertBefore (data => ", data, " before =>", before, " )")
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		cur := head.Next
		prev := cur.Prev
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		for {
			// we have visited the first node,
			// and we are back to the first node again
			if visitedFirstNode && cur.Next == head {
				fmt.Println("End of the list...")
				break
			}

			if cur.Data == before {
				fmt.Println("Found data match")
				// node data found
				// create a new node
				node := Node{
					Data: data,
					Prev: prev,
					Next: cur,
				}

				prev.Next = &node
				cur.Prev = &node
				// done inserting
				created = 1
				break
			}
			prev = cur
			cur = cur.Next
			visitedFirstNode = true
		}
	}
	if created == 1 {
		return created, nil
	}
	return 0, errors.New("Node with value " + strconv.Itoa(before) + " not found in the list.")
}

// InsertAfter adds a node after specified node's data
func InsertAfter(head *Node, data int, after int) (int, error) {
	fmt.Println("Inside InsertAfter (data => ", data, " after =>", after, " )")
	created := 1
	// check if the list is empty or not
	if head.Next != nil {
		cur := head.Next
		next := cur.Next
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		for {
			if cur.Data == after {
				fmt.Println("Found data match")
				// node data found
				// create a new node
				node := Node{
					Data: data,
				}
				// Linking the node between the current and next node
				// next will be after the new node
				node.Next = next
				// node's prev is the current
				node.Prev = cur
				// next's prev node will be the new node
				next.Prev = &node
				// current's next node will be the new node
				cur.Next = &node

				// done inserting
				created = 1
				break
			}
			// we have visited the first node,
			// and we are back to the first node again, we have not
			// found the node to add new node after
			if visitedFirstNode && next == head {
				fmt.Println("End of the list...")
				break
			}
			visitedFirstNode = true

			cur = cur.Next
			next = cur.Next
		}
	}
	if created == 1 {
		return created, nil
	}

	return 0, errors.New("Node with value " + strconv.Itoa(after) + " not found in the list.")
}

// TraverseList traverses a singly linked list
func TraverseList(head Node) error {
	// check if it is the head node
	if head.Data != 0 {
		return errors.New("this node is not the header")
	}

	// Check if the list is empty or not
	if head.Next != nil {
		node := head.Next
		listStr := "[ "
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		// visitedFirstNode := false
		firstNode := head.Next

		for {
			data := node.Data
			// convert int to string and add to listStr
			listStr = listStr + strconv.Itoa(data)
			// do we have a next node? if yes then add a ", "
			// or else break out of loop
			if node.Next != nil && node.Next.Next != firstNode {
				listStr = listStr + ", "
			} else {
				break
			}
			// // if we have visited the first node,
			// // and we are back to the first node again then break out of loop
			// if visitedFirstNode && node.Next.Next == firstNode {
			// 	fmt.Println("End of the list...")
			// 	break
			// }
			// visitedFirstNode = true
			node = node.Next
		}
		listStr = listStr + " ]"
		fmt.Println("LinkedList: ", listStr)
		return nil
	}
	return errors.New("the List is empty")
}

// Find searches the node from the linked list
func Find(head Node, data int) (bool, int) {
	if head.Next != nil {
		firstNode := head.Next
		visitedFirstNode := false
		pos := 0
		cur := firstNode
		found := false
		for {
			if data == cur.Data {
				found = true
				break
			}
			if visitedFirstNode && cur.Next == firstNode {
				break
			}
			visitedFirstNode = true
			cur = cur.Next
			pos++
		}
		if found {
			return true, cur.Data
		}
	}
	return false, -1
}

/**
* Deleting a node from the list
* There are several ways we can delete a node in the linked list
* 1. DeleteFront - Delete a node from front(remove first node after head)
* 2. DeleteEnd - Delete the last node
* 3. Delete(Node.Data) - Delete a specific node (delete any)
**/

// DeleteFront deletes the node from the front of the linked list
func DeleteFront(head *Node) (bool, int) {
	if head.Next == nil {
		fmt.Println("The list is empty")
		return false, -1
	}
	// get the first element
	first := head.Next

	// if there are no nodes left, we clear the head's next & prev reference
	if head.Prev == first {
		head.Next = nil
		head.Prev = nil
		return true, first.Data
	}

	second := head.Next.Next

	head.Next = second
	second.Prev = head

	first.Next = nil
	first.Prev = nil

	return true, first.Data
}

// DeleteEnd deletes the node from the end of the linked list
func DeleteEnd(head *Node) (bool, int) {
	// if head is the only node in the list
	if head.Next == nil && head.Prev == nil {
		return false, -1
	}

	// go to the last node
	// we can go backwards from head in doubly circular linked list!
	// The last node
	last := head.Prev
	// get the last node's data
	deleted := last.Data

	// if there are no nodes left, we clear the head's next & prev reference
	if head.Next == last {
		head.Next = nil
		head.Prev = nil

		return true, deleted
	}

	// get the second last node (last's prev)
	newlast := last.Prev

	// add head as next node to the second last node
	newlast.Next = head
	// head's prev should be second last node
	head.Prev = newlast

	last.Prev = nil
	last.Next = nil

	// last node will be garbage collected, hence deleted!
	return true, deleted
}

// Delete deletes the specified node in the linked list
func Delete(head *Node, data int) (bool, int) {
	if head.Next == nil {
		return false, -1
	}

	cur := head.Next
	next := cur.Next
	visitedFirstNode := false

	for {
		if data == cur.Data {
			// found the node
			// we have to delete this (cur) node!
			// get the prev node of the current node
			prev := cur.Prev

			// prev should now point to the next
			prev.Next = next

			// next should point back to pre
			next.Prev = prev

			cur.Next = nil
			cur.Prev = nil

			// we have no nodes left!
			if next == head && head.Prev == cur && head.Next == cur {
				head.Next = nil
				head.Prev = nil
			}

			return true, cur.Data
		}

		if visitedFirstNode && next == head {
			break
		}

		visitedFirstNode = true
		cur = cur.Next
		next = cur.Next
	}
	return false, -1
}

// printMsg utility function for printing message
//func printMsg(msg string) {
//	fmt.Println(msg)
//}
