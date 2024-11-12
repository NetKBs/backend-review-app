package main

import (
	"net/http"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/images"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	images.RegisterRoutes(r)
	review.RegisterRoutes(r)

	r.Run()
}
