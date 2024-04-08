package structs

import "fmt"

type Subscriber struct {
	name   string
	rate   float64
	Active bool
	Address
}

type Address struct {
	Street, City, State, PostalCode string
}

type car struct {
	name     string
	topSpeed float64
}

func printSubscriber(s Subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Active:", s.Active)
	fmt.Println("Street:", s.Street)
	fmt.Println("City:", s.City)
	fmt.Println("State:", s.State)
	fmt.Println("PostalCode:", s.PostalCode)

}

func applyDiscount(s *Subscriber) {
	s.rate = 0.5 * s.rate
}

func Test() {
	var subscriber1 Subscriber
	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.Active = true
	subscriber1.Street = "221B Baker St"
	subscriber1.City = "London"
	subscriber1.State = "Secret"
	subscriber1.PostalCode = "Secret"

	var subscriber2 Subscriber
	subscriber2.name = "Chaman Singh"
	subscriber2.rate = 8.23
	subscriber2.Active = true

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
