package linkedList

import "fmt"

/*
	RecPrintAllElementsNL prints recursively in stdout every node of the
	List l. Every node is printed on a new line.
*/
func RecPrintAllElementsNL(l List) {
	if l != nil {
		fmt.Println(l.Content)
		l = l.Next
		RecPrintAllElementsNL(l)
	}
}

/*
	ItPrintAllElementsArr prints recursively in stdout every node of the List l.
	The nodes are printed all in the same line, separated by an arrow.
*/
func RecPrintAllElementsArr(l List) {
	if l != nil {
		fmt.Print(l.Content)
		l = l.Next
		if l != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
