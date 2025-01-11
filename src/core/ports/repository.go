package ports

import (
	"golang.org/x/net/context"
	"sample.code/dataflow/src/core/domain"
)

type SyncDataRepository interface {
	InsertOrUpdateOrderData(context.Context, domain.OrderData) error
}

type RevenueRepository interface {
	GetRevenue(ctx context.Context, dataRange domain.DateRangeRequest) (domain.Revenue, error)
}
