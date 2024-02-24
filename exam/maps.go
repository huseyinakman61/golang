package main

import "fmt"

func main() {

	//maps - exam1
	var phones map[string]int
	phones = make(map[string]int, 0)

	phones["iphone15"] = 60000
	phones["iphone15pro"] = 80000
	phones["iphone15promax"] = 100000

	fmt.Println("phones", phones)
	fmt.Println("iphone15", phones["iphone15"])
	fmt.Println("none value", phones["iphone16"])

	//maps - exam2
	phoneDetail := map[string]int{
		"256gb": 10000,
		"512gb": 20000,
		"1tb":   30000,
		"2tb":   40000,
	}
	fmt.Println("phoneDetail", phoneDetail)

	//maps -> delete
	delete(phoneDetail, "2tb")
	fmt.Println("phoneDetail", phoneDetail)

}
