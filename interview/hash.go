package main

import "fmt"

type hash map[int]int

func (h hash) insertItem(k int, value int) {
	h[k] = value
}

func (h hash) printMap() {
	for key, value := range h {
		fmt.Println("key:", key, "->", " value:", value)
	}
}

// func searcPairs()
