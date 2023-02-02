// Package list impliments linked list.
//
// It supports insert, search, predecessor, and delete operations
// on a list.
package list

// A sint represents a node in a linked list.
// Here, val is an integer value, and the next points
// to the next node if exists, or otherwise to the nil pointer.
type sint struct {
	val  int
	next *sint
}

// SinglyInt initiates a singly integer linked list
// with an initial value. It returns the pointer
// to the first node.
func SinglyInt(v int) (p *sint) {
	return &sint{v, nil}
}

// Insert appends a node with the given value
// into the list.
func (l *sint) Insert(v int) {
	if l.next == nil {
		l.next = &sint{v, nil}
		return
	} else {
		l.next.Insert(v)
	}
}

// Search look for the given value in
// the list. It returns the pointer to the value
// if found, or nil if doesn't.
func (l *sint) Search(v int) *sint {
	if l == nil {
		return nil
	}
	if (*l).val == v {
		return l
	} else {
		return l.next.Search(v)
	}

}

// Pred find the predecessor of the given value in
// the list. It returns the pointer to the value
// if found, or nil if doesn't.
func (l *sint) Pred(v int) *sint {
	if l == nil || l.next == nil || l.val == v {
		return nil
	}
	if (*l.next).val == v {
		return l
	} else {
		return l.next.Pred(v)
	}
}

// Delete removes the node containing the given value
// from the list, if exists, of course.
func (l *sint) Delete(v int) {
	p := l.Search(v)
	if p == nil {
		return
	}
	pred := l.Pred(v)
	if pred == nil {
		l.val = l.next.val
		l.next = l.next.next
		return
	}
	pred.next = p.next
}
