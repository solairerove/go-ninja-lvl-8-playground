package jedi2

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

// Ex2 tbd
func Ex2() {
	fmt.Println("\nEx2")

	s := `[
		{
			"First": "James",
			"Last": "Bond",
			"Age": 32,
			"Sayings": [
				"Shaken, not stirred",
				"Youth is no guarantee of innovation",
				"In his majesty's royal service"
			]
		},
		{
			"First": "Miss",
			"Last": "Moneypenny",
			"Age": 27,
			"Sayings": [
				"James, it is soo good to see you",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."
			]
		},
		{
			"First": "M",
			"Last": "Hmmmm",
			"Age": 54,
			"Sayings": [
				"Oh, James. You didn't.",
				"Dear God, what has James done now?",
				"Can someone please tell me where James Bond is?"
			]
		}
	]`

	var people []person

	json.Unmarshal([]byte(s), &people)

	for _, p := range people {
		fmt.Println("\n", p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Println(s)
		}
	}
}
