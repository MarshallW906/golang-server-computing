package synchronous

import (
	"approach"
)

// Start ..
func Start() [10]approach.Info {
	customer := approach.GetCustomerDetails()
	destinations := approach.GetRecommendedDestinations(customer)

	var ret [10]approach.Info
	var quotes [10]approach.Quoting
	var weathers [10]approach.Weather

	for i := range quotes {
		quotes[i] = approach.GetQuote(destinations[i])
	}
	for i := range weathers {
		weathers[i] = approach.GetWeatherForecast(destinations[i])
	}

	for idx, dest := range destinations {
		ret[idx] = approach.Info{Destination: dest, Quote: quotes[idx], Weather: weathers[idx]}
	}
	return ret
}
