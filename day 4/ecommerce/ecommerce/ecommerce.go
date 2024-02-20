package ecommerce

// Product represents a product in the e-commerce system
type Product struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

// NewProduct creates a new product with the provided details
func NewProduct(id int, name string, price float64, quantity int) *Product {
    return &Product{
        ID:       id,
        Name:     name,
        Price:    price,
        Quantity: quantity,
    }
}
