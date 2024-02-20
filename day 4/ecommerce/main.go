package main

import (
    "fmt"
    "ecommerce/ecommerce"
)

func main() {
    // Create a new product
    laptop := ecommerce.NewProduct(1, "Laptop", 1000, 5)

    // Display product details
    fmt.Println("Product ID:", laptop.ID)
    fmt.Println("Product Name:", laptop.Name)
    fmt.Println("Product Price:", laptop.Price)
    fmt.Println("Product Quantity:", laptop.Quantity)
}
