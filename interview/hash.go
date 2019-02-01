package main

import "fmt"

type hash map[string]int

func (h hash) insertItem(k string, value int) {
	h[k] = value
}

func (h hash) printMap() {
	for key, value := range h {
		fmt.Println("key:", key, "->", " value:", value)
	}
}

func (h hash) deleteItem(k string) {
	delete(h, k)
}

// func searcPairs()
