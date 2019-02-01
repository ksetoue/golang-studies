package main

type Node struct {
	value   int
	pointer *Node
}

type List struct {
	length int
	start  *Node
}

func (l *List) insertItem(value int) {

}
