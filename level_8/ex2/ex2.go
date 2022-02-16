package ex2

import (
	"encoding/json"
	"fmt"
)

var s = `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

// Create a struct type of person followed the JSON structure
type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Using json.Marshal to convert a JSON to a struct
func Ex2() {

	fmt.Println(s)

	// First, We need to convert the JSON to a byte slice before using Unmarshal
	bs := []byte(s)

	// Next, we need a slice of type person
	people := []person{}

	// Use Unmarshal to add the byte slice of s to people
	// Specially, we need to reference to people pointer.
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	fmt.Println()

	// Loop through people, a slice of type person
	for i, person := range people {
		fmt.Print("Person number ", i)
		fmt.Println(" with identity of", person.First, person.Last, person.Age, "has said:")
		for _, say := range person.Sayings {
			fmt.Println("\t", say)
		}
	}
}
