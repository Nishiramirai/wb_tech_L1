package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
//  две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий: в Go тип int справится с такими числами, но обратите внимание
// на возможное переполнение для ещё больших значений.
// Для очень больших чисел можно использовать math/big.

func main() {
	a := new(big.Int)
	a.SetString("20000000000000000000", 10)

	b := new(big.Int)
	b.SetString("50000000000000000000", 10)

	sum := new(big.Int)
	diff := new(big.Int)
	product := new(big.Int)
	quotient := new(big.Int)

	sum.Add(a, b)
	diff.Sub(a, b)
	product.Mul(a, b)

	if b.Sign() != 0 {
		quotient.Div(a, b)
	}

	fmt.Printf("%s + %s = %s\n", a, b, sum)
	fmt.Printf("%s - %s = %s\n", a, b, diff)
	fmt.Printf("%s * %s = %s\n", a, b, product)
	fmt.Printf("%s / %s = %s\n", a, b, quotient)

}
