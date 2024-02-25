package main

import "fmt"

func main() {
	//en üstte olsa bile fonksiyon sonlandığında çalışacaktır. panic ile program durdurulsa bile defer ile gelen kodlar çalışır. güzel bir özellik...
	//dipnot: elektrik kesintileri veya server down gibi durumlarda burası çalışmayabilir. buranın kurgusu doğru yapılmalıdır. bu kod her zaman 100% çalışacak gibi mantık kurulmamalıdır.
	defer closeDb()

	panic("Stop") //exit veya end komutu ile aynı işlemi görür uggulamayı durdurur

	defer fmt.Println("panic after defer") //not: yalnız panic ile kaçış uygulandıktan sonra sonra eklenen defer kodu çalışmaz.
}

func closeDb() {
	fmt.Println("Databsse Close")
}
