package common

import "strings"

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Intersect[T comparable](a, b []T) []T {
	var set = make([]T, 0)
	var hashTable = make(map[T]struct{})

	for _, val := range a {
		hashTable[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := hashTable[val]; ok {
			set = append(set, val)
		}
	}

	return set
}

func IntersectStrings(a, b string) string {
	var result = ""
	var hashTable = make(map[rune]struct{})

	for _, val := range a {
		hashTable[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := hashTable[val]; ok && !strings.ContainsRune(result, val) {
			result += string(val)
		}
	}

	return result
}

func Last[T any](a []T) T {
	return a[len(a)-1]
}
