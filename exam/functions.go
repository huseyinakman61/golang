package main

import (
	"fmt"
	"strings"
)

func main() {
	fullname("Hüseyin", "Akman") //no return
	//fmt.Println("heloo")
	fmt.Println("status: ", checkId(12))               //return bool
	name, surname := multiResponse("Hüseyin", "Akman") //return multiple
	fmt.Println("Multi Response: ", name, surname)
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
