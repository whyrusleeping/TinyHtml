package tinyhtml

//A linked list style implementation of a queue of bytes
type Queue struct {
	head, tail *sqNode
	length int
}

type sqNode struct {
	val byte
	next *sqNode
}

func (sq *Queue) Push(b byte) {
	nn := new(sqNode)
	nn.val = b
	if sq.length == 0 {
		sq.head = nn
		sq.tail = nn
	} else {
		sq.tail.next = nn
		sq.tail = nn
	}
	sq.length++
}

func (sq *Queue) PushMany(b []byte) {
	for _,v := range b {
		sq.Push(v)
	}
}

func (sq *Queue) Pop() (b byte) {
	if sq.length == 0 {
		panic("Why are you popping? Theres nothing here!")
	}
	b = sq.head.val
	sq.head = sq.head.next
	sq.length--
	return
}

func (sq *Queue) Top() byte {
	return sq.head.val
}

func (sq *Queue) Size() int {
	return sq.length
}
