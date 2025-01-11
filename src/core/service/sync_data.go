package service

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"sample.code/dataflow/src/core/domain"
	"sample.code/dataflow/src/core/ports"
	"sample.code/dataflow/src/core/util"
)

type syncDataService struct {
	filename     string
	syncDataRepo ports.SyncDataRepository
}

var validate *validator.Validate

func NewSyncDataService(filename string, sdr ports.SyncDataRepository) ports.SyncDataService {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("valid_date", util.IsValidDate)
	}

	return &syncDataService{
		filename:     filename,
		syncDataRepo: sdr,
	}
}

func (sds *syncDataService) SyncData(ctx context.Context) {
	// Open the CSV file
	file, err := os.Open(sds.filename)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header row
	if err != nil {
		log.Fatalf("Error reading header: %v", err)
	}
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		// Map the CSV record to OrderData
		order := domain.OrderData{
			OrderID:         record[0],
			ProductID:       record[1],
			CustomerID:      record[2],
			ProductName:     record[3],
			Category:        record[4],
			Region:          record[5],
			DateOfSale:      util.ParseDate(record[6]),
			QuantitySold:    util.Atoi(record[7]),
			UnitPrice:       util.Atof(record[8]),
			Discount:        util.Atof(record[9]),
			ShippingCost:    util.Atof(record[10]),
			PaymentMethod:   record[11],
			CustomerName:    record[12],
			CustomerEmail:   record[13],
			CustomerAddress: record[14],
		}

		// Validate the order
		err = validate.Struct(order)
		if err != nil {
			log.Printf("Validation failed for OrderID %s: %v", order.OrderID, err)
			continue
		}

		// Insert or update the data
		err = sds.syncDataRepo.InsertOrUpdateOrderData(ctx, order)
		if err != nil {
			log.Println(err)
		}
	}
}
