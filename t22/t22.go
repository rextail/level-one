package t22

import "math/big"

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a, b, значение которых > 2^20.

//Хоть 2^20 и не является каким-то сверхъестественным числом, но т.к. это нижний предел,
//возможно, стоит задуматься о том, что числа могут быть куда больше.
//Для работы с числами, длина которых превышает машинное слово, команда Rob Pike and Co.
//Разработала в рамках пакета math пакет big, реализующий длинную арифметику.

//Пример:
//unsigned int u = 1e9;
//long signed int i = -1;
//i + u???

type Test interface {
	~int | ~float64 | ~uint
}

func Add[S Test](a, b S) S {
	return a + b
}

func AddFloats(a, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}
func SubFloats(a, b *big.Float) *big.Float {
	return new(big.Float).Sub(a, b)
}
func DivFloats(a, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}
func MulFloats(a, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

func AddInts(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}
func SubInts(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}
func DivInts(a, b *big.Int) *big.Int {
	return new(big.Int).Quo(a, b)
}
func MulInts(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func AddRats(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Add(a, b)
}
func SubRats(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Sub(a, b)
}
func DivRats(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Quo(a, b)
}
func MulRats(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Mul(a, b)
}
