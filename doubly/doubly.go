// Package doubly provides implementation of Doubly Linkedlist
package doubly

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

// InsertFront adds a node at front of the list
func InsertFront(head *Node, data int) {
	// fmt.Println("Doubly LinkedList -> InsertFront ", data)
	// create a new node
	node := Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	// fmt.Println("Node : ", node)

	// check if the head has next node
	if head.Next != nil {
		// Get the node next to head
		headNext := head.Next
		fmt.Println("The Previous Node's value: ", headNext.Data)
		// make node as headNext's prev node
		headNext.Prev = &node
		// make node's next node as head's next node
		node.Next = headNext
		node.Prev = head
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
		node.Prev = next
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

		for {
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
		for {
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
		return nil
	}
	return errors.New("The List is empty")

}

/**
* Searching the node in the link list
 */

// Find search the node in linklist
func Find(head Node, data int) (bool, int) {
	if head.Next == nil {
		return false, -1
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
	if found {
		return true, nodeData
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
		// and add new reference to head's next,
		// so we are letting GC to know that element should be GC'd
		// deleted the first element
		head.Next = next
		next.Prev = head
		return true, deleted
	}
	return false, -1
}

// DeleteEnd deletes the node from the end of list
func DeleteEnd(head *Node) (bool, int) {
	if head != nil && head.Next != nil {

		// we have only one node inlist
		if head.Next.Next == nil {
			deleted := head.Next.Data
			head.Next.Prev = nil
			head.Next = nil
			return true, deleted
		}

		// we have multiple nodes in list
		cur := head.Next
		prev := cur
		// go to end
		for {
			if cur.Next != nil {
				prev = cur
				cur = cur.Next
			} else {
				break
			}
		}
		deleted := cur.Data
		prev.Next = nil
		cur.Prev = nil
		return true, deleted
	}
	return false, -1
}

// Delete removes the node specified by its data passed in argument
func Delete(head *Node, data int) (bool, int) {
	if head != nil && head.Next != nil {
		cur := head.Next
		prev := cur
		found := false

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
			deleted := cur.Data

			prev.Next = cur.Next
			cur.Next.Prev = prev
			cur.Next = nil

			return true, deleted
		}
	}
	return false, -1
}
