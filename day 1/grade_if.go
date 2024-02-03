package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Prompt the user to enter the numerical grade
    fmt.Println("Please enter the numerical grade:")
    
    // Read user input with error handling
    numericalGrade, err := getUserInput()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Determine the corresponding letter grade
    var letterGrade string
    switch {
    case numericalGrade >= 90 && numericalGrade <= 100:
        letterGrade = "A"
    case numericalGrade >= 80 && numericalGrade <= 89:
        letterGrade = "B"
    case numericalGrade >= 70 && numericalGrade <= 79:
        letterGrade = "C"
    case numericalGrade >= 60 && numericalGrade <= 69:
        letterGrade = "D"
    default:
        letterGrade = "F"
    }

    // Print the corresponding letter grade
    fmt.Println("Letter grade:", letterGrade)
}

// getUserInput prompts the user for input and returns the numerical grade entered
func getUserInput() (int, error) {
    var input string
    fmt.Print(">> ")
    _, err := fmt.Scanln(&input)
    if err != nil {
        return 0, err
    }
    numericalGrade, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("invalid input: %v", err)
    }
    return numericalGrade, nil
}
