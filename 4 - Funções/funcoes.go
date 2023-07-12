package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)

		return txt
	}

	result := f("testando")
	fmt.Println(result)

	resultSoma, resultSub := calculosMatematicos(10, 25)
	fmt.Println(resultSoma, resultSub)

	teste1, _ := calculosMatematicos(10, 25)
	fmt.Println(teste1)

}
