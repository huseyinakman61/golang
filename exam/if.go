package main

import "fmt"

func main() {
	const age int = 19 //const
	//age = 22 //const değeri değiştirilemez
	if age == 18 {
		fmt.Println("Yaşınız 18")
	} else if age >= 18 {
		fmt.Println("Yaşınız 18den büyük")
	} else {
		fmt.Println("Yaşınız 18den küçük")
	}
}
