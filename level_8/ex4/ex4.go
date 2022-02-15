package ex4

import (
	"fmt"
	"sort"
)

func sliceSort() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi  INLINE
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs 	INLINE
	sort.Strings(xs)
	fmt.Println(xs)
}

func Ex4() {
	sliceSort()
}
