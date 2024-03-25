package guessinggame

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readNumberInput(promptMessage string, maxAllowedInput int) int {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptMessage)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess >= maxAllowedInput {
		log.Fatal("You can't select more than ", maxAllowedInput)
	}
	return guess
}

func generateTarget() int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	return target
}

func PlayGame() {
	allowedChances := readNumberInput("How many chances you want to allow a player? ", 100)

	target := generateTarget()
	fmt.Println("I've chosen a random numer between 1 and 100. Can you guess it? You will have", allowedChances, "chances to guess")

	var chance int
	for chance = 1; chance < allowedChances+1; chance++ {
		fmt.Println("Chance# ", chance)

		guess := readNumberInput("Make a guess: ", 100)
		if guess < target {
			fmt.Println("Oops, your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops, your guess was HIGH")
		} else {
			fmt.Println("You got it right!")
			break
		}
	}

	if chance == allowedChances+1 {
		fmt.Println("You ran out of all the chances. Sorry, better luck next time! The target was", target)
	}
}
