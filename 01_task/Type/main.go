package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var (
		number string = "104"
		z      int    = 35
	)
	a := strconv.Itoa(z)
	b, _ := strconv.Atoi(number)
	fmt.Println("104 is", reflect.TypeOf(a), ", 35 is", reflect.TypeOf(b))
}
