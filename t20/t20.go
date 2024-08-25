package t20

import "strings"

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

//Первое, что приходит на ум - strings.Split и метод двух указателей
//В таком случае нам не стоит беспокоиться о кодировке и кропотливо
//работать со строками. Просто разобьем на слайсы и двумя указателями
//поменяем слова местами, вернем результат через Join.

func ReverseWords(str string) string {
	splited := strings.Split(str, " ")
	for i, j := 0, len(splited)-1; i < j; i, j = i+1, j-1 {
		splited[i], splited[j] = splited[j], splited[i]
	}
	return strings.Join(splited, " ")
}
