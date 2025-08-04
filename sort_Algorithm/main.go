package main

import "fmt"

func main() {

	// GetUserDay()
	
	// Сортуваня слацсу строк за алфовітом за допомогою Bubble Sorrt ( Бульбашкове сортуваня )

	leters := []string{"Банан", "Ава", "Лимон", "Персик", "Ківі", "z", "r"}

	sortLeters := SortLetes(leters)

	fmt.Println("No sort leters:", leters)

	fmt.Println("Sorting letters alphabetically:", sortLeters)

	// Функція яка сортує слайс за допомогою алгоритму QuickSort ( Швидке сортування )

	quickSortStart()
}
