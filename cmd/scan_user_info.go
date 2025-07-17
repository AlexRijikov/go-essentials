package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func ScanUserInfo() {
	userName := getUserName()

	userAge := getUserAge()

	userHeight := getUserHeight()

	fmt.Println("Helo,", userName, ",your age is", userAge, "years old.", "Your height:", userHeight, "cm")
	fmt.Printf("Hello %s,your age is %d years old. Your height %f cm ", userName, userAge, userHeight)
}

func getUserName() string {
	var userName string

	fmt.Print("Enter yoar name:")

	if _, err := fmt.Scan(&userName); err != nil {
		panic(err)
	}

	// Задача: якщо є пробіл то обрізати його , якщо літера з маленької літери то підняти її.
	var result strings.Builder

	userName = strings.ToLower(userName)
	userName = strings.TrimSpace(userName)

	var firstTime bool // по defolt завжди true
	var aLatterIndex []int
	for i, letter := range userName { //[     M + asdfasdfasdf]
		if letter == ' ' {
			// Пропускаємо пробіл
			continue

		} else if unicode.IsLower(letter) && !firstTime { //true - false
			// Перетворення маленьких літер на великі
			result.WriteRune(unicode.ToUpper(letter))
			firstTime = true
		} else {
			result.WriteRune(letter)
		}

		if userName[i] == 'a' {
			aLatterIndex = append(aLatterIndex, i)
		}
	}

	fmt.Println(aLatterIndex)
	return result.String()
}

func getUserAge() int {

	// var userAge1 string

	// fmt.Println( "Enter your age: ")

	// fmt.Scan(&userAge1)

	// ageUint, err := strconv.ParseUint(userAge1,10,0)
	// if err != nil{
	// 	log.Fatal ( "Incorrect value")
	// }

	// if ageUint <=10{
	// 	log.Fatal( "Age must be at least 10")
	// }
	// fmt.Println("Age - % d\n", ageUint)

	//задекларувати строкову змінну (вік) *
	//отримати занчення від користувача *
	//сконвертувати занчення в uint
	//Перевірити що вік був більше 18, якщо меньше вивсети log.Fatal("- -")
	//вивести форматовану строку, де буде "Вік - 50"

	//2. Так само але дня float значень на прикладі зросту людини
	var userAge string

	fmt.Print("Enter yoar age:")

	if _, err := fmt.Scan(&userAge); err != nil {
		panic(err)
	}

	userAgeInt, err := strconv.Atoi(userAge)
	if err != nil {
		log.Fatal(err)
	}

	if userAgeInt <= 18 {
		log.Fatal(`Age should be more then 18 years`)
	}

	if userAgeInt >= 110 {
		log.Fatal("Your age is too great, it is dangerous.")
	}
	return userAgeInt

}

//2. Так само але дня float значень на прикладі зросту людин
// Знайти в масиві кому і замінити її на крапку

func getUserHeight() float64 {

	var userHeight string

	fmt.Println("Enter yoar height:")

	if _, err := fmt.Scan(&userHeight); err != nil {
		panic(err)

	}
	userHeightFloat, err := strconv.ParseFloat(userHeight, 64)
	if err != nil {
		log.Fatal(err)
	}

	if userHeightFloat <= 130.1 {
		log.Fatal(" Your height is less than 131 cm. Sorry, you do not duality")
	}

	if userHeightFloat >= 221.1 {
		log.Fatal(" Your height is more than 220 cm. Sorry, you do not duality")
	}

	return userHeightFloat

}
