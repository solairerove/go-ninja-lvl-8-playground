package jedi4

import (
	"fmt"
	"sort"
)

// Ex4 tbd
func Ex4() {
	fmt.Println("\nEx4")

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{
		"random",
		"rainbow",
		"delights",
		"in",
		"torpedo",
		"summers",
		"under",
		"gallantry",
		"fragmented",
		"moons",
		"across",
		"magenta",
	}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
