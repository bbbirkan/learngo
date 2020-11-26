package main

import (
	"fmt"
)

func main() {

	//Convert
	/*
		var yazi = "1"
		var x = 10
		var f float32 = 2.0

		fmt.Println(yazi, x, f)
		// stringten integire donusturme

		sayi, _ := strconv.Atoi(yazi)
		fmt.Println(sayi)
	*/
	var i int = 55
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
}
