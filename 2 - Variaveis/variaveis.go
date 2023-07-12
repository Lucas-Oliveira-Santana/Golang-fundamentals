package main

import "fmt"

func main() {
	var variavel1 string = "var 1"
	variavel2 := "var 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var 3"
		variavel4 string = "var 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var 5", "var 6"

	fmt.Println(variavel5, variavel6)

	const constate1 string = "const 1"

	fmt.Println(constate1)
}
