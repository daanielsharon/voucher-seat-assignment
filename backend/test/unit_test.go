package test

import (
	"strconv"
	"testing"
	"unicode"
	"voucher-seat-assignment/utils"

	"slices"

	"github.com/stretchr/testify/assert"
)

func parseSeat(seat string) (int, string, error) {
	var rowStr, colStr string

	for _, r := range seat {
		if unicode.IsDigit(r) {
			rowStr += string(r)
		} else {
			colStr += string(r)
		}
	}

	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return 0, "", err
	}

	return row, colStr, nil
}

func isValidRow(row int, maxRow int, minRow int) bool {
	return row >= minRow && row <= maxRow
}

func isValidSeatLetter(allowedLetters []string, letter string) bool {
	return slices.Contains(allowedLetters, letter)
}

func TestGenerateRandomSeats(t *testing.T) {
	t.Run("GenerateRandomSeats", func(t *testing.T) {
		t.Run("ATR", func(t *testing.T) {
			maxRow := 18
			minRow := 1
			allowedLetters := []string{"A", "C", "D", "F"}

			seats, err := utils.GenerateRandomSeats("ATR")
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, 3, len(seats))

			seen := make(map[string]bool)
			for _, seat := range seats {
				row, col, err := parseSeat(seat)
				if err != nil {
					t.Fatal(err)
				}

				assert.True(t, isValidRow(row, maxRow, minRow))
				assert.True(t, isValidSeatLetter(allowedLetters, col))
				assert.False(t, seen[seat])

				seen[seat] = true
			}
		})

		t.Run("Airbus 320", func(t *testing.T) {
			maxRow := 32
			minRow := 1
			allowedLetters := []string{"A", "B", "C", "D", "E", "F"}

			seats, err := utils.GenerateRandomSeats("Airbus 320")
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, 3, len(seats))

			seen := make(map[string]bool)
			for _, seat := range seats {
				row, col, err := parseSeat(seat)
				if err != nil {
					t.Fatal(err)
				}

				assert.True(t, isValidRow(row, maxRow, minRow))
				assert.True(t, isValidSeatLetter(allowedLetters, col))
				assert.False(t, seen[seat])

				seen[seat] = true
			}
		})

		t.Run("Boeing 737 Max", func(t *testing.T) {
			maxRow := 32
			minRow := 1
			allowedLetters := []string{"A", "B", "C", "D", "E", "F"}

			seats, err := utils.GenerateRandomSeats("Boeing 737 Max")
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, 3, len(seats))

			seen := make(map[string]bool)
			for _, seat := range seats {
				row, col, err := parseSeat(seat)
				if err != nil {
					t.Fatal(err)
				}

				assert.True(t, isValidRow(row, maxRow, minRow))
				assert.True(t, isValidSeatLetter(allowedLetters, col))
				assert.False(t, seen[seat])

				seen[seat] = true
			}
		})
	})
}
