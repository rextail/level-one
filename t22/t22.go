package t22

import "math/big"

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a, b, значение которых > 2^20.

//Хоть 2^20 и не является каким-то сверхъестественным числом, но т.к. это нижний предел,
//возможно, стоит задуматься о том, что числа могут быть куда больше.
//Для работы с числами, длина которых превышает машинное слово, команда Rob Pike and Co.
//Разработала в рамках пакета math пакет big, реализующий длинную арифметику.

func AddFloats(a, b *big.Float) *big.Float {
	c := new(big.Float).Add(a, b)
	return c
}
func SubFloats(a, b *big.Float) *big.Float {
	c := new(big.Float).Sub(a, b)
	return c
}
func DivFloats(a, b *big.Float) *big.Float {
	c := new(big.Float).Quo(a, b)
	return c
}
func MulFloats(a, b *big.Float) *big.Float {
	c := new(big.Float).Mul(a, b)
	return c
}

func AddInts(a, b *big.Int) *big.Int {
	c := new(big.Int).Add(a, b)
	return c
}
func SubInts(a, b *big.Int) *big.Int {
	c := new(big.Int).Sub(a, b)
	return c
}
func DivInts(a, b *big.Int) *big.Int {
	c := new(big.Int).Quo(a, b)
	return c
}
func MulInts(a, b *big.Int) *big.Int {
	c := new(big.Int).Mul(a, b)
	return c
}

func AddRats(a, b *big.Rat) *big.Rat {
	c := new(big.Rat).Add(a, b)
	return c
}
func SubRats(a, b *big.Rat) *big.Rat {
	c := new(big.Rat).Sub(a, b)
	return c
}
func DivRats(a, b *big.Rat) *big.Rat {
	c := new(big.Rat).Quo(a, b)
	return c
}
func MulRats(a, b *big.Rat) *big.Rat {
	c := new(big.Rat).Mul(a, b)
	return c
}
