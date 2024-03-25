package hello

import (
	"fmt"
	"headfirst-go/greeting"
)

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func SayHello() {

	/*
		fmt.Println("Hello, Go!")
		fmt.Println(math.Floor(2.75))
		fmt.Println(reflect.TypeOf(42))
		fmt.Println(reflect.TypeOf(3.1415))
		fmt.Println(reflect.TypeOf(true))
		fmt.Println(reflect.TypeOf("hello, go!"))

		quantity := 4
		length, width := 1.2, 2.4
		customerName := "Maulin Vasavada"

		fmt.Println(customerName)
		fmt.Println("has ordered", quantity, "sheets")
		fmt.Println("each with an area of")
		fmt.Println(length*width, "suqare meters")

		fmt.Print("Enter a grade: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(reflect.TypeOf(input))
			input = strings.TrimSpace(input)
			grade, err := strconv.ParseFloat(input, 64)
			if err != nil {
				panic(err)
			}
			fmt.Println("Grade is", grade)
			if grade >= 60 {
				fmt.Println("You passed!")
			} else {
				fmt.Println("You failed!")
			}
		}

		myInt := 4
		fmt.Println(myInt)
		myIntPointer := &myInt
		*myIntPointer = 8
		fmt.Println("My Int Pointer ", *myIntPointer)
		fmt.Println("My Int ", myInt)

	*/

	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)

	greeting.Greetings()
}
