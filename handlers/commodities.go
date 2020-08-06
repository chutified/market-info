package handlers

import "github.com/gin-gonic/gin"

// GetCommity serves the commodity data to the client.
func (h *Handler) GetCommodity(c *gin.Context) {

	// get requested the commodity
	name := c.Param("name")
	cmd, err := h.commoditySrv.GetCommodity(name)
	if err != nil {
		c.JSON() // TODO
	}

}
