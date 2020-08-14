package handlers

import "github.com/gin-gonic/gin"

// GetCurrency provides the data about the given currency.
func (h *Handler) GetCurrency(c *gin.Context) {

	name := c.Params("name")

	crn, err := h.currencySrv.GetCurrency(name)
	if err != nil {
		c.JSON(400, gin.H{"error": gin.H{"error": err.Error()}})
	}

	c.JSON(200, crn)
}

// GetRate handles the request and calculates the exchange rate
// between two currencies.
func (h *Handler) GetRate(c *gin.Context) {

	base := c.Params("base")
	dest := c.Params("dest")

	// calculate the rate
	rate, err := h.cryptoSrv.GetRate(base, dest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, rate)
}
