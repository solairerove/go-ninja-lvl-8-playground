package jedi1

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

// Ex1 tbd
func Ex1() {
	fmt.Println("\nEx1")

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
