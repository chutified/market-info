package models

// Currency model holds the data about the currency.
type Currency struct {

	// Name is the currency code of the currency.
	Name string
	// The name of the country where the currency came from.
	Country string
	// Fullname of the currency.
	Description string

	// Latest currency change in percentages.
	Change float32
	// Exchange rate to USD.
	RateUSD float32
	// Last time updated.
	UpdatedAt string
}
