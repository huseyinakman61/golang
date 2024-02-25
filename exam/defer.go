package main

import "fmt"

//defer içinde bulunduğu alandaki(fonksiyon içindeki) tüm tüm işlemler tamamlandıktan sonra çalışır.
//önemli not: defer -> tersten çalışır. en üstteki defer en son, en sondaki defer ilk çalışır
func main() {
	fmt.Println("Hello")
	fmt.Println(" Wolrd !")

	fmt.Println("-------------------")

	defer fmt.Println("Hello ")
	fmt.Println(" Wolrd !")

	deferWrite()
}

func deferWrite() {
	fmt.Println("-------------------")
	//aşağıda defer değerlerinin tersten yazdığını görebiliriz
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")
}
