package main

import "fmt"

func main() {

	//ScanUserInfo() // Робота з користувачем

	//PrintZeroValues() // Робота з нульовими значеннями

	//PrintVarbsFormatting() //  Робота дієслів форматування з ЗАГАЛЬНИМИ типа

	//VarbsForWholeNmders() //  Робота дієслів форматування з Цілими числами

	// VarbsForString()  // Робота дієслів форматування з Страками

	// LogicalVarbs()  // Логічні дієслова  (так само як при використанні %v)

	// VarbsForFloat()// Дієслова форматування Float

	//  WorkingWithAnArray() // Робота з масивами

	//  printNameFerstName("Katya")
	//  printNameFerstName("Alex")
	//  printNameFerstName("Monika")

	// Створив слайс інтів який має різні числоа
	numbers := []int{9, 3, 22, 12, 7}

	// Створив зміну сорт яка дорівнює  функцію сортслайс та має в собі параметр цифр
	sortNumber := SortSlice(numbers)

	sumNumbers := SumSlice(numbers)

	sumEvenSlice := SumEvenSlice(numbers)

	minVelues := MinVelues(numbers)

	maxVelues := MaxVelues(numbers)

	// Вивід слайсів до сортування та після сортування
	fmt.Println("Original slice: ", numbers)
	fmt.Println("Sorted slice:", sortNumber)
	// Вивід суми всях чисел 
	fmt.Println("Sum:",sumNumbers)
	// Вивід суми парних цифр
	fmt.Println("Sum Even Numbers: ",sumEvenSlice)
// Вивід мінам і макс значення 

fmt.Println("Min Velues :",minVelues)
fmt.Println("Max Velues :",maxVelues)
}
