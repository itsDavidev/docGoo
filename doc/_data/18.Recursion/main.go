package main

import "fmt"

var Esponte int = 0

func main() {
	Expon(3, 3, 1)

}

// 2 2 = 4
// 2 3 = 8

func Expon(num int, exp int, i int) {
	if i == 1 {
		Esponte = num
	}
	if i == exp {
		return
	}

	var resultado = num * Esponte

	fmt.Println(num, " * ", Esponte, "=", resultado)

	Expon(resultado, exp, i+1)
}
