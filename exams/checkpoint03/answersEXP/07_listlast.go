package piscine

// ListLast returns the data stored in the last node of l.
func ListLast(l *List) interface{} {
	if l == nil || l.Head == nil {
		return nil
	}

	if l.Tail != nil {
		return l.Tail.Data
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
