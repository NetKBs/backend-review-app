package images

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/images/:name", func(c *gin.Context) {
		name := c.Param("name")
		width := c.Query("width")
		height := c.Query("height")

		file, err := os.Open(fmt.Sprintf("images/%s", name))

		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("\"%s\" not found", name)})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		decoded, _, err := image.Decode(file)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		parsedWidth, err := strconv.ParseUint(width, 10, 32)

		if err != nil {
			parsedWidth = 0
			err = nil
		}

		parsedHeight, err := strconv.ParseUint(height, 10, 32)

		if err != nil {
			parsedHeight = 0
			err = nil
		}

		resized := resize.Resize(uint(parsedWidth), uint(parsedHeight), decoded, resize.Lanczos3)

		buf := new(bytes.Buffer)
		png.Encode(buf, resized)

		c.Data(http.StatusOK, "image/png", buf.Bytes())
	})
}
