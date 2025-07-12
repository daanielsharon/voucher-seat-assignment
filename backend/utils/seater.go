package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomSeats(aircraft string) ([]string, error) {
	layouts := map[string]struct {
		Rows  int
		Seats []string
	}{
		"ATR":            {18, []string{"A", "C", "D", "F"}},
		"Airbus 320":     {32, []string{"A", "B", "C", "D", "E", "F"}},
		"Boeing 737 Max": {32, []string{"A", "B", "C", "D", "E", "F"}},
	}

	config, ok := layouts[aircraft]
	if !ok {
		return nil, errors.New("invalid aircraft type")
	}

	// This ensures the generated seat numbers are different every time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	seen := map[string]bool{}
	var result []string

	for len(result) < 3 {
		row := r.Intn(config.Rows) + 1 // x + 1 because row is 1-indexed
		col := config.Seats[r.Intn(len(config.Seats))]
		code := fmt.Sprintf("%d%s", row, col)
		if !seen[code] {
			seen[code] = true
			result = append(result, code)
		}
	}

	return result, nil
}
