// Scenario: Grade Calculator

// Suppose you are developing a program to calculate the grade of a student based on their test score. The grading system is as follows:

// A grade of 90 or above is an "A".
// A grade of 80-89 is a "B".
// A grade of 70-79 is a "C".
// A grade of 60-69 is a "D".
// Anything below 60 is an "F".

package main 
import "fmt"

func calculateGrade(score int) string{
	var grade string
	if score >= 90{
		grade = "A"
	}else if score >= 80 && score <= 89 {
		grade = "B"
	}else if score >= 70 && score <= 79{
		grade = "C"
	}else if score < 60 && score <= 69{
		grade = "D"
	}else{
        grade = "F"
	}
	return grade
}
func main(){
  var score int
  fmt.Printf("Enter Your Score: \n")
  fmt.Scanln(&score)
  grade := calculateGrade(score)
    fmt.Println("Your grade is:", grade)
}