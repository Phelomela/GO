package main

import "fmt"

func main() {
	weekend := []string{"Сб", "Вскр"}
	work := []string{"Пн", "Вт", "Ср", "Чт", "Пт"}
	week := append(work, weekend...)
	fmt.Println(week)
}
