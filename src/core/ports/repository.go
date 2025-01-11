package ports

import "sample.code/dataflow/src/core/domain"

type SyncDataRepository interface {
	InsertOrUpdateOrderData(data domain.OrderData) error
}
