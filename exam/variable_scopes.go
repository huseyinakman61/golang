package main

import "fmt"

var globalValue int = 10   //global değer
var globalValue2 int = 50  //global değer 2
var globalValue3 int = 150 //global değer 3

//burada var kullanımı çok önemli: eğer global değeri değiştirmek istiyorsan var kullanma.
//eğer tanımlı değişkeni anlık kullanmak istiyorsan function, if vb. yerlerde var kullanarak global değişkeni değiştirMEden işlem yapabilirsin
func main() {
	//global değişkene işlem yapıyoruz
	fmt.Println("global 1", globalValue)  //global değeri yazdırıyoruz
	fmt.Println("global 2", globalValue2) //global2 değeri yazdırıyoruz
	fmt.Println("global 3", globalValue3) //global3 değeri yazdırıyoruz
	var globalValue2 = 100                //globalValue2 değerini var ile function içinde değiştiriyoruz
	changeValue()                         //global değeri function içinde değiştirip yazdırıyoruz
	fmt.Println("global 1", globalValue)  //global değeri her yerde değişiyor
	fmt.Println("global 2", globalValue2) //global değeri bu function içinde kullanılan her yerde değişiyor
	fmt.Println("global 3", globalValue3) //global değeri changeValue içinde var ile değiştirdiğimiz için burada global değeri değiştirmiyor
}

func changeValue() {
	globalValue = 20
	var globalValue3 = 500
	fmt.Println("global 1", globalValue)
	fmt.Println("global 2", globalValue2)
	fmt.Println("global 3", globalValue3)
}
