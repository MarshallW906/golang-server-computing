package approach

import (
	"time"
)

// Weather ..
type Weather struct{}

// Destination ..
type Destination struct{}

// Quoting ..
type Quoting struct{}

// Customers ..
type Customers struct{}

// Info ..
type Info struct {
	Destination Destination
	Quote       Quoting
	Weather     Weather
}

// GetCustomerDetails ..
func GetCustomerDetails() Customers {
	time.Sleep(150 * time.Millisecond)
	return Customers{}
}

// GetRecommendedDestinations ..
func GetRecommendedDestinations(customers Customers) [10]Destination {
	time.Sleep(250 * time.Millisecond)
	return [10]Destination{}
}

// GetQuote ..
func GetQuote(destination Destination) Quoting {
	time.Sleep(170 * time.Millisecond)
	return Quoting{}
}

// GetWeatherForecast ..
func GetWeatherForecast(destination Destination) Weather {
	time.Sleep(330 * time.Millisecond)
	return Weather{}
}
