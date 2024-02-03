package main
import "fmt"
func main(){
	var num1 int
	var num2 int
	var num3 int
	fmt.Println("Enter first Number:\n");
	fmt.Scanln(&num1);
	fmt.Println("Enter Second Number:\n");
	fmt.Scanln(&num2);
	fmt.Println("Enter third Number:\n");
	fmt.Scanln(&num3);
	multiplying :=num1*num2*num3
	fmt.Println("After multiplying three numbers\n");
	fmt.Println(multiplying);
}