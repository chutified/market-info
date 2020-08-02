package models

// Crypto model holds the data about the commodity.
type Crypto struct {

	// Name of the currency.
	Name string
	// Symbol of the currency.
	Symbol string

	// Market capitalization in USD.
	MarketCapUSD float64
	// Cuurrent value of the currency.
	Price float64
	// The total value of the currently available amount of the currencies.
	CirculatingSupply float64
	// Mineable indicates if the currency is mineable.
	Mineable bool
	// Volume is the total value of the currencies in USD which was traded in the last 24 hours.
	Volume float64

	// The percentage changes in the last hour/day/week.
	ChangeHour string
	ChangeDay  string
	ChangeWeek string
}
