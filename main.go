package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	minus(1)
	plus()
	Random()

	for i := range []int{1, 2, 3} {
		fmt.Print(i)
	}
}

func plus() {
	fmt.Print("Hello World")
}

func minus(a int) {
	fmt.Print("Hello World")
}
