package service

import (
	"log"
	"time"

	"sample.code/dataflow/src/core/domain"
	"sample.code/dataflow/src/core/ports"
)

type syncDataService struct {
	filename     string
	syncDataRepo ports.SyncDataRepository
}

func NewSyncDataService(filename string, sdr ports.SyncDataRepository) ports.SyncDataService {
	return &syncDataService{
		filename:     filename,
		syncDataRepo: sdr,
	}
}

func (sds *syncDataService) SyncData() {
	orderData := domain.OrderData{
		OrderID:         123,
		ProductID:       1,
		CustomerID:      34,
		ProductName:     "Product A",
		Category:        "Category 1",
		Region:          "Region 1",
		DateOfSale:      time.Now(),
		QuantitySold:    5,
		UnitPrice:       20.5,
		Discount:        5.0,
		ShippingCost:    3.5,
		PaymentMethod:   "Credit Card",
		CustomerName:    "John Doe",
		CustomerEmail:   "john@example.com",
		CustomerAddress: "123 Main St",
	}

	err := sds.syncDataRepo.InsertOrUpdateOrderData(orderData)
	if err != nil {
		log.Println(err)
		return
	}
}
