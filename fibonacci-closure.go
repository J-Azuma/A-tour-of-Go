package main

import "fmt"

func fibonacci() func() int {
	prev, next := 0, 1
	return func() int {
		prev, next = next, prev+next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
