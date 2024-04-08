package mytypes

import "fmt"

type Liters float64
type Gallons float64

func Test() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)
	fmt.Println(carFuel, busFuel)
}

func LegalButIncorrectConversion() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(Liters(40.0))
	busFuel = Liters(Gallons(63.0))
	fmt.Printf("Gallons: %0.1f, Liters %0.1f\n", carFuel, busFuel)
}

func CorrectConversion() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f, Liters %0.1f\n", carFuel, busFuel)
}

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MillilitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.000264)
}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMilliliters(g Gallons) Liters {
	return Liters(g * 3785.41)
}

func ConversionWithFuncs() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += LitersToGallons(Liters(40.0))
	busFuel += GallonsToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)

}
