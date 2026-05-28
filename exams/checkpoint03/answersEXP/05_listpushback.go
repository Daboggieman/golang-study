package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushBack appends a new node with data to the end of l.
func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}

	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	if l.Tail == nil {
		l.Tail = l.Head
	}

	l.Tail.Next = newNode
	l.Tail = newNode
}
