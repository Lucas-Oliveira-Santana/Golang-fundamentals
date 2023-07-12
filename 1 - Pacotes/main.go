package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Printf("Ecrevendo no main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("lucas@gmail.com")
	fmt.Println(erro)
}
