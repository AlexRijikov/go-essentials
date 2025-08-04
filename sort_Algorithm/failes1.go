package main

// import "fmt"

// Сортування слайсу строк за алфовітом

func SortLetes(leters []string) []string {

	arrStr := make([]string, len(leters))
	copy(arrStr, leters)

	lenArrStr := len(arrStr)

	// bubble sort - бульбашковтй метод порівняння
	for i := 0; i < lenArrStr-1; i++ {
		for j := 0; j < lenArrStr-i-1; j++ {
			// fmt.Println("f",lenArrStr-i-1) // Вивиде на єкран скільки циклів проведенно в кожній ітерації 
			if arrStr[j] > arrStr[j+1] { // Оператор > для рядків у Go означає лексикографічне порівняння, тобто:

				// "яблуко" > "банан" — true, бо "я" пізніше за "б" у алфавіті.

				// "груша" < "яблуко" — true, бо "г" раніше за "я".

				arrStr[j], arrStr[j+1] = arrStr[j+1], arrStr[j]
			}
		}
	}
	return arrStr

}
