// write two numbers using user input
package main 
import "fmt"
func main(){
	var num1,num2 int
	fmt.Println("Enter num1:")
	fmt.Scanln(&num1)
	fmt.Println("Enter num2:")
	fmt.Scanln(&num2)
	sum := num1 +num2
	fmt.Println("Sum of two number is",sum)

}