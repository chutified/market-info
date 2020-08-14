package handlers

import "github.com/gin-gonic/gin"

// GetCurrency provides the data about the given currency.
func (h *Handler) GetCurrency(c *gin.Context) {

	name := c.Param("name")

	crn, err := h.currencySrv.GetCurrency(name)
	if err != nil {
		c.JSON(400, gin.H{"error": gin.H{"error": err.Error()}})
	}

	c.JSON(200, crn)
}

// GetRate handles the request and calculates the exchange rate
// between two currencies.
func (h *Handler) GetRate(c *gin.Context) {

	base := c.Param("base")
	dest := c.Param("dest")

	// calculate the rate
	rate, err := h.currencySrv.GetRate(base, dest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, rate)
}
