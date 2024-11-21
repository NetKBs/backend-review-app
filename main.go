package main

import (
	"net/http"
	"os"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/geoapify"
	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/maps"
	"github.com/NetKBs/backend-reviewapp/src/social/auth"
	"github.com/NetKBs/backend-reviewapp/src/social/place"
	"github.com/NetKBs/backend-reviewapp/src/social/review"

	//"github.com/V-enekoder/backend-review-app/src/social/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
	geoapify.SetGeoapifyKey(os.Getenv("GEOAPIFY_KEY"))
}

/*func main() {
	r := gin.Default()
	r.Use(cors.Default()) // Habilita CORS (considera restringirlo en producción)

	 // Registra la ruta para el controlador

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}*/

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	//r.GET("/user/:id", GetUserByIdController)
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

	r.Run()
}
