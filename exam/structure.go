package main

import "fmt"

var customer2 = Customer{id: 2, name: "Berat Akman", age: 14}

func main() {

	var customer1 = Customer{id: 1, name: "Hüseyin Akman", age: 36, address: Address{country: "Turkey", city: "İstanbul", postalCode: 34150}} //değerleri toplu tanımlıyoruz
	//değerleri toplu tanımlıyoruz

	fmt.Println("customer1", customer1)
	customer1.age = 37 //tek değeri manuel değiştiriyoruz
	fmt.Println("customer1", customer1)
	customer1.toString() //Customers objesinin içindeki ToString fonksiyonu çağırıyoruz
	fmt.Println("customer2", customer2)
	customer2.changeName("Burak Akman")
	fmt.Println("customer2", &customer2)
}

type Customer struct { //Customer adında obje tanımlıyoruz
	id      int32
	name    string
	age     int
	address Address //Address obje değerlerini Customer içine ekliyoruz
}

type Address struct { //Address adında obje tanımlıyoruz
	country    string
	city       string
	postalCode int
}

func (customer Customer) toString() { //Customers objesinin içine ToString fonksiyonu ekliyoruz
	fmt.Printf("Name: %s, Age: %d\n", customer.name, customer.age)
}

func (customer *Customer) changeName(name string) { //Customers içindeki değeri pointer ile değiştiriyoruz
	//customer2.name = name
	customer.name = name
}
