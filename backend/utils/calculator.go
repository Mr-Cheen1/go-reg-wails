package utils

import (
	"strconv"
	"strings"
)

// CalculateTime вычисляет общее время из строки с формулой
func CalculateTime(timeStr string) float64 {
	if strings.TrimSpace(timeStr) == "" {
		return 0
	}
	parts := strings.Split(timeStr, "+")
	var total float64
	for _, part := range parts {
		val, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err == nil {
			total += val
		}
	}
	return total
}
