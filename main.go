package main

import (
	"net/http"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/maps"
	"github.com/NetKBs/backend-reviewapp/src/social/auth"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	image.RegisterRoutes(r)
	maps.RegisterRoutes(r)
	review.RegisterRoutes(r)
	auth.RegisterRoutes(r)

	r.Run()
}
