package t12

import (
	"fmt"
	"strings"
)

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func MakeSetArray() {
	//Массивы comparable, но способ совершенно негибкий, мы привязаны к размеру
	s := [5]string{"cat", "cat", "dog", "cat", "tree"}

	sets := make(map[[5]string]struct{})

	sets[s] = struct{}{}

	fmt.Println(sets)
}

type SeqSet map[string]struct{}

func New() SeqSet {
	return SeqSet{}
}

func (s SeqSet) String() string {
	size := len(s) + 2 //выделим еще слота под скобочки
	strs := make([]string, size)
	strs[0], strs[size-1] = "(", ")"
	ind := 1
	for k, _ := range s {
		strs[ind] = k
		ind++
	}
	return strings.Join(strs, "")
}

func (s SeqSet) AddStrings(strs ...string) bool {
	if len(s) == 1 {
		//гарантируем, что множество состоит
		//максимум из одного элемента
		return false
	}
	return s.add(strs)
}

func (s SeqSet) add(strs []string) bool {
	//соединим всю переданную нам последовательность
	//в одну строку, чтобы не зависеть от длины последовательности
	str := strings.Join(strs, ",")
	if _, ok := s[str]; !ok {
		s[str] = struct{}{}
		return true
	}
	return false
}
