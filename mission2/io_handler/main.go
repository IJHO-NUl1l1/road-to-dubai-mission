package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		fmt.Printf("Enter command: ")
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		} else if input == "hello" {
			fmt.Println("Hello, world!")
		} else if input == "even" {
			for i := 0; i <= 10; i += 2 {
				fmt.Printf("%d\n", i)
			}
			fmt.Println()
		} else if input == "odd" {
			for i := 1; i <= 10; i += 2 {
				fmt.Printf("%d\n", i)
			}
			fmt.Println()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
