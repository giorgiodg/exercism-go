// Package weather provides a library for obtaining weather forecast information.
package weather

// CurrentCondition is a variable holding the current weather condition.
var CurrentCondition string
// CurrentLocation is a variable holding the current location.
var CurrentLocation string

// Forecast returns the forecast for the given city and its condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
