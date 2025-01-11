package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"sample.code/dataflow/src/core/domain"
	"sample.code/dataflow/src/core/ports"
)

type RevenueRepository struct {
	dbPool *pgxpool.Pool
}

func NewRevenueRepository(dbPool *pgxpool.Pool) ports.RevenueRepository {
	return &RevenueRepository{
		dbPool: dbPool,
	}
}

func (rr *RevenueRepository) GetRevenue(ctx context.Context, dataRange domain.DateRangeRequest) (
	domain.Revenue, error,
) {
	// Define the SQL queries
	totalRevenueQuery := `
		SELECT
			SUM(om.no_of_units_ordered * p.unit_price * (1 - o.discount / 100) + o.shipping_cost) AS total_revenue
		FROM 
			dataflow.orders o
		JOIN 
			dataflow.order_mapping om ON om.order_id = o.order_id
		JOIN 
			dataflow.products p ON p.product_id = om.product_id
		WHERE 
			o.date_of_sale BETWEEN $1 AND $2
	`

	productRevenueQuery := `
		SELECT
			SUM(om.no_of_units_ordered * p.unit_price * (1 - o.discount / 100) + o.shipping_cost) AS total_revenue,
			p.product_name
		FROM 
			dataflow.orders o
		JOIN 
			dataflow.order_mapping om ON om.order_id = o.order_id
		JOIN 
			dataflow.products p ON p.product_id = om.product_id
		WHERE 
			o.date_of_sale BETWEEN $1 AND $2
		GROUP BY 
			p.product_name
	`

	categoryRevenueQuery := `
		SELECT
			SUM(om.no_of_units_ordered * p.unit_price * (1 - o.discount / 100) + o.shipping_cost) AS total_revenue,
			p.category
		FROM 
			dataflow.orders o
		JOIN 
			dataflow.order_mapping om ON om.order_id = o.order_id
		JOIN 
			dataflow.products p ON p.product_id = om.product_id
		WHERE 
			o.date_of_sale BETWEEN $1 AND $2
		GROUP BY 
			p.category
	`

	regionRevenueQuery := `
		SELECT
			SUM(om.no_of_units_ordered * p.unit_price * (1 - o.discount / 100) + o.shipping_cost) AS total_revenue,
			o.region
		FROM 
			dataflow.orders o
		JOIN 
			dataflow.order_mapping om ON om.order_id = o.order_id
		JOIN 
			dataflow.products p ON p.product_id = om.product_id
		WHERE 
			o.date_of_sale BETWEEN $1 AND $2
		GROUP BY 
			o.region
	`

	var revenue domain.Revenue
	var totalRevenue float64

	// Execute the Total Revenue Query
	err := rr.dbPool.QueryRow(
		ctx, totalRevenueQuery, dataRange.StartDate, dataRange.EndDate,
	).Scan(&totalRevenue)
	if err != nil {
		return revenue, fmt.Errorf("error executing total revenue query: %v", err)
	}

	// Execute the Product Revenue Query
	productRows, err := rr.dbPool.Query(ctx, productRevenueQuery, dataRange.StartDate, dataRange.EndDate)
	if err != nil {
		return revenue, fmt.Errorf("error executing product revenue query: %v", err)
	}
	defer productRows.Close()

	categoryRows, err := rr.dbPool.Query(ctx, categoryRevenueQuery, dataRange.StartDate, dataRange.EndDate)
	if err != nil {
		return revenue, fmt.Errorf("error executing category revenue query: %v", err)
	}
	defer categoryRows.Close()

	// Execute the Region Revenue Query
	regionRows, err := rr.dbPool.Query(ctx, regionRevenueQuery, dataRange.StartDate, dataRange.EndDate)
	if err != nil {
		return revenue, fmt.Errorf("error executing region revenue query: %v", err)
	}
	defer regionRows.Close()

	productRevenues := make([]domain.ProductRevenue, 0)
	for productRows.Next() {
		var productName string
		var productRevenue float64
		if err := productRows.Scan(&productRevenue, &productName); err != nil {
			return revenue, fmt.Errorf("error scanning product row: %v", err)
		}
		productRevenues = append(
			productRevenues, domain.ProductRevenue{
				ProductName:  productName,
				TotalRevenue: productRevenue,
			},
		)
	}

	categoryRevenues := make([]domain.CategoryRevenue, 0)
	for categoryRows.Next() {
		var category string
		var categoryRevenue float64
		if err := categoryRows.Scan(&categoryRevenue, &category); err != nil {
			return revenue, fmt.Errorf("error scanning category row: %v", err)
		}
		categoryRevenues = append(
			categoryRevenues, domain.CategoryRevenue{
				Category:     category,
				TotalRevenue: categoryRevenue,
			},
		)
	}

	regionRevenues := make([]domain.RegionRevenue, 0)
	for regionRows.Next() {
		var region string
		var regionRevenue float64
		if err := regionRows.Scan(&regionRevenue, &region); err != nil {
			return revenue, fmt.Errorf("error scanning region row: %v", err)
		}
		regionRevenues = append(
			regionRevenues, domain.RegionRevenue{
				Region:       region,
				TotalRevenue: regionRevenue,
			},
		)
	}

	revenue.TotalRevenueByProduct = productRevenues
	revenue.TotalRevenueByCategory = categoryRevenues
	revenue.TotalRevenueByRegion = regionRevenues
	revenue.TotalRevenue = totalRevenue

	return revenue, nil
}
