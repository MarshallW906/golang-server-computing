package asynchronous

import (
	"approach"
)

// Start ..
func Start() [10]approach.Info {
	customer := approach.GetCustomerDetails()
	destinations := approach.GetRecommendedDestinations(customer)

	var ret [10]approach.Info

	quotes := [10]chan approach.Quoting{}
	weathers := [10]chan approach.Weather{}

	for i := range quotes {
		quotes[i] = make(chan approach.Quoting)
	}

	for i := range weathers {
		weathers[i] = make(chan approach.Weather)
	}

	for idx, dest := range destinations {
		// must be a final variable
		i := idx
		dst := dest

		go func() {
			quotes[i] <- approach.GetQuote(dst)
		}()
		go func() {
			weathers[i] <- approach.GetWeatherForecast(dst)
		}()
	}

	for idx, dest := range destinations {
		ret[idx] = approach.Info{Destination: dest, Quote: <-quotes[idx], Weather: <-weathers[idx]}
	}
	return ret
}
