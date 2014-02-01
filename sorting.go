package main

import (
	. "fmt"
	"sort"
)

type ReverseSort []int

func (r ReverseSort) Len() int {
	return len(r)
}

func (r ReverseSort) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r ReverseSort) Less(i, j int) bool {
	//return r[i] < r[j]
	if r[i] < r[j] {
		return false
	} else {
		return true
	}
}

func main() {
	// Simple sort
	nums := []int{6, 2, 7, 3, 10, 22, 9}
	sort.Ints(nums)
	Println("Sorted:", sort.IntsAreSorted(nums))
	Println(nums)

	strs := []string{"hello", "world", "and", "johnny"}
	sort.Strings(strs)
	Println("Sorted:", sort.StringsAreSorted(strs))
	Println(strs)

	// Custom function sort
	newNums := []int{6, 2, 7, 3, 10, 22, 9}
	sort.Sort(ReverseSort(newNums))
	Println(newNums)
}
