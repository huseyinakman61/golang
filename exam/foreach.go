package main

import "fmt"

func main() {

	var colors = [6]string{"Blue", "Red", "Green", "Yellow", "Black", "White"}

	fmt.Println("colors", colors)

	//for each -> index + value
	for index, value := range colors {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
		//fmt.Println("color :", colors[index])
	}

	fmt.Println("-------------------------------")

	//for each -> only value
	for _, value := range colors {
		fmt.Println("value: ", value)
	}

	fmt.Println("-------------------------------")

	//for each -> only index
	for index := range colors {
		//fmt.Println("index", index)
		fmt.Println("color:", colors[index])
	}

	fmt.Println("-------------------------------")

	//for each -> string
	var name string = "Hüseyin Akman"
	for _, char := range name {
		//fmt.Println("index", index)
		fmt.Printf("\n%c", char)
	}

	fmt.Println("-------------------------------")

	//for each -> maps
	phones := map[string]int{
		"iphone12": 10000,
		"iphone13": 20000,
		"iphone14": 30000,
		"iphone15": 40000,
	}
	for key, value := range phones {
		fmt.Printf("Key: %s, Value: %d \n", key, value)
	}
}
