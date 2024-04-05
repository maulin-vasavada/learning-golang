package structs

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type car struct {
	name     string
	topSpeed float64
}

func printSubscriber(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Active:", s.active)

}

func applyDiscount(s *subscriber) {
	s.rate = 0.5 * s.rate
}

func Test() {
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.active = true

	var subscriber2 subscriber
	subscriber2.name = "Chaman Singh"
	subscriber2.rate = 8.23
	subscriber2.active = true

	printSubscriber(subscriber1)
	applyDiscount(&subscriber1)
	printSubscriber(subscriber1)

	printSubscriber(subscriber2)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Car name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)

	myCar := car{name: "Corvette", topSpeed: 337}
	fmt.Println("Car name:", myCar.name)
	fmt.Println("Top Speed:", myCar.topSpeed)
}
