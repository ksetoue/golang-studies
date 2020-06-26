package main

import(
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 1; i < n; i++ {
		x,y = y, x+y
		fmt.Println(x)
	}
	return x
}

func main() {
	for _, arg := range os.Args[1:] {
		s, err := strconv.Atoi(arg)
		fmt.Println(err)
		fib(s)
	}
}