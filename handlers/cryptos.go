package handlers

import "github.com/gin-gonic/gin"

// GetCrypto handles the GetCrypto request and response with the
// latest available data about the given cryptocurrency.
func (h *Handler) GetCrypto(c *gin.Context) {

	name := c.Param("name")
	crp, err := h.cryptoSrv.GetCrypto(name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, crp)
}
