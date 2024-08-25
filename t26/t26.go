package t26

import (
	"unicode"
	"unicode/utf8"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

//Например:
//abcd — true
//abCdefAaf — false

//1.  хэш таблица гарантирует уникальность

//2.  при итерации по строке через range получаем int32 (rune)

//3.  struct{} не занимает память

//4.  не все символы кодируются одним байтом, соответственно
//	  размер мапы может раздуться, если того пожелает партия 中国,
//    так что можно использовать utf8.RuneCountInString
//    вместо len при создании мапы - пусть она лишний раз пройдется по строке,
//    но зато можно будет больше не бояться партия 中国

func IsUnique(str string) bool {
	used := make(map[int32]struct{}, utf8.RuneCountInString(str))
	for _, char := range str {
		c := unicode.ToLower(char)
		if _, ok := used[c]; ok {
			return false
		}
		used[c] = struct{}{}
	}
	return true
}
