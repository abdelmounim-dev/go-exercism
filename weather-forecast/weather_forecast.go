// Package weather provides current weather conditions.
package weather

// CurrentCondition is package level variable.
var CurrentCondition string
// CurrentLocation is package level variable.
var CurrentLocation string

// Forecast returns a weather forecast for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
