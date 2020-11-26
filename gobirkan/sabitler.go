package main

//go enum yapamak -verileri paket haline getirip erismek
import (
	"fmt"
)

type Hey string

const (
	thy Hey = "Turkish Airlines"
	tai Hey = "Tai"
)

func PrintHey(hey Hey) {
	fmt.Println(hey)
}

func main() {
	PrintHey(thy)
	PrintHey(tai)

}
