package ports

import (
	"golang.org/x/net/context"
	"sample.code/dataflow/src/core/domain"
)

type SyncDataRepository interface {
	InsertOrUpdateOrderData(context.Context, domain.OrderData) error
}
