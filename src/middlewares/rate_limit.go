package middlewares

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ipData struct {
	requests int
	lastTime time.Time
}

var ipMap = make(map[string]*ipData)
var mutex = &sync.Mutex{}

func RateLimitMiddleware(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Get("userId")
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
			c.Abort()
			return
		}
		userIdStr := strconv.Itoa(int(userId.(uint)))

		mutex.Lock()
		defer mutex.Unlock()

		if _, ok := ipMap[userIdStr]; !ok {
			ipMap[userIdStr] = &ipData{
				requests: 1,
				lastTime: time.Now(),
			}
		} else {

			diff := time.Since(ipMap[userIdStr].lastTime)
			if diff > duration {
				ipMap[userIdStr].requests = 1
				ipMap[userIdStr].lastTime = time.Now()

			} else if ipMap[userIdStr].requests >= limit {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
				c.Abort()
				return
			} else {
				ipMap[userIdStr].requests++
				ipMap[userIdStr].lastTime = time.Now()
			}
		}

		c.Next()
	}
}
