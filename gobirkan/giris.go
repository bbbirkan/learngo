package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// konsol girisi
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Yazini Yaz: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Numarani yaz: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sayi degeri:", f)
	}
}
