package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"sample.code/dataflow/src/core/domain"
	"sample.code/dataflow/src/core/ports"
)

type SyncDataRepository struct {
	dbPool *pgxpool.Pool
}

func NewSyncDataRepository(dbPool *pgxpool.Pool) ports.SyncDataRepository {
	return &SyncDataRepository{
		dbPool: dbPool,
	}
}

func (sdr *SyncDataRepository) InsertOrUpdateOrderData(ctx context.Context, data domain.OrderData) error {
	// Begin Transaction
	tx, err := sdr.dbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx) // rollback in case of error

	// Insert or else Update into Order Table
	orderQuery := `
		INSERT INTO dataflow.orders (order_id, date_of_sale, region, discount, shipping_cost, payment_method, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		ON CONFLICT (order_id)
		DO UPDATE SET
			date_of_sale = EXCLUDED.date_of_sale,
			region = EXCLUDED.region,
			discount = EXCLUDED.discount,
			shipping_cost = EXCLUDED.shipping_cost,
			payment_method = EXCLUDED.payment_method,
			updated_at = NOW();`

	_, err = tx.Exec(
		ctx, orderQuery, data.OrderID, data.DateOfSale, data.Region, data.Discount, data.ShippingCost,
		data.PaymentMethod,
	)
	if err != nil {
		return fmt.Errorf("error upserting order %s: %v", data.OrderID, err)
	}

	// Insert or Update on the Product Table
	productQuery := `
		INSERT INTO dataflow.products (product_id, unit_price, product_name, category, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		ON CONFLICT (product_id)
		DO UPDATE SET
			unit_price = EXCLUDED.unit_price,
			product_name = EXCLUDED.product_name,
			category = EXCLUDED.category,
			updated_at = NOW();`

	_, err = tx.Exec(
		ctx, productQuery, data.ProductID, data.UnitPrice, data.ProductName, data.Category,
	)
	if err != nil {
		return fmt.Errorf("error upserting product %s: %v", data.ProductID, err)
	}

	// Insert or Update the Customer Table
	customerQuery := `
		INSERT INTO dataflow.customers (customer_id, name, email, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		ON CONFLICT (customer_id)
		DO UPDATE SET
			name = EXCLUDED.name,
			email = EXCLUDED.email,
			address = EXCLUDED.address,
			updated_at = NOW();`

	_, err = tx.Exec(
		ctx, customerQuery, data.CustomerID, data.CustomerName, data.CustomerEmail,
		data.CustomerAddress,
	)
	if err != nil {
		return fmt.Errorf("error upserting customer %s: %v", data.CustomerID, err)
	}

	// Insert or Update the Order Mapping Table
	orderMappingQuery := `
		INSERT INTO dataflow.order_mapping (order_id, product_id, customer_id, no_of_units_ordered, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		ON CONFLICT (order_id, product_id, customer_id)
		DO UPDATE SET
			no_of_units_ordered = EXCLUDED.no_of_units_ordered,
			updated_at = NOW();`

	_, err = tx.Exec(
		ctx, orderMappingQuery, data.OrderID, data.ProductID, data.CustomerID, data.QuantitySold,
	)
	if err != nil {
		return fmt.Errorf("error upserting order mapping for order %s: %v", data.OrderID, err)
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	// return nil if now error
	return nil
}
