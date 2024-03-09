package piscine

func ListReverse(l *List) {
	var p *NodeL = nil
	c := l.Head
	var n *NodeL = nil

	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}

	l.Head = p
}
