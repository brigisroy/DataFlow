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
	OrderID         int       `csv:"order_id"`
	ProductID       int       `csv:"product_id"`
	CustomerID      int       `csv:"customer_id"`
	ProductName     string    `csv:"product_name"`
	Category        string    `csv:"category"`
	Region          string    `csv:"region"`
	DateOfSale      time.Time `csv:"date_of_sale"`
	QuantitySold    int       `csv:"quantity_sold"`
	UnitPrice       float64   `csv:"unit_price"`
	Discount        float64   `csv:"discount"`
	ShippingCost    float64   `csv:"shipping_cost"`
	PaymentMethod   string    `csv:"payment_method"`
	CustomerName    string    `csv:"customer_name"`
	CustomerEmail   string    `csv:"customer_email"`
	CustomerAddress string    `csv:"customer_address"`
}
