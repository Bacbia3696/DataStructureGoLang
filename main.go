package main

import "DataStructureGolang/linklist"

func main() {
	l := linklist.LinkedList{}
	l.InsertBegin(1)
	l.InsertBegin(2)
	l.InsertBegin(3)
	l.InsertEnd(4)
	_ = l.Display()
}
