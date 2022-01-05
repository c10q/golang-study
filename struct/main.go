package main

import "fmt"

type skill struct {
	name     string
	damage   int
	manaCost int
	coolTime int
}

type champion struct {
	name string
	q    skill
	w    skill
	e    skill
	r    skill
}

func main() {
	amumu := champion{
		name: "아무무",
		q: skill{
			name:     "붕대날리기",
			damage:   50,
			manaCost: 30,
			coolTime: 9,
		},
		w: skill{
			name:     "울기",
			damage:   4,
			manaCost: 5,
			coolTime: 8,
		},
		e: skill{
			name:     "긁기",
			damage:   30,
			manaCost: 10,
			coolTime: 4,
		},
		r: skill{
			name:     "궁극기",
			damage:   140,
			manaCost: 100,
			coolTime: 50,
		},
	}

	fmt.Println(amumu)
	fmt.Println(amumu.q.name)

}
