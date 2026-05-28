package piscine

// ListRemoveIf removes all nodes from l whose Data equals data_ref.
func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	l.Tail = current
}
