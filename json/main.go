package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "KNOH",
		Age:  25,
	}
	bytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(bytes)
	// [123 34 110 97 109 101 34 58 34 75 78 79 72 34 44 34 97 103 101 34 58 50 53 125]

	fmt.Println(string(bytes))
	/*{
		"name": "KNOH",
		"age": 25
	}*/

	var user1 User

	err = json.Unmarshal(bytes, &user1)
	if err != nil {
		return
	}

	fmt.Println(user1.Age)
	fmt.Println(user1.Name)

}
