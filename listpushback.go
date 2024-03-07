package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	i := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = i
		l.Tail = i
	} else {
		l.Tail.Next = i
		l.Tail = i
	}
}
