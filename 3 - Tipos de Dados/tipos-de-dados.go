package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int32 = 100000000
	var numero2 int = 100000000
	numero3 := 100000000
	fmt.Println(numero, numero2, numero3)

	var numero4 uint = 1000
	fmt.Println(numero4)

	// INT32 = rune
	var numero5 rune = 12345
	fmt.Println(numero5)

	// UINT8 = byte
	var numero6 byte = 123
	fmt.Println(numero6)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12345.67
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'a'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	var booleano3 bool
	fmt.Println(booleano3)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Error interno")
	fmt.Println(erro2)
}
