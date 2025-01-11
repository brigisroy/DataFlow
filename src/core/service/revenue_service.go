package service

import (
	"context"
	"fmt"

	"sample.code/dataflow/src/core/domain"
	"sample.code/dataflow/src/core/ports"
)

type revenueService struct {
	revenueRepo ports.RevenueRepository
}

func NewRevenueService(rr ports.RevenueRepository) ports.RevenueService {
	return &revenueService{
		revenueRepo: rr,
	}
}

func (rs *revenueService) GetRevenueDetails(ctx context.Context, dataRange domain.DateRangeRequest) (
	domain.Revenue, error,
) {
	revenue, err := rs.revenueRepo.GetRevenue(ctx, dataRange)
	if err != nil {
		return domain.Revenue{}, fmt.Errorf("failed to get revenue: %v", err)
	}
	return revenue, nil
}
