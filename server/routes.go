package server

import (
	"github.com/chutified/market-info/handlers"
	"github.com/gin-gonic/gin"
)

// SetRoutes applies all the routings.
func SetRoutes(r *gin.Engine, h *handlers.Handler) {

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// v1 := r.Group("/v1")

	// cmd := v1.Group("/commodity")
	// {

	// }

	// crn := v1.Group("/currency")
	// {

	// }

	// crp := v1.Group("/crypto")
	// {

	// }
}
