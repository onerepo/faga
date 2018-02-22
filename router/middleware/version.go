package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/onerepo/faga/internal/version"
)

// Version is a middleware function that appends the Drone version information
// to the HTTP response. This is intended for debugging and troubleshooting.
func Version(c *gin.Context) {
	c.Header("X-FAGA-VERSION", version.Version.String())
}
