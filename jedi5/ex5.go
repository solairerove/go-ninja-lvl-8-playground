package jedi5

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u *user) printUser() {
	fmt.Println("\n", u.First, u.Last, u.Age)
	for _, s := range u.Sayings {
		fmt.Println(s)
	}
}

func printUsers(users []user) {
	for _, u := range users {
		u.printUser()
	}
}

type sortByAge []user

func (a sortByAge) Len() int           { return len(a) }
func (a sortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type sortByLast []user

func (l sortByLast) Len() int           { return len(l) }
func (l sortByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l sortByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

// Ex5 tbd
func Ex5() {
	fmt.Println("\nEx5")

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	printUsers(users)

	fmt.Println("\nsortByAge")
	sort.Sort(sortByAge(users))
	printUsers(users)

	fmt.Println("\nsortByLast")
	sort.Sort(sortByLast(users))
	printUsers(users)
}
