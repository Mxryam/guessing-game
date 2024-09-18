package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("guessing game!")
	rand.Seed(time.Now().UnixNano())
	var maxNumber int
	fmt.Println("1.easy(0 until 10)")
	fmt.Println("2.medium(0 until 100)")
	fmt.Println("3.hard(0 until 1000)")

	var difficulty int
	fmt.Print("your choice is:")
	fmt.Scan(&difficulty)

	switch difficulty {
	case 1:
		maxNumber = 10
	case 2:
		maxNumber = 100
	case 3:
		maxNumber = 1000
	default:
		fmt.Println("invalid number!the game is continue in medium level.")
		maxNumber = 100

	}
	targetNumber := rand.Intn(maxNumber + 1)

	var guess int
	attempts := 0
	for {
		fmt.Print("your guess:")
		fmt.Scan(&guess)
		attempts++

		if guess < targetNumber {
			fmt.Println("guess a bigger number!")
		} else if guess < targetNumber {
			fmt.Println("guess a smaller number!")
		} else {
			fmt.Printf("congratulation!!you find correct number in %d guess\n", attempts)
			break
		}
	}
}
