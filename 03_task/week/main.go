package main

import "fmt"

func main() {
	weekend := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вскр"}
	work := weekend[:5]
	weekend = weekend[5:]
	fmt.Println("Рабочие дни: ", work, "\nВыходные: ", weekend)
}
