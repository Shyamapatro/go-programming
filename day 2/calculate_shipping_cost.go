// Shipping Cost Calculation:

// You are building an e-commerce platform. Write a Go program that calculates the shipping cost based on the weight of the items ordered. The shipping cost is calculated according to the following rules:

// If the total weight of the items is less than or equal to 5 kg, the shipping cost is $10.
// If the total weight is greater than 5 kg but less than or equal to 10 kg, the shipping cost is $20.
// If the total weight is greater than 10 kg, the shipping cost is $30.
// Write a program that prompts the user to enter the weights of each item in the order, calculates the total weight, and then calculates and displays the shipping cost.
package main 
import (
"fmt"
)

func calculateShippingCost(itemWeight float64) float64 {
	var totalCost float64
	if itemWeight<5 {
		totalCost+=10
	}else if itemWeight>5 && itemWeight<=10{
		totalCost+=20
	}else if itemWeight>10{
		totalCost+=30
	}
	return totalCost
}

func main(){
    var itemWeight float64
	fmt.Println("Enter the weight of shipping item")
    fmt.Scanln(&itemWeight)
    shippingCost:=calculateShippingCost(itemWeight)
	fmt.Printf("So the shipping cost for the order with total weight %.2f kg is: $%.2f\n", itemWeight, shippingCost)
}
