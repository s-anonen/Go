// Package weather, func Forecast.
package weather

var (
    // CurrentCondition - condition.
	CurrentCondition string
    // CurrentLocation - city.
	CurrentLocation  string
)

// Forecast returns a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}