package ports

import "context"

type SyncDataService interface {
	SyncData(context.Context)
}
