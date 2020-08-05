package handlers

// InvalidParam defines the error returned by an invalid parameter.
type InvalidParam struct {
	// Parameter is the name of the parameter.
	Parameter string `json:"parameter"`
	// Value is the value of the Parameter.
	Value string `json:"value"`
	// Error is the returned error.
	Error error `json:"error"`
}
