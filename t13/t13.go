package t13

import "fmt"

//Поменять местами два числа без создания временной переменной.

func SwapValuesMath() {
	//математический способ
	a, b := 5, 2

	a = a + b //a = 5 + 2 = 7
	b = a - b //b = 7 - 2 = 5 (!)
	a = a - b //a = 7 - 5 = 2 (!)
	fmt.Println(a, b)
}
func SwapValuesXOR() {
	//битовые манипуляции через XOR
	a, b := 5, 2

	a = a ^ b //101 ^ 010 -> 111 -> 7
	b = a ^ b //010 ^ 111 -> 101 -> 5 (!)
	a = a ^ b //111 ^ 101 -> 010 -> 2 (!)
	fmt.Println(a, b)
}
