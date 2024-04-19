package main

import (
	"fmt"
	"gobyexamples/generics"
)

func main() {
	myIntList := generics.List[int]{}
	myIntList.Push(10)
	myIntList.Push(20)
	myIntList.Push(15)
	myIntList.Push(8)
	intElements := myIntList.GetAll()
	fmt.Println(intElements)

	myStringList := generics.List[string]{}
	myStringList.Push("jerry")
	myStringList.Push("tom")
	myStringList.Push("harry")

	stringElements := myStringList.GetAll()
	fmt.Println(stringElements)
}
