package main

import "fmt"

func main() {
	var (
		books map[string]map[string][]int = map[string]map[string][]int{}
	)
	books["Старик"] = map[string][]int{
		"Пророк": {1, 2, 3},
		"Ищейка": {3, 4, 5},
	}

	books["Молодой"] = map[string][]int{
		"Сила":  {3, 25, 36},
		"Воля":  {33, 431, 5134},
		"Слава": {3, 4, 6},
	}

	fmt.Println(len(books))
	for man, manBooks := range books {
		fmt.Println(man, "имеет", len(manBooks), "книги.")
	}
}
