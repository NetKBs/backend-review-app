package main

import (
	"net/http"
	"os"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/geoapify"
	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/maps"
	"github.com/NetKBs/backend-reviewapp/src/social/auth"
	"github.com/NetKBs/backend-reviewapp/src/social/comment"
	"github.com/NetKBs/backend-reviewapp/src/social/place"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
	"github.com/NetKBs/backend-reviewapp/src/social/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
	geoapify.SetGeoapifyKey(os.Getenv("GEOAPIFY_KEY"))
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
	place.RegisterRoutes(r)
	user.RegisterRoutes(r)
	comment.RegisterRoutes(r)
	reaction.RegisterRoutes(r)
	r.Run()
}
