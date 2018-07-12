package main

import (
	"fmt"

	"github.com/hirenchauhan2/linkedlist/circular/doublycircular"
	"github.com/hirenchauhan2/linkedlist/circular/singlycircular"
	"github.com/hirenchauhan2/linkedlist/doubly"
	"github.com/hirenchauhan2/linkedlist/singly"
)

func main() {
	singlyLinkedlistDemo()
	doublyLinkedlistDemo()
	singlyCircularLinkedlistDemo()
	doublyCircularLinkedlistDemo()
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

func doublyLinkedlistDemo() {
	fmt.Println("----------------------------------------")
	fmt.Println("\tDoubly Linkedlist Demo")
	fmt.Println("----------------------------------------")

	fmt.Println("Create a new List")
	fmt.Println("----------------------------------------")

	// create a head node
	head := doubly.Node{
		Next: nil,
		Prev: nil,
	}
	fmt.Println("Adding elements at begining")
	fmt.Println("----------------------------------------")

	// add a node at front
	fmt.Println("Add => 10")
	doubly.InsertFront(&head, 10)
	fmt.Println("Add => 8")
	doubly.InsertFront(&head, 8)
	fmt.Println("Add => 6")
	doubly.InsertFront(&head, 6)
	fmt.Println("Add => 4")
	doubly.InsertFront(&head, 4)
	fmt.Println("Add => 2")
	doubly.InsertFront(&head, 2)
	fmt.Println("Add => 0")
	doubly.InsertFront(&head, 0)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	// traverse a list and print elements
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding elements at end")
	fmt.Println("----------------------------------------")
	// add node to last
	fmt.Println("Add => 12")
	doubly.InsertEnd(&head, 12)
	fmt.Println("Add => 14")
	doubly.InsertEnd(&head, 14)
	fmt.Println("Add => 16")
	doubly.InsertEnd(&head, 16)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")

	// traverse a list and print elements
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element before 14")
	inserted, err := doubly.InsertBefore(&head, 13, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 13, "before 14")
	} else {
		fmt.Println("Not Inserted")
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element after 14")
	inserted, err = doubly.InsertAfter(&head, 15, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 15, "after 14")
	} else {
		fmt.Println("Not Inserted")
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Searching elements")
	fmt.Println("----------------------------------------")

	fmt.Printf("Search 15: ")
	found, nodeData := doubly.Find(head, 15)
	if found {
		fmt.Printf("Found %d\n", nodeData)
	} else {
		fmt.Printf("Not found!\n")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Deleting nodes from front: ")
	fmt.Println("----------------------------------------")

	deleted, data := doubly.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	deleted, data = doubly.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting nodes from end: ")
	fmt.Println("----------------------------------------")
	deleted, data = doubly.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	deleted, data = doubly.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the elements randomly: ")
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the element 12: ")

	deleted, data = doubly.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}

	fmt.Println("Deleting the element 12: ")

	deleted, data = doubly.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doubly.TraverseList(head)
	fmt.Println("----------------------------------------")
}

// singlyCircularLikedlistDemo demonstrate the Singly Circular linked list
func singlyCircularLinkedlistDemo() {
	fmt.Println("----------------------------------------")
	fmt.Println("\tSingly Circular Linkedlist Demo")
	fmt.Println("----------------------------------------")

	fmt.Println("Create a new Singly Circular Linkedlist")
	fmt.Println("----------------------------------------")

	// create a head node
	head := singlycircular.Node{
		Next: nil,
	}

	fmt.Println("Adding elements at begining")
	fmt.Println("----------------------------------------")

	// add a node at front
	fmt.Println("Add => 10")
	singlycircular.InsertFront(&head, 10)
	fmt.Println("Add => 8")
	singlycircular.InsertFront(&head, 8)
	fmt.Println("Add => 6")
	singlycircular.InsertFront(&head, 6)
	fmt.Println("Add => 4")
	singlycircular.InsertFront(&head, 4)
	fmt.Println("Add => 2")
	singlycircular.InsertFront(&head, 2)
	fmt.Println("Add => 0")
	singlycircular.InsertFront(&head, 0)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	// traverse a list and print elements
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding elements at end")
	fmt.Println("----------------------------------------")
	// add node to last
	fmt.Println("Add => 12")
	singlycircular.InsertEnd(&head, 12)
	fmt.Println("Add => 14")
	singlycircular.InsertEnd(&head, 14)
	fmt.Println("Add => 16")
	singlycircular.InsertEnd(&head, 16)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")

	// traverse a list and print elements
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element before 14")
	inserted, _ := singlycircular.InsertBefore(&head, 13, 14)
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
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element after 14")
	inserted, _ = singlycircular.InsertAfter(&head, 15, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 15, "after 14")
	} else {
		fmt.Println("Not Inserted")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Searching elements")
	fmt.Println("----------------------------------------")

	fmt.Printf("Search 15: ")
	found, nodeData := singlycircular.Find(head, 15)
	if found {
		fmt.Printf("Found %d\n", nodeData)
	} else {
		fmt.Printf("Not found!\n")
	}
	fmt.Printf("Search 20: ")
	found, nodeData = singlycircular.Find(head, 20)
	if found {
		fmt.Printf("Found %d\n", nodeData)
	} else {
		fmt.Printf("Not found!\n")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting nodes from front: ")
	fmt.Println("----------------------------------------")

	deleted, data := singlycircular.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	deleted, data = singlycircular.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting nodes from end: ")
	fmt.Println("----------------------------------------")
	deleted, data = singlycircular.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	deleted, data = singlycircular.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the elements randomly: ")
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the element 12: ")

	deleted, data = singlycircular.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}

	fmt.Println("Deleting the element 12: ")

	deleted, data = singlycircular.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	singlycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
}

// doublyCircularLikedlistDemo demonstrate the Singly Circular linked list
func doublyCircularLinkedlistDemo() {
	fmt.Println("----------------------------------------")
	fmt.Println("\tDoubly Circular Linkedlist Demo")
	fmt.Println("----------------------------------------")

	fmt.Println("Create a new Singly Circular Linkedlist")
	fmt.Println("----------------------------------------")

	// create a head node
	head := doublycircular.Node{
		Next: nil,
	}

	fmt.Println("Adding elements at beginning")
	fmt.Println("----------------------------------------")

	// add a node at front
	fmt.Println("Add => 10")
	doublycircular.InsertFront(&head, 10)
	fmt.Println("Add => 8")
	doublycircular.InsertFront(&head, 8)
	fmt.Println("Add => 6")
	doublycircular.InsertFront(&head, 6)
	fmt.Println("Add => 4")
	doublycircular.InsertFront(&head, 4)
	fmt.Println("Add => 2")
	doublycircular.InsertFront(&head, 2)
	fmt.Println("Add => 0")
	doublycircular.InsertFront(&head, 0)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	// traverse a list and print elements
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding elements at end")
	fmt.Println("----------------------------------------")
	// add node to last
	fmt.Println("Add => 12")
	doublycircular.InsertEnd(&head, 12)
	fmt.Println("Add => 14")
	doublycircular.InsertEnd(&head, 14)
	fmt.Println("Add => 16")
	doublycircular.InsertEnd(&head, 16)
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")

	// traverse a list and print elements
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element before 14")
	inserted, err := doublycircular.InsertBefore(&head, 13, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 13, "before 14")
	} else {
		fmt.Println("Not Inserted", err)
	}
	fmt.Println("Adding element before 10")
	inserted, err = doublycircular.InsertBefore(&head, 9, 10)
	if inserted == 1 {
		fmt.Println("Inserted", 9, "before 10")
	} else {
		fmt.Println("Not Inserted", err)
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Adding element 15 after 14")

	inserted, _ = doublycircular.InsertAfter(&head, 15, 14)
	if inserted == 1 {
		fmt.Println("Inserted", 15, "after 14")
	} else {
		fmt.Println("Not Inserted")
	}
	fmt.Println("Adding element 11 after 10")

	inserted, _ = doublycircular.InsertAfter(&head, 11, 10)
	if inserted == 1 {
		fmt.Println("Inserted", 11, "after 10")
	} else {
		fmt.Println("Not Inserted")
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Done adding elements")
	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Searching elements")
	fmt.Println("----------------------------------------")

	fmt.Printf("Search 15: ")
	found, nodeData := doublycircular.Find(head, 15)
	if found {
		fmt.Printf("Found %d\n", nodeData)
	} else {
		fmt.Printf("Not found!\n")
	}
	fmt.Printf("Search 20: ")
	found, nodeData = doublycircular.Find(head, 20)
	if found {
		fmt.Printf("Found %d\n", nodeData)
	} else {
		fmt.Printf("Not found!\n")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Deleting nodes from front: ")
	fmt.Println("----------------------------------------")

	deleted, data := doublycircular.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}

	deleted, data = doublycircular.DeleteFront(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting nodes from end: ")
	fmt.Println("----------------------------------------")

	deleted, data = doublycircular.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}

	deleted, data = doublycircular.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}

	deleted, data = doublycircular.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}

	deleted, data = doublycircular.DeleteEnd(&head)
	if deleted {
		fmt.Println(data, "deleted successfully!")
	} else {
		fmt.Println("Empty list!")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")

	fmt.Println("Deleting the elements randomly: ")
	fmt.Println("----------------------------------------")
	fmt.Println("Deleting the element 12: ")

	deleted, data = doublycircular.Delete(&head, 12)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}

	fmt.Println("Deleting the element 12: ")

	deleted, data = doublycircular.Delete(&head, 8)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}

	fmt.Println("Deleting the element 10 ")

	deleted, data = doublycircular.Delete(&head, 10)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}

	fmt.Println("Deleting the element 6")

	deleted, data = doublycircular.Delete(&head, 6)
	if deleted {
		fmt.Println(data, "Deleted successfully!")
	} else {
		fmt.Println("Could not found the node or the list is empty")
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Traversing the list...")
	fmt.Println("----------------------------------------")
	doublycircular.TraverseList(head)
	fmt.Println("----------------------------------------")
}
