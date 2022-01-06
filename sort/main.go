package main

import (
	"fmt"
	"sort"
)

type User struct {
	name  string
	point int
}

type Users []User

var users = []User{
	{name: "user1", point: 1200},
	{name: "user2", point: 1380},
	{name: "user3", point: 1420},
	{name: "user4", point: 1150},
	{name: "user5", point: 1020},
}

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].point < u[j].point
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	fmt.Println(users)
	sort.Sort(Users(users))
	fmt.Println(users)
}
