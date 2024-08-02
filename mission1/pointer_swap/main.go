package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 5, 10
	fmt.Println("Before swap: a =", a, "b =", b)
	swap(&a, &b)
	fmt.Println("After swap: a =", a, "b =", b)
}
