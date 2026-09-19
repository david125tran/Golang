package file_operations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Helper Functions -------------------------------------
func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetFloatFromFile(accountBalanceFile string) (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	// File not found
	if err != nil {
		return 1000, errors.New("Error: Failed to read file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	// Conversion error
	if err != nil {
		return 1000, errors.New("Error: Failed to parse stored balance value.")
	}

	return value, nil
}
