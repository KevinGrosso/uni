package linkedList

import "fmt"

/*
	ItPrintAllElementsNL prints iteratively in stdout every node of the
	List l. Every node is printed on a new line.
*/
func ItPrintAllElementsNL(l List) {
	for l != nil {
		fmt.Println(l.Content)
		l = l.Next
	}
}

/*
	ItPrintAllElementsArr prints iteratively in stdout every node of the List l.
	The nodes are printed all in the same line, separated by an arrow.
*/
func ItPrintAllElementsArr(l List) {
	for l != nil {
		fmt.Print(l.Content)
		l = l.Next

		if l != nil { // if i'm not printing the last node
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
