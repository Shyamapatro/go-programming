package main
import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	input := ""
	fmt.Print("Enter your age: \n")
	fmt.Scanln(&input)
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid age.")
		os.Exit(1)
	}

	// Check if the user is eligible to vote
	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else {
		fmt.Println("You are not eligible to vote yet.")
	}
}