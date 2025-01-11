package domain

type Revenue struct {
	TotalRevenue float64 `json:"total_revenue"`
}

type ProductRevenue struct {
	ProductName  string  `json:"product_name"`
	TotalRevenue float64 `json:"total_revenue"`
}

type CategoryRevenue struct {
	Category     string  `json:"category"`
	TotalRevenue float64 `json:"total_revenue"`
}

type RegionRevenue struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
}

type TopProducts struct {
	ProductName string `json:"product_name"`
	TotalSold   int    `json:"total_sold"`
}

type CustomerAnalysis struct {
	TotalCustomers int     `json:"total_customers"`
	TotalOrders    int     `json:"total_orders"`
	AvgOrderValue  float64 `json:"avg_order_value"`
}
