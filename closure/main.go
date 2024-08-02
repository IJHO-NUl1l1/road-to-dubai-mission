package main

import (
	"fmt"
)

func outer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	increment := outer()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
