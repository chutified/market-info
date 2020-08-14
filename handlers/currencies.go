package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) GetCurrency(c *gin.Context) {

	name := c.Params("name")
}

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
