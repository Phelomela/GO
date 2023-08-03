package main

import "fmt"

func main() {
	fmt.Println(contains([]string{"a", "b", "c", "d"}, "c"))
	fmt.Println(getMax(1, 2, 3, 4, 5, 6, 7))
	Max := func(numbers ...int) (max int) {
		max = numbers[0]
		for _, i := range numbers {
			if max < i {
				max = i
			}
		}
		return
	}
	fmt.Println(Max(1, 2, 3, 4, 5, 6, -7))

}

func contains(a []string, x string) (flag bool) {
	flag = false
	for _, str := range a {
		if str == x {
			flag = true
			break
		}
	}

	return
}

func getMax(numbers ...int) (max int) {
	max = numbers[0]
	for _, i := range numbers {
		if max < i {
			max = i
		}
	}
	return
}
