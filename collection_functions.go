package main

import (
	. "fmt"
	"strings"
)

func Index(strs []string, target string) int {
	for i, str := range strs {
		if str == target {
			return i
		}
	}

	return -1
}

func Include(strs []string, target string) bool {
	return Index(strs, target) >= 0
}

func Any(strs []string, f func(string) bool) bool {
	for _, str := range strs {
		if f(str) {
			return true
		}
	}
	return false
}

func All(strs []string, f func(string) bool) bool {
	for _, str := range strs {
		if !f(str) {
			return false
		}
	}
	return true
}

func Filter(strs []string, f func(string) bool) []string {
	ret := make([]string, 0)
	for _, str := range strs {
		if f(str) {
			ret = append(ret, str)
		}
	}
	return ret
}

func Map(strs []string, f func(string) string) []string {
	ret := make([]string, len(strs))
	for i, str := range strs {
		ret[i] = f(str)
	}
	return ret
}

func main() {
	strs := []string{"hello", "world", "and", "johnny"}
	Println("Index:", Index(strs, "world"))
	Println("Include:", Include(strs, "world"))

	Println("Any:", Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "w")
	}))

	Println("All:", All(strs, func(s string) bool {
		return strings.HasPrefix(s, "j")
	}))

	Println("Filter:", Filter(strs, func(s string) bool {
		return strings.Contains(s, "o")
	}))

	Println("Map:", Map(strs, strings.ToUpper))
}
