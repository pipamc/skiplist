package skiplist

import "math/rand"

const MaxLevel = 20

type Comparable interface {
	Compare(lhs, rhs interface{}) bool
}
// each node struct, contains next multiple points and value
type Node struct {
	nextPointers []*Node
	key interface{}
	value interface{}
}

type SkipList struct {
	head Node
	level int
	keyFunc Comparable
	randSource rand.Source
	length int64
}


func New(keyFunc Comparable) *SkipList {
	return &SkipList{
		head: Node{nextPointers: make([]*Node, MaxLevel)},
		level: MaxLevel,
		keyFunc: keyFunc,
		randSource: defaultRandSource{},
	}
}

//Gets list length
func (sl *SkipList) Len() int64 {
	return sl.length
}

// Gets skiplist head
func (sl *SkipList) Head() Node {
	return sl.head
}


// getPrevNode
func (sl *SkipList) getPrevNodes(key interface{}) []*Node {
	prev := &sl.head
	var next *Node
	for i := sl.level - 1; i >= 0; i-- {
		next = prev.nextPointers[i]
		for next != nil && (sl.keyFunc.Compare(key, next.key)) {
			prev = next
		}
	}
}

//Set a value in skiplist
func (sl *SkipList) Set(key, value interface{}) error {
	prevs := sl.
}

