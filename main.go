package main

import (
	"net/http"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/maps"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
}

/*func main() {
	r := gin.Default()
	r.Use(cors.Default()) // Habilita CORS (considera restringirlo en producción)

	r.GET("/user/:id", GetUserByIdController) // Registra la ruta para el controlador

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}*/

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

	r.Run()
}
