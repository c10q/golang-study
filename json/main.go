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
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
