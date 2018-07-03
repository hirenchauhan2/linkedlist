// Package singlycircular provides the implementation of singly circular linked list
// @author: Hiren Chauhan
package singlycircular

import (
	"errors"
	"fmt"
	"strconv"
)

// A Node is the type of each element in the LinkedList
type Node struct {
	Data int
	Next *Node
}

// InsertFront adds a node into the list at the end
func InsertFront(head *Node, data int) {
	// fmt.Println("LinkedList -> InsertFront ", data)
	// create a new node
	node := Node{
		Data: data,
		Next: nil,
	}

	// fmt.Println("Node : ", node)

	// check if the head has next node
	if head.Next != nil {
		// Get the node next to head
		headNext := head.Next

		// make node's next node as head's next node
		node.Next = headNext
	}
	// set the new node as head's next node
	head.Next = &node
}

// InsertEnd adds the node into start of list after the head
func InsertEnd(head *Node, data int) {
	// create a node
	node := Node{
		Data: data,
		// Set the new node's next as head's next, circle
		Next: head.Next,
	}

	// check if the head is the only one node in the list
	if head.Next == nil {
		head.Next = &node
	} else {
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		firstNode := head.Next

		// Head has next node
		cur := head.Next
		// Traverse to end
		for {
			// we have visited the first node,
			// and we are back to the first node again
			if visitedFirstNode && cur.Next == firstNode {
				break
			}

			visitedFirstNode = true
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
		// we are at the end of list, last node
		// add new node as last node's next node
		cur.Next = &node
	}
}

// InsertBefore adds a node before specified node's data
func InsertBefore(head *Node, data int, before int) (int, error) {
	// fmt.Println("Inside InsertBefore (data => ", data, " before =>", before, " )")
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		cur := head.Next
		prev := head

		firstNode := head.Next
		visitedFirstNode := false

		for {
			if visitedFirstNode && cur.Next == firstNode {
				// fmt.Println("Back to the first!")
				break
			}
			if cur.Data == before {
				// found the match!
				// create a new node
				node := Node{
					Data: data,
				}
				prev.Next = &node
				node.Next = cur

				// done inserting
				created = 1
				break
			}
			prev = cur
			cur = cur.Next
		}
	}

	if created == 1 {
		return created, nil
	}

	return 0, errors.New("Could not insert a node")
}

// InsertAfter adds a node after specified node's data
func InsertAfter(head *Node, data int, after int) (int, error) {
	fmt.Println("Inside InsertAfter (data => ", data, " after =>", after, " )")
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		cur := head.Next
		firstNode := head.Next
		visitedFirstNode := false

		for {
			if visitedFirstNode && cur.Next == firstNode {
				fmt.Println("Back to the first!")
				break
			}
			if cur.Data == after {
				fmt.Println("Found data match")
				// node data found
				// create a new node
				node := Node{
					Data: data,
					Next: cur.Next,
				}
				cur.Next = &node
				// done inserting
				created = 1
				break
			}

			cur = cur.Next
		}
	}
	if created == 1 {
		return created, nil
	}

	return 0, errors.New("Could not insert a node")
}

// TraverseList traverses a singly linked list
func TraverseList(head Node) {
	// Check if the list is empty or not
	if head.Next != nil {
		firstNode := head.Next
		node := firstNode
		listStr := "[ "
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		for {
			// we have visited the first node,
			// and we are back to the first node again
			if visitedFirstNode && node == firstNode {
				// fmt.Println("End of the list...")
				break
			}

			if node != nil {
				// we have visited the first node!
				visitedFirstNode = true
				data := node.Data
				// convert int to string and add to listStr
				listStr = listStr + strconv.Itoa(data)
				// do we have a next node? if yes then add a ", "
				// or else break out of loop
				if node.Next != nil && node.Next != firstNode {
					listStr = listStr + ", "
				} else {
					break
				}
			}
			node = node.Next
		}
		listStr = listStr + " ]"
		fmt.Println("LinkedList: ", listStr)
	} else {
		fmt.Println("Empty List: [ ]")
	}
}

// Find finds the node in the list
func Find(head Node, data int) (bool, int) {
	if head.Next != nil {
		cur := head.Next
		found := false

		firstNode := cur

		for {
			if data == cur.Data {
				found = true
				break
			} else if cur.Next == firstNode {
				break
			} else if cur.Next != nil {
				cur = cur.Next
			}
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

// DeleteFront Deletes a node from front of the linked list
func DeleteFront(head *Node) (bool, int) {
	if head != nil && head.Next != nil {
		// get the 2nd element from head
		next := head.Next.Next
		deleted := head.Next.Data

		// remove reference from head
		// add new reference to head's next,
		// so we are letting GC to know that element should be GC'd
		head.Next = next

		// go to the last node of the list

		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		firstNode := head.Next

		// Head has next node
		cur := head.Next
		// Traverse to end
		for {
			// we have visited the first node,
			// and we are back to the first node again
			if visitedFirstNode && cur.Next.Next == firstNode {
				break
			}

			visitedFirstNode = true
			cur = cur.Next
		}

		cur.Next = firstNode

		// deleted the first element
		return true, deleted
	}
	return false, -1
}

// DeleteEnd deletes the node from the end of linked list
func DeleteEnd(head *Node) (bool, int) {
	if head == nil || head.Next == nil {
		fmt.Println("The list is empty")
		return false, -2
	}

	// go to the last node of the list

	// flag for the first node's visit, if we reached to it again
	// we have traversed the linked list
	visitedFirstNode := false
	firstNode := head.Next

	// Head has next node
	cur := head.Next
	// Traverse to end
	for {
		// we have visited the first node,
		// and we are back to the first node again
		if visitedFirstNode && cur.Next.Next == firstNode {
			break
		}

		visitedFirstNode = true
		cur = cur.Next
	}
	deleted := cur.Next.Data
	// remove the firstNode's reference from the last node
	cur.Next.Next = nil
	// remove the last node by overwriting the current node's next as the first node.
	cur.Next = firstNode

	return true, deleted
}

// Delete deletes any node with the given data.
// if found, the node will be deleted
func Delete(head *Node, data int) (bool, int) {
	if head == nil && head.Next == nil {
		fmt.Println("The list is empty")
		return false, -1
	}

	// search the node
	firstNode := head.Next
	cur := firstNode
	prev := cur
	found := false
	for {
		if data == cur.Data {
			found = true
			break
		}

		if cur.Next == firstNode {
			break
		}

		prev = cur
		cur = cur.Next
	}

	if found {
		deleted := cur.Data
		// detach the current node from the list
		prev.Next = cur.Next
		// detach every reference from current node
		cur.Next = nil

		return true, deleted
	}
	return false, -1
}
