package domain

type Revenue struct {
	TotalRevenue           float64           `json:"totalRevenue"`
	TotalRevenueByProduct  []ProductRevenue  `json:"product"`
	TotalRevenueByCategory []CategoryRevenue `json:"category"`
	TotalRevenueByRegion   []RegionRevenue   `json:"region"`
}

type ProductRevenue struct {
	ProductName  string  `json:"product"`
	TotalRevenue float64 `json:"revenue"`
}

type CategoryRevenue struct {
	Category     string  `json:"category"`
	TotalRevenue float64 `json:"revenue"`
}

type RegionRevenue struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"revenue"`
}
