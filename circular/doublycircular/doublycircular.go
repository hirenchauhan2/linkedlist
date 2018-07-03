// Package doublycircular provides implementation of Doubly Circular Linkedlist
// @author: Hiren Chauhan
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
		Prev: nil,
	}

	// fmt.Println("Node : ", node)

	// check if the head has next node (list is not empty)
	if head.Next != nil {
		// Get the node next to head
		headNext := head.Next
		fmt.Println("The Previous Node's value: ", headNext.Data)
		// make node as headNext's prev node
		node.Prev = headNext.Prev
		headNext.Prev = &node
		// make node's next node as head's next node
		node.Next = headNext
	}
	// set the new node as head's next node
	head.Next = &node
}

// InsertEnd adds the node at end of list
func InsertEnd(head *Node, data int) {
	str := fmt.Sprintf("Inside InsertEnd => (%d)", data)
	fmt.Println(str)
	// create a node
	node := Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	// check if the head is the only one node in the list
	if head.Next == nil {
		head.Next = &node
		node.Prev = head
	} else {
		// Head has next node
		cur := head.Next

		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		firstNode := head.Next

		// Traverse to end
		for {
			// we have visited the first node,
			// and we are back to the first node again
			if visitedFirstNode && cur.Next == firstNode {
				fmt.Println("End of the list...")
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
		node.Prev = cur
		node.Next = firstNode
	}
}

// InsertBefore adds a node before specified node's data
func InsertBefore(head *Node, data int, before int) (int, error) {
	// fmt.Println("Inside InsertBefore (data => ", data, " before =>", before, " )")
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		next := head.Next
		prev := head
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		firstNode := head.Next
		for {
			// we have visited the first node,
			// and we are back to the first node again
			if visitedFirstNode && next.Next == firstNode {
				fmt.Println("End of the list...")
				break
			}

			visitedFirstNode = true

			if next.Data == before {
				fmt.Println("Found data match")
				// node data found
				// create a new node
				node := Node{
					Data: data,
				}

				node.Prev = prev
				prev.Next = &node
				node.Next = next

				// done inserting
				created = 1
				break
			}
			prev = next
			next = next.Next
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
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		next := head.Next
		prev := next
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		firstNode := head.Next
		for {

			// we have visited the first node,
			// and we are back to the first node again, we have not
			// found the node to add new node after
			if visitedFirstNode && next.Next == firstNode {
				fmt.Println("End of the list...")
				break
			}
			visitedFirstNode = true

			if next.Data == after {
				fmt.Println("Found data match")
				// node data found
				// create a new node
				node := Node{
					Data: data,
					Next: next.Next,
				}
				node.Prev = prev
				next.Next = &node
				// done inserting
				created = 1
				break
			}
			prev = next
			next = next.Next
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
		return errors.New("This node is not the header")
	}

	// Check if the list is empty or not
	if head.Next != nil {
		node := head.Next
		listStr := "[ "
		// flag for the first node's visit, if we reached to it again
		// we have traversed the linked list
		visitedFirstNode := false
		firstNode := head.Next

		for {
			// if we have visited the first node,
			// and we are back to the first node again then break out of loop
			if visitedFirstNode && node == firstNode {
				fmt.Println("End of the list...")
				break
			}
			visitedFirstNode = true

			if node != nil {
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
		return nil
	}
	return errors.New("The List is empty")
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
	if head == nil && head.Next == nil {
		fmt.Println("The list is empty")
		return false, -1
	}
	// TODO: Finish this
	deleted := head.Next.Data
	next := head.Next.Next
	next.Prev = head
	head.Next = next
	return true, deleted
}
