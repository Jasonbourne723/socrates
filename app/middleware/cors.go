package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://localhost:7037", "http://localhost:5274"}
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"New-Token", "New-Expires-In", "Content-Disposition"}

	return cors.New(config)
}
