package t16

import (
	"cmp"
	"slices"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

//Можно написать самому, можно использовать стандартную гошную сортировку, внутри которой
//уже реализован quicksort.

//Само собой, разработчики Go не могут себе позволить обычный quicksort.
//Их алгоритм основан на pattern-defeating варианте алгоритма, который в нужный
//момент меняет стратегию сортировки, тем самым предотвращая худший случай O(n^2)

func QuickByGoTeam[C ~[]E, E cmp.Ordered](arr C) {
	slices.Sort(arr)
}

//В реализации ниже, если массив уже изначально отсортирован, то приходим к худшему времени работы O(n^2)
//Если же нет, то при удачном разбиении имеем O(nlog(n))

func Quicksort[C ~[]E, E cmp.Ordered](arr C, l int, r int) {
	if l < r {
		q := partition(arr, l, r)
		Quicksort(arr, l, q)
		Quicksort(arr, q+1, r)
	}
}

func partition[C ~[]E, E cmp.Ordered](arr C, l int, r int) int {
	v := arr[int(uint((l+r)>>1))]
	i := l
	j := r
	for i <= j {
		for arr[i] < v {
			i++
		}
		for arr[j] > v {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return j
}
