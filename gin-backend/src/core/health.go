package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHealthRoute expone un endpoint liviano para probes (Fly.io, monitoreo, Play Store).
func RegisterHealthRoute(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("healthy\n"))
	})
}
