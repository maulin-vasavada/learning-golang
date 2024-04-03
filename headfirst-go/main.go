package main

import (
	"headfirst-go/maps"
)

func main() {
	/*
		fmt.Println("Hello from main")
		greeting.Greetings()
		//guessinggame.PlayGame()
		hello.SayHello()
		array.HelloArrays()
		array.Print()
		numbers, err := files.GetFloats("data.txt")
		if err != nil {
			log.Fatal(err)
		}

		numbers := make([]float64, 0)
		arguments := os.Args[1:]
		for _, argument := range arguments {
			number, err := strconv.ParseFloat(argument, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)
		}
		array.Average(numbers...)
	*/
	maps.Test()
	maps.VoteCount()
}
