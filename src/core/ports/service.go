package ports

import (
	"context"

	"sample.code/dataflow/src/core/domain"
)

type SyncDataService interface {
	SyncData(context.Context)
}

type RevenueService interface {
	GetRevenueDetails(ctx context.Context, dataRange domain.DateRangeRequest) (domain.Revenue, error)
}
