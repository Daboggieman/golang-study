package piscine

// ListReverse reverses the linked list l in place.
func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	l.Tail = l.Head
	var prev *NodeL
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
