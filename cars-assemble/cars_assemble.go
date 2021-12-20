package cars

import (
	"math"
)

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch {
	case speed == 0:
		return 0.0
	case speed <= 4:
		return 1.0
	case speed <= 8:
		return 0.9
	case speed <= 10:
		return 0.77
	default:
		return -1.0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed) * 221.0 * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(math.Floor(CalculateProductionRatePerHour(speed)) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	return math.Min(CalculateProductionRatePerHour(speed), limit)
}
