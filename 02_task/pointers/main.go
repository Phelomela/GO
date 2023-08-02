package main

import "fmt"

func main() {
	var (
		B int = 1
		A *int
	)
	A = &B
	fmt.Println(*A)
	*A = 2
	fmt.Println(B)
}
