// Package weather provides a forecast function that returns the weather condition based on the location.
package weather

var (
    // CurrentCondition contains the current weather condition.
	CurrentCondition string
    // CurrentLocation is the current location.
	CurrentLocation  string
)

// Forecast returns a string with the current weather condition based on location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
