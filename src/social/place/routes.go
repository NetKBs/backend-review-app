package place

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/places", middlewares.AuthMiddleware())

	group.GET("/details", getPlaceDetailsController)
	group.GET("/autocomplete", getAutocompleteResultController)
	group.GET("/", getPlacesController)
}
