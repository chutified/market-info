package server

import (
	handlers "github.com/chutified/market-info/handlers"
	"github.com/gin-gonic/gin"
)

// SetRoutes applies all the routings.
func SetRoutes(r *gin.Engine, h *handlers.Handler) {

	// ping to test server status
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/v1")
	{

		// comodities
		cmd := v1.Group("/commodity")
		{
			cmd.GET("/:name", h.GetCommodity)
		}

		// currencies
		crn := v1.Group("/currency")
		{
			crn.GET("/:name", h.GetCurrency)
			crn.POST("/:base/:dest", h.GetRate)
		}

		// cryptocurrencies
		crp := v1.Group("/crypto")
		{
			crp.GET("/:name", h.GetCrypto)
		}
	}
}
