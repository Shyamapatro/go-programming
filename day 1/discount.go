// Scenario: Online Shopping Cart Discount

// Suppose you are developing an e-commerce platform, and you need to implement a discount system for shopping cart items. The discount rules are as follows:

// If the total price of the items in the cart is greater than $100, apply a 10% discount.
// If the total price is between $50 and $100 (inclusive), apply a flat $5 discount.
// If the total price is less than $50, no discount is applied.

package main

import "fmt"

func applyDiscount(amount float64) float64 {
	var discountedTotal float64
	if amount > 100 {
		discountedTotal = amount - (amount * 0.1)
	} else if amount >= 50 && amount <= 100 {
		discountedTotal = amount + 5
	} else if amount < 50 {
		discountedTotal = amount
	}

	return discountedTotal
}
func main() {
	var amount float64
	fmt.Println("Please enter a card Amount:\n ")
	fmt.Scanln(&amount)
	discountedTotal := applyDiscount(amount)
	fmt.Printf("Total price after discount: $%.2f\n", discountedTotal)

}
