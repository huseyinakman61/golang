package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	var age int
	var isMarried bool
	var price float32

	name = "Hüseyin Akman"
	age = 36
	isMarried = true
	price = 1250.25

	fmt.Println("Name:", name)
	fmt.Printf("%T\n", name)
	fmt.Println("Name:", reflect.TypeOf(name))
	fmt.Println("Age:", age)
	fmt.Printf("%T\n", age)
	fmt.Println("Age:", reflect.TypeOf(age))
	fmt.Printf("%T\n", isMarried)
	fmt.Println("Married:", reflect.TypeOf(isMarried))

	fmt.Println("price:", price)
}
