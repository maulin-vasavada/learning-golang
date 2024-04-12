package main

import (
	"fmt"
	"headfirst-go/prose"
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
	*/

	/*
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

	/*
		maps.Test()
		maps.VoteCount()
		maps.OrderTest()
	*/

	/*
		structs.Test()
	*/

	/*
		mytypes.Test()
		mytypes.LegalButIncorrectConversion()
		mytypes.CorrectConversion()
		mytypes.ConversionWithFuncs()
	*/

	/*
		date := calendar.Date{}
		err := date.SetYear(2019)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date.Year())

		err = date.SetMonth(5)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date.Month())

		err = date.SetDay(10)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date.Day())
	*/

	/*
		var toy interfaces.NoiseMaker
		toy = interfaces.Honk("Toyco Blaster")
		toy.MakeSound()
		toy = interfaces.Whistle("Toyco Camry")
		toy.MakeSound()
		interfaces.Play(toy)
	*/

	/*
		channels.Test()
		urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
		channels.FetchHttpResponseSizes(urls)
	*/

	fmt.Println(prose.JoinWithSeparator([]string{"apple", "orange", "pear", "banana"}, ","))
	fmt.Println(prose.JoinWithSeparator([]string{"apple", "orange"}, ","))
}
