package module02

import (
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.

// https://en.wikipedia.org/wiki/Insertion_sort#Algorithm
func InsertionSortInt(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

// the solutin from the course
func InsertionSortIntInsertFunc(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	copy(list, sorted)
}

func insert(sorted []int, item int) []int {
	// fmt.Println(item, sorted)
	for j, val := range sorted {
		if item < val {
			// fmt.Println(sorted[:j], sorted[j:])
			return append(sorted[:j], append([]int{item}, sorted[j:]...)...)
		}
	}
	return append(sorted, item)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
}
