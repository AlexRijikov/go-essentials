package main

import (
	"fmt"
)

const WAVE = 10

func WorkingWithAnArray() {

	// Синтаксис оголошення масиву

	arr := [3]int{1, 2, 3} //Фіксований розмір масиву

	fmt.Println(arr)

	//Ініціалізація масивів
	arrA := [5]int{}              //не ініціалізовано
	arrB := [5]int{1, 2}          //частково ініціалізовано
	arrC := [5]int{1, 2, 3, 4, 5} //повністю ініціалізовано

	fmt.Println(arrA)
	fmt.Println(arrB)
	fmt.Println(arrC)

	//Практика: створити 3 різних масивів

	arrCar := [5]string{"BMW", "ZAZ", "Ford", "Opel", "Renault"}

	arrTemp := [4]float32{12.2, 13.5, 14.3, 22.6}

	fmt.Println(arrCar[2]) //Доступ до елементів за індексом

	fmt.Println(arrTemp)

	// Робота з елементами масиву

	arrSchool := [8]string{"Class1", "Class2", "Class3", "Class4", "Class5", "Class6", "Class7", "Class8"}

	arrSchool[2] = "NoClass" //Зміна значень елементів

	fmt.Println(arrSchool)

	//Ітерація через масив (for, range)

	arrProblem := [5]string{"Engine", "Chassis", "Body", "Electricity", "Salon"}

	fmt.Println(len(arrProblem), "What could be the car's faults: ", arrProblem) //len() функція для масивів для визначення довжини масиву


	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Last Value:", AlwaysGetLastValue(numbers)) 

	ConstValue()
}

// Практика: через цикли додати значення 10 до кожного елементу масиву через константу

func ConstValue() {

	arrValue := [6]int{1, 2, 3, 4, 5, 6}

	if len(arrValue) == 0 {
		return
	}

	for i := 0; i < len(arrValue); i++ {
		arrValue[i] += WAVE
	}

	fmt.Println("New Valume ", arrValue)

}	

	// Вивести останній елемент масиву в окремій функції

func AlwaysGetLastValue(numbers []int) int {

	return numbers[len(numbers)-1] 

}