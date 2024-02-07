package main

import (
	"fmt"
	"math"
)

// Formula de DELTA
//  d = b ^ 2 - 4 x a x c
/*
Se Δ > 0, a equação possui duas soluções reais.

Se Δ = 0, a equação possui uma solução real.

Se Δ < 0, a equação não possui solução real.

*/
func calcularDelta(a int, b int, c int) int {

	d := int(math.Pow(float64(b), 2)) - (4 * a * c)
	return d
}

func main() {
	var soma int
	valores := [3]int{1, 2, 3}
	for _, val := range valores {
		soma += val
	}
	var delta = calcularDelta(valores[0], valores[1], valores[2])

	if delta > 0 {
		fmt.Println("Possui duas raizes")
	} else if delta == 0 {
		fmt.Println("Possui uma raiz")
	} else {
		fmt.Println("Não possui uma solução REAL")
	}

}
