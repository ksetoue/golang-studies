package main

import "fmt"

func main() {
	h := make(hash)

	// h.insertItem(1, 2)
	// h.insertItem(10, 29)
	// h.insertItem(4, 3)
	// h.insertItem(6, 7)
	h.insertItem("test", 1)
	h.insertItem("test2", 2)
	h.insertItem("test3", 3)
	h.insertItem("test4", 4)

	h.printMap()

	h.deleteItem("test")
	fmt.Println("\n after deletion")
	h.printMap()

}
