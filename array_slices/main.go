package main

import "fmt"

func main() {

	//array
	var country [3]string
	country[0] = "Turkey"
	country[1] = "Germany"
	country[2] = "USA"

	var colors = [6]string{"Blue", "Red", "Green", "Yellow", "Black", "White"}

	//slices
	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println("country 1:", country[0])
	fmt.Println("country 2:", country[1])
	fmt.Println("country 2:", country[2])
	fmt.Println("country 2:", colors)
	fmt.Println("country 2:", colors[3:5])
	fmt.Println("country 2:", days)
}
