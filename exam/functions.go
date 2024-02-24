package main

import (
	"fmt"
	"strings"
)

func main() {
	fullname("Hüseyin", "Akman") //no return

	fmt.Println("status: ", checkId(12)) //return bool

	name, surname := multiResponse("Hüseyin", "Akman") //return multiple
	fmt.Println("Multi Response: ", name, surname)

	var num = []int{3, 5, 10, 7, 4, 8}
	fmt.Println("Calc: ", calc(num)) //slices(array) --> return int

	fmt.Println("Sum 1: ", sum(3, 5, 10, 7, 4, 8)) //multiple values --> return int
	fmt.Println("Sum 2: ", sum(11, 4, 5))          //multiple values --> return int
	fmt.Println("Sum 3: ", sum(16, 22, 33, 8))     //multiple values --> return int

	fmt.Println("Multi Values Any Response: ", multiValues(34, "Hüseyin", "Akman", 61)) //multiple type values --> return any

}

// function -> no return
func fullname(name string, surname string) {
	fmt.Println(name + " " + surname)
}

// function -> return bool(true/false)
func checkId(id int) bool {
	fmt.Println("id:", id)
	if id == 10 {
		return true
	} else {
		return false
	}
}

// function -> return multiple (string, string)
func multiResponse(name string, surname string) (string, string) {
	return strings.ToUpper(name), strings.ToUpper(surname)
}

// function(slices(array)) --> return int
func calc(num []int) int {
	count := 0
	for _, value := range num {
		count += value
	}
	return count
}

// function multiple values(1,5,8,4,15,n..) --> return int
func sum(num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}

// function multiple values(1,5,8,4,15,n..) --> return int
func multiValues(num ...any) any {
	for _, value := range num {
		fmt.Println("value: ", value)
	}
	return "Test"
}
