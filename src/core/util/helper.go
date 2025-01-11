package util

import (
	"log"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// Atoi converts a string to an integer. If it fails, it logs the error and returns 0.
func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Conversion error for string to int: %v", err)
		return 0
	}
	return result
}

// Atof converts a string to a float64. If it fails, it logs the error and returns 0.
func Atof(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Conversion error for string to float: %v", err)
		return 0.0
	}
	return result
}

// ParseDate converts a string to a time.Time object. If it fails, it logs the error and returns the zero value of time.
func ParseDate(s string) time.Time {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Printf("Date parse error for string: %v", err)
		return time.Time{} // return zero time
	}
	return date
}

func IsValidDate(fl validator.FieldLevel) bool {
	// Get the value of the field
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	// Compare the year, month, and day only, ignoring the time
	zeroTime := time.Time{}
	if date.Year() == zeroTime.Year() && date.Month() == zeroTime.Month() && date.Day() == zeroTime.Day() {
		return false
	}
	return true
}
