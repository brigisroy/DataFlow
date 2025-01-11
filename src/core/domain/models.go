package domain

import "time"

type Order struct {
	OrderID       int       `json:"order_id"`
	DateOfSale    time.Time `json:"date_of_sale"`
	Region        string    `json:"region"`
	Discount      float64   `json:"discount"`
	ShippingCost  float64   `json:"shipping_cost"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Product struct {
	ProductID   int       `json:"product_id"`
	UnitPrice   float64   `json:"unit_price"`
	ProductName string    `json:"product_name"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Customer struct {
	CustomerID int       `json:"customer_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderMapping struct {
	OrderID          int       `json:"order_id"`
	ProductID        int       `json:"product_id"`
	CustomerID       int       `json:"customer_id"`
	NoOfUnitsOrdered int       `json:"no_of_units_ordered"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type OrderData struct {
	OrderID         string    `csv:"order_id" validate:"required"`                // Order ID is required
	ProductID       string    `csv:"product_id" validate:"required"`              // Product ID is required
	CustomerID      string    `csv:"customer_id" validate:"required"`             // Customer ID is required
	ProductName     string    `csv:"product_name" validate:"required"`            // Product name is required
	Category        string    `csv:"category" validate:"required"`                // Category is required
	Region          string    `csv:"region" validate:"required"`                  // Region is required
	DateOfSale      time.Time `csv:"date_of_sale" validate:"required,valid_date"` // Date of sale must be valid and formatted as YYYY-MM-DD
	QuantitySold    int       `csv:"quantity_sold" validate:"required,gt=0"`      // Quantity sold must be greater than 0
	UnitPrice       float64   `csv:"unit_price" validate:"required,gt=0"`         // Unit price must be greater than 0
	Discount        float64   `csv:"discount" validate:"gte=0"`                   // Discount must be 0 or greater
	ShippingCost    float64   `csv:"shipping_cost" validate:"required,gt=0"`      // Shipping cost must be greater than 0
	PaymentMethod   string    `csv:"payment_method" validate:"required"`          // Payment method is required
	CustomerName    string    `csv:"customer_name" validate:"required"`           // Customer name is required
	CustomerEmail   string    `csv:"customer_email" validate:"required,email"`    // Valid email address required
	CustomerAddress string    `csv:"customer_address" validate:"required"`        // Customer address is required
}
