package linkedList

func RemoveAtIndex(i int, l List) List {
	len := CountAllElements(l)

	if l == nil {
		return l
	} else if i == len-1 {
		last := l
		for j := 0; j < len-2; j ++ {
			last = last.Next
		}
		last.Next = nil
		return l
	} else if i == 0 {
		return l.Next
	}

	prev := l
	next := prev.Next.Next

	for j := 0; j < i-1; j ++ {
		prev = prev.Next
		next = prev.Next
	}

	prev.Next = next.Next

	return l
}
