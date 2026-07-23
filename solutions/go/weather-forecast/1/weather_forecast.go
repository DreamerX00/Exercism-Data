// Package weather to get weather forecast.
package weather

var (
    // CurrentCondition of the weather.
	CurrentCondition string
    // CurrentLocation of the city to get exact weather forecast.
	CurrentLocation  string
)

// Forecast the weather of current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
