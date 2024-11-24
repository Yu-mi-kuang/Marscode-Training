package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	secretNumber := r.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess:")

	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again:", err)
			continue
		}

		fmt.Println("Your guess is:", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again.")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
