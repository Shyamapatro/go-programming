// create a program that calculates the total revenue generated from a list of orders

package main 
import "fmt"

func calculateTotalRevenue(prices []float64)float64 {
	var totalRevenue float64
	for _, price := range prices{
		totalRevenue += price
	}
	return totalRevenue 
}

func main(){
	var numOrders int
	fmt.Println("Enter Number of orders")
	fmt.Scanln(&numOrders)
	orders := make([]float64, numOrders)
	for i := 0; i < numOrders; i++ {
		fmt.Println("Enter order price",i+1)
	    fmt.Scanln(&orders[i])
	}
	totalRevenue := calculateTotalRevenue(orders)
	fmt.Printf("Total revenue generated from the orders: $%.2f\n", totalRevenue)


}

