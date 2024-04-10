package interfaces

import "fmt"

type NoiseMaker interface {
	MakeSound()
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Honk string

func (h Honk) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func Play(n NoiseMaker) {
	n.MakeSound()
}

func Test() {
	fmt.Println("Interfaces test")
}
