package main

import (
	"fmt"
)

func GetUserDay() int {

	var getUserDay int

	fmt.Println("What day of the week is it today?")

	fmt.Scan(&getUserDay)

	userDay := getUserDay

	switch userDay {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")

	default:
		fmt.Println("Invalid day number")
	}
return getUserDay
}
