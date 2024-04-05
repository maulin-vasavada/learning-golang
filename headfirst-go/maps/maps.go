package maps

import (
	"fmt"
	"headfirst-go/files"
	"log"
)

func Test() {
	ranks := map[string]int{"red": 10, "blue": 20}
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["red"])
	fmt.Println(ranks["violet"])
	count, ok := ranks["violet"]
	fmt.Println("Violet count and ok are ", count, ok)
}

func VoteCount() {
	lines, err := files.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for candidate, votes := range counts {
		fmt.Println(candidate, "got", votes, "votes")
	}
}

func OrderTest() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	for student, grade := range grades {
		fmt.Println(student, "'s grade is", grade)
	}
}
