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

	// TODO
	// v1 := r.Group("/v1")

	// cmd := v1.Group("/commodity")
	// {
	// TODO GetCommodity{name}
	// }

	// crn := v1.Group("/currency")
	// {

	// }

	// crp := v1.Group("/crypto")
	// {

	// }
}
