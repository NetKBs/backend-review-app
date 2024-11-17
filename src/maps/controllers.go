package maps

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetMapStyles(c *gin.Context) {
	resp, err := http.Get(fmt.Sprintf("https://maps.geoapify.com/v1/styles/osm-bright/style.json?apiKey=%s", os.Getenv("GEOAPIFY_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", respbody)
}
