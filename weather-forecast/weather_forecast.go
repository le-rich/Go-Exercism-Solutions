// Package weather
// Gives the forecast for a given city and condition.
package weather

// CurrentCondition: Describes a weather condition.
var CurrentCondition string

// CurrentLocation: Describes the location of a forecast.
var CurrentLocation string

// Forecast: Returns a formatted forecast statement for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
