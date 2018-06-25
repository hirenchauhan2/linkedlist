// Package singly provides the implementation of Singly Linkedlist
// Implementation of Singly Linked List
// @author: Hiren Chauhan
package singly

import (
	"fmt"
	"strconv"
)

// A Node is the type of each element in the LinkedList
type Node struct {
	Data int
	Next *Node
}

/** Inserting nodes into the linked list
* There are 4 ways we can insert a node into the linked list
* 1. InsertFront - Add node at the begining
* 2. InsertEnd -  Add node at the end of list
* 3. InsertBefore - Add a node before a specified node
* 4. InsertAfter - Add a node after a specified node
 */

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
		Next: nil,
	}
	// check if the head is the only one node in the list
	if head.Next == nil {
		head.Next = &node
	} else {
		// Head has next node
		next := head.Next
		// Traverse to end
		for {
			if next.Next != nil {
				next = next.Next
			} else {
				break
			}
		}
		// we are at the end of list, last node
		// add new node as last node's next node
		next.Next = &node
	}
}

// InsertBefore adds a node before specified node's data
func InsertBefore(head *Node, data int, before int) int {
	fmt.Println("Inside InsertBefore (data => ", data, " before =>", before, " )")
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		next := head.Next
		prev := head

		for {
			if next.Data == before {
				// fmt.Println("Found data match")
				// create a new node
				node := Node{
					Data: data,
				}
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

	return created
}

// InsertAfter adds a node after specified node's data
func InsertAfter(head *Node, data int, after int) int {
	fmt.Println("Inside InsertAfter (data => ", data, " after =>", after, " )")
	created := 0
	// check if the list is empty or not
	if head.Next != nil {
		next := head.Next
		for {
			if next.Data == after {
				fmt.Println("Found data match")
				// node data found
				// create a new node
				node := Node{
					Data: data,
					Next: next.Next,
				}
				next.Next = &node
				// done inserting
				created = 1
				break
			}

			next = next.Next
		}
	}

	return created
}

/**
* Traverse the list to the end and print out each node's data
 */

// TraverseList traverses a singly linked list
func TraverseList(head Node) {
	// Check if the list is empty or not
	if head.Next != nil {
		node := head.Next
		listStr := "[ "
		for {
			if node != nil {
				data := node.Data
				// convert int to string and add to listStr
				listStr = listStr + strconv.Itoa(data)
				// do we have next node? if yes then add ,
				// or else break out of loop
				if node.Next != nil {
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
		fmt.Println("[ ]")
	}
}

/**
* Searching the node in the link list
 */

// Find searches the node in linklist
func Find(head Node, data int) (bool, int) {
	if head.Next == nil {
		return false, 0
	}
	found := false
	cur := head.Next
	var nodeData int
	for {
		if cur.Data == data {
			nodeData = cur.Data
			found = true
			break
		}
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	return found, nodeData
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
		head.Next = nil
		// add new reference to head's next,
		// so we are letting GC to know that element should be GC'd
		head.Next = next
		// deleted the first element
		return true, deleted
	}
	return false, -1
}

// DeleteEnd deletes a node from end of the list
func DeleteEnd(head *Node) (bool, int) {
	if head != nil && head.Next != nil {
		next := head.Next

		// only one node in the list after head
		if next.Next == nil {
			deleted := head.Next.Data
			head.Next = nil
			return true, deleted
		}

		// list has more nodes, go to the end
		for {
			if next.Next != nil && next.Next.Next != nil {
				next = next.Next
			} else {
				break
			}
		}
		// get the data of deleted node
		deleted := next.Next.Data
		// set the pointer to nil
		next.Next = nil
		return true, deleted
	}
	return false, -1
}

// Delete delete any node by it's data
func Delete(head *Node, data int) (bool, int) {
	if head != nil && head.Next != nil {
		// get the first node
		cur := head.Next
		prev := cur
		found := false
		deleted := -1
		// iternate the list
		for {
			if data == cur.Data {
				found = true
				break
			} else if cur.Next != nil {
				prev = cur
				cur = cur.Next
			} else {
				break
			}
		}
		if found {
			deleted = cur.Data
			// remove the reference of cur node
			prev.Next = cur.Next
			cur.Next = nil
			return true, deleted
		}
	}
	return false, -1
}
