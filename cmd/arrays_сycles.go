package main
//const WAVE = 10

// func WorkingWithAnArray() {

// Синтаксис оголошення масиву

// 	arr := [3]int{1, 2, 3} //Фіксований розмір масиву

// 	fmt.Println(arr)

// 	//Ініціалізація масивів
// 	arrA := [5]int{}              //не ініціалізовано
// 	arrB := [5]int{1, 2}          //частково ініціалізовано
// 	arrC := [5]int{1, 2, 3, 4, 5} //повністю ініціалізовано

// 	fmt.Println(arrA)
// 	fmt.Println(arrB)
// 	fmt.Println(arrC)

// 	//Практика: створити 3 різних масивів

// 	arrCar := [5]string{"BMW", "ZAZ", "Ford", "Opel", "Renault"}

// 	arrTemp := [4]float32{12.2, 13.5, 14.3, 22.6}

// 	fmt.Println(arrCar[2]) //Доступ до елементів за індексом

// 	fmt.Println(arrTemp)

// 	// Робота з елементами масиву

// 	arrSchool := [8]string{"Class1", "Class2", "Class3", "Class4", "Class5", "Class6", "Class7", "Class8"}

// 	arrSchool[2] = "NoClass" //Зміна значень елементів

// 	fmt.Println(arrSchool)

// 	//Ітерація через масив (for, range)

// 	arrProblem := [5]string{"Engine", "Chassis", "Body", "Electricity", "Salon"}

// 	fmt.Println(len(arrProblem), "What could be the car's faults: ", arrProblem) //len() функція для масивів для визначення довжини масиву

// 	numbers := []int{10, 20, 30, 40, 50}

// 	fmt.Println("Last Value:", AlwaysGetLastValue(numbers))

// 	ConstValue()

// }

// // Практика: через цикли додати значення 10 до кожного елементу масиву через константу

// func ConstValue() {

// 	arrValue := [6]int{1, 2, 3, 4, 5, 6}

// 	if len(arrValue) == 0 {
// 		return
// 	}

// 	for i := 0; i < len(arrValue); i++ {
// 		arrValue[i] += WAVE
// 	}

// 	fmt.Println("New Valume ", arrValue)

// }

// // Вивести останній елемент масиву в окремій функції

// func AlwaysGetLastValue(numbers []int) int {

// 	return numbers[len(numbers)-1]
// }

//func printNameFerstName(fname string) {

// fmt.Println("Hello", fname, "Ryzhykov")

//}

// Cтворив функцію сортування  яка приймає слайс інтів та повертає слайс інтів
func SortSlice(numbers []int) []int {

	// Ствоив зманну ар за допомогою фунції make яка має переметр цілих чисил та довжину слайсу numbers (5)
	arr := make([]int, len(numbers))
	// За допомогою функії copy роблю комію слайса numbers в слайс arr
	copy(arr, numbers)

	// створюю допоміжну змінну n яка містить довжину скопійованого слайсу
	n := len(arr)
	// Через цикл фор записую що після кожного проходу  найбільше число записується в кінець
	// тому з кожним проходом треба робити на -1 (н-1) перевіру менше
	for i := 0; i < n-1; i++ {
		// внутрений цикл проходить по усім не пройденим цифрам
		// а умова (n-і-1) що треба пропустити вже відсортовані цифри
		for j := 0; j < n-i-1; j++ {
			// Певеріряємо два сусідні елементи якщо  ліве число більше за праве то міняемо їх місцями
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// ключове слово яке повертає результат значення arr
	return arr

}
// Функція яка приймає слайс чисел та повертає їх суму
func SumSlice(numbers []int) int {

	//Створив змінну сума и дав їй зачення 0
	sum := 0
	//  За допомогою циклу та допоміжного слова range перебераю слайс numbers
	for _, value := range numbers {
		// в змінну сума додаю значення зміної value конен раз циклу
		sum += value
	}
	// Повертаю значення циклу (суму)
	return sum

}
//Функція яка рахує тільки парні цифри
func SumEvenSlice(numbers []int) int {

	sumEven := 0

	for _, velue := range numbers {
		if velue%2 == 0 {
			sumEven += velue
		}
	}
	return sumEven
}
func MinVelues(nambers []int) int{

	minVelue := nambers[0] 

	for  i := 1;i< len(nambers);i++{
		if nambers[i]< minVelue{
			minVelue = nambers[i]
 
		}
	}
	return minVelue }

func MaxVelues( nambers []int) int {
	
    maxVelue := nambers[0]

	for i := 1;i < len(nambers);i++{
		if nambers[i]>maxVelue{
			maxVelue = nambers[i]
		}
	}
	return maxVelue

}
