package main

import "fmt"

func main() {
	var number int = 22
	number2 := number
	var point *int
	point = &number
	fmt.Println("number:", number)
	fmt.Println("number2=number:", number2) //number2 değeri number ile aynı değere sahip
	fmt.Println("point:", point)
	fmt.Println("point in value:", *point)

	fmt.Println("-------------------------------")

	*point = 30 //change pointer in value(number)

	fmt.Println("number:", number)
	fmt.Println("number2=number:", number2) //number2 değişmiyor. çünkü bellekten sadece number değişkeni değiştirildi
	fmt.Println("point:", point)            //ram point value: not change
	fmt.Println("point in value:", *point)

	fmt.Println("-------------------------------")

	//number3 değerini fonksiyon içinde değiştirmek için pointer kullanıyoruz.
	//normalde global değişken tanımlayıp onu her yerden değiştirmek daha mantıklı bir çözüm fakat bu dilde yok :)
	//pointer ile ramdeki değerini değiştiriyoruz ve bu sayede fonksiyon içinde number3 değeri değişmiş oluyor.
	//bu saçma yapıyı unutmamam lazım bana çok lazım olacak
	var number3 int = 10
	fmt.Println("number3:", number3)
	changePublicVal(&number3)
	fmt.Println("number3:", number3)

	fmt.Println("-------------------------------")
	//buradaa pointer yok. test için eklendi
	//neyseki bu saçma sistem array ve slices değerlerinde yok :)
	var numbers = []int{12, 415, 51, 3}
	fmt.Println("numbers:", numbers)
	changeArr(numbers)               //fakat func içinde array değerini değiştiriiyor
	fmt.Println("numbers:", numbers) //çıktıda değişken içindeki değerde değişiyors
}

//bu fonksiyon ile gelen değerin(pointer değeri) global değerini değiştiriyoruz.
func changePublicVal(i *int) {
	*i = *i + 10
}

//pointer kullanmadan array değerleri fonksiyon içinde değişebiliyor.
func changeArr(i []int) {
	i[0] = 250
}
