package main

import "fmt"

func main() {
	//konsol :veri cikisi

	yazi1 := "Birkan Kalyon"
	yazi2 := "Goland ogreniyor"
	yazi3 := "Ama pek neseli degil :)"
	sayi1 := 3
	secenek1 := true

	stringLength, _ := fmt.Println(yazi1, yazi2, yazi3, sayi1, secenek1)
	fmt.Println("Yazi Uzunlugu : ", stringLength)

	deneme := fmt.Sprintf("Konsol cikislari %s", "2")
	fmt.Println(deneme)

	fmt.Printf("Datanin Turu:\n yazi 1: %T\n yazi 2: %T\n yazi 3: %T\n and sayi1: %T\n secenek1: %T \n", yazi1, yazi2, yazi3, sayi1, secenek1)
	hey := fmt.Sprintf("Datanin Turu:%T, %T, %T, %T, ve %T", yazi1, yazi2, yazi3, sayi1, secenek1)
	fmt.Println(hey)

}
