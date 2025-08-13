package main

import "fmt"

func myMessage() {

// Використання (map + slice) Та сорування серез  цикл for 

	hause := map[int][]string{

		1: {"Петров", "Сідоров"},

		2: {"Колобко"},

		3: {"Сабко", "Шкрабатько"},

		4: {"Бондар"},

		5: {"Бондар", "Усачев"},
	}

	for apartNumer, residentes := range hause {
		fmt.Println("Квартира - ", apartNumer)
		for _, name := range residentes {
			fmt.Println(" - ", name)
		}
	}
	fmt.Println()
}
