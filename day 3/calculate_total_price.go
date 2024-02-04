package main 
import "fmt"
func calculateTotalPrice(items map[string]float64)float64 {
total := 0.0
for _, item := range items {
	total+=item
}
return total
}
func main(){
	var numberOfItem int
    fmt.Println("Enter number of Items")
	fmt.Scanln(&numberOfItem)
    items := make(map[string]float64)

	
	for i := 0; i < numberOfItem; i++ {
		var itemName string
		var itemPrice float64
		fmt.Printf("Enter the name of item %d: ", i+1)
		fmt.Scanln(&itemName)
		fmt.Printf("Enter the price of item %s: ", itemName)
		fmt.Scanln(&itemPrice)
		items[itemName] = itemPrice
	}
	totalPrice := calculateTotalPrice(items)
	fmt.Printf("Total price of the order: $%.2f\n", totalPrice)
}