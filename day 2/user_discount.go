// You are developing an e-commerce application. 
// Write a Go program that prompts the user to enter the original price of an item and the discount percentage applicable. 
// Calculate the discounted price and display it to the user.

package main 
import "fmt"
func calculateTotalDiscount(originalPrice float64,discountPercentage int) float64{
	var priceAfterDiscount float64 
	discount := originalPrice * (float64(discountPercentage) / 100)
	priceAfterDiscount=originalPrice - discount
	return priceAfterDiscount
}
func main(){
	var orginalPrice float64
	fmt.Println("Enter the original price of an item\n")
	fmt.Scanln(&orginalPrice)
	var discountPercentage int
	fmt.Println("Enter the discount\n")
	fmt.Scanln(&discountPercentage)
	price:=calculateTotalDiscount(orginalPrice,discountPercentage)
	fmt.Println("After the discount item price is:",price)
	
}