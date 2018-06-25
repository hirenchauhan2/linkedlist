package main

import (
	"fmt"

	"github.com/hirenchauhan2/linkedlist/singly"
)

func main() {
	singlyLinkedlistDemo()
	// doublyLinkedlistDemo()
	// singlyCircularLinkedlistDemo()
	// doublyCircularLinkedlistDemo()
}

// singlyLinkedlistDemo demonstrate the Singly Linked list
func singlyLinkedlistDemo() {
	fmt.Println("----------------------------------------")
	fmt.Println("\tSingly Linked List Demo")
	fmt.Println("----------------------------------------")

	fmt.Println("Create a new List")
	fmt.Println("----------------------------------------")

	// create a head node
	head := singly.Node{
		Next: nil,
	}
	fmt.Println("Adding elements at begining")
	fmt.Println("----------------------------------------")

	// add a node at front
	fmt.Println("Add => 10")
	singly.InsertFront(&head, 10)
	fmt.Println("Add => 8")
	singly.InsertFront(&head, 8)
	fmt.Println("Add => 6")
	singly.InsertFront(&head, 6)
	fmt.Println("Add => 4")
	singly.InsertFront(&head, 4)
	fmt.Println("Add => 2")
	singly.InsertFront(&head, 2)
	fmt.Println("Add => 0")
	singly.InsertFront(&head, 0)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	// traverse a list and print elements
	singly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding elements at end")
	fmt.Println("----------------------------------------")
	// add node to last
	fmt.Println("Add => 12")
	singly.InsertEnd(&head, 12)
	fmt.Println("Add => 14")
	singly.InsertEnd(&head, 14)
	fmt.Println("Add => 16")
	singly.InsertEnd(&head, 16)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")

	// traverse a list and print elements
	singly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element before 14")
	inserted := singly.InsertBefore(&head, 13, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 13, "before 14")
	} else {
		fmt.Println("Not Inserted")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding 15 after 14")
	inserted = singly.InsertAfter(&head, 15, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 15, "after 14")
	} else {
		fmt.Println("Not Inserted")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Searching elements")
	fmt.Println("----------------------------------------")

	fmt.Printf("Search 15: ")
	found, nodeData := singly.Find(head, 15)
	if found {
		fmt.Printf("Found %d\n", nodeData)
	} else {
		fmt.Printf("Node Not found!\n")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the elements from front: ")
	deleted, data := singly.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("The list is empty")
	}
	deleted, data = singly.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
		fmt.Println("----------------------------------------")
		fmt.Println("Traversing the list...")
		fmt.Println("----------------------------------------")
		singly.TraverseList(head)
	} else {
		fmt.Println("The list is empty")
	}
	fmt.Println("----------------------------------------")

	fmt.Println("Deleting the elements from end: ")

	deleted, data = singly.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("The list is empty")
	}

	deleted, data = singly.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
		fmt.Println("----------------------------------------")
		fmt.Println("Traversing the list...")
		fmt.Println("----------------------------------------")
		singly.TraverseList(head)

	} else {
		fmt.Println("Could not found the node or the list is empty")
	}
	fmt.Println("----------------------------------------")

	fmt.Println("Deleting the elements randomly: ")
	deleted, data = singly.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}
	deleted, data = singly.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singly.TraverseList(head)
	fmt.Println("----------------------------------------")

}
