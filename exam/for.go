package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println("i", i)
	}

	count := 1
	for count <= 10 {
		fmt.Println("count", count)
		count++
	}

}
