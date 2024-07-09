package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

func sum(a, b *big.Int) *big.Int {
	var res big.Int
	res.Add(a, b)
	return &res
}

func sub(a, b *big.Int) *big.Int {
	var res big.Int
	res.Sub(a, b)
	return &res
}

func mul(a, b *big.Int) *big.Int {
	var res big.Int
	res.Mul(a, b)
	return &res
}

func div(a, b *big.Int) *big.Int {
	var res big.Int
	res.Quo(a, b)
	return &res
}
func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("222222222222222222222222222222222222222", 10)
	b.SetString("111111111111111111111111111111111111111", 10)

	fmt.Printf("sum = %d\n", sum(a, b))
	fmt.Printf("sub = %d\n", sub(a, b))
	fmt.Printf("mul = %d\n", mul(a, b))
	fmt.Printf("div = %d\n", div(a, b))
}
