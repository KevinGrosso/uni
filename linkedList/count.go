package linkedList

/*
	CountAllElements counts and returns the number of nodes in the list l
*/
func CountAllElements(l List) int {
	counter := 0
	for l != nil {
		counter++
		l = l.Next
	}
	return counter
}
