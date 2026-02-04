package main

import (
	"fmt"
	"strings"
)

func main() {

	var day = "monday"
	isWDay(day)

	day = "WedNesdaY"
	isWDay(day)

	/*-----------------------------------------*/

	var users = map[string]int{
		"Vasya":  20,
		"Kolya":  21,
		"Vpiska": 23,
		"WHO!":   432,
	}
	// or
	// var users map[string]int
	// or
	// users := make(map[string]int){
	// 	"Vasya":  20,
	// 	"Kolya":  21,
	// 	"Vpiska": 23,
	// 	"WHO!":   432,
	// }

	fmt.Println(users)
	var sumOfAge int
	for _, age := range users {
		sumOfAge += age
	}
	fmt.Println(sumOfAge)

	age, exist := users["Vasya"]
	fmt.Println(age, exist)
	age, exist = users["Petr"]
	fmt.Println(age, exist)

	users["Misha"] = 32
	fmt.Println(users["Misha"])

	delete(users, "Petr")
	delete(users, "Misha")
	fmt.Println(users)
	fmt.Println(users)

}

func isWDay(day string) {

	day = strings.ToLower(day)

	switch day {
	case "wednesday":
		fmt.Println("its wednesday dude")
	default:
		fmt.Println("bro(")
	}
}
