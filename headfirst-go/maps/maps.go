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
	fmt.Println(counts)
}
