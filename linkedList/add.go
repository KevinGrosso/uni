package linkedList

/*
	AddFront creates a new node with content s in front of
	the List l, and returns the new List (a List is just a
	pointer to the first node).
*/
func AddFront(s string, l List) List{
	var n List
	n = new(Node)
	n.Content = s
	n.Next = l
	return n
}

/*
	AddTail creates a new node with content s and appends it
	as the last node of the returned List l.
*/
func AddTail(s string, l List) List{
	var n List
	n = new(Node)
	n.Content = s
	n.Next = nil
	if l == nil {
		return n
	}
	first := l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = n
	return first
}

/*
	AddInOrder creates a new node with content s and appends
	it in the right position of List l, so that if l was
	lexicographically sorted before, it remains sorted.
	The new list is then returned.
*/
func AddInOrder(s string, l List) List {
	var n List
	n = new(Node)
	n.Content = s
	var prev, curr List
	prev = nil
	curr = l
	for curr != nil && curr.Content < s {
		prev = curr
		curr = curr.Next
	}
	if prev == nil {
		n.Next = curr
		return n
	}
	prev.Next = n
	n.Next = curr
	return l
}

/*
	AddAtIndex creates a new node with content s and inserts it in the list l,
	in the position specified by the index i. Then new list is then returned
*/
func AddAtIndex(s string, i int, l List) List {
	n := new(Node)
	n.Content = s

	len := CountAllElements(l)

	if l == nil {
		return n
	} else if i == len {
		return AddTail(s, l)
	} else if i == 0{
		return AddFront(s, l)
	}

	var prev, next *Node
	prev = l
	next = prev.Next

	// a b c . d e
	for j := 0; j < i-1; j++ {
		prev = prev.Next
		next = prev.Next
	}

	prev.Next = n
	n.Next = next

	return l
}
