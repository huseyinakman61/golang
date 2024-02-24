package main

import "fmt"

func main() {

	//default for
	for i := 1; i < 10; i++ {
		fmt.Println("i", i)
	}

	fmt.Println("-------------------------------")

	//inline for
	count := 1
	for count <= 10 {
		fmt.Println("count", count)
		count++
	}

	fmt.Println("-------------------------------")

	//array for
	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for a := 0; a < len(days); a++ {
		fmt.Println(a+1, days[a])
	}

	fmt.Println("-------------------------------")

	//break for
	for i := 1; i < 10; i++ {
		if i == 5 {
			fmt.Println("bitir", i)
			break
		}
		fmt.Println("break: ", i)
	}

	fmt.Println("-------------------------------")

	//continue for
	for i := 1; i <= 10; i++ {
		if i == 5 || i == 10 {
			fmt.Println("atla", i)
			continue
		}
		fmt.Println("continue: ", i)
	}

}
