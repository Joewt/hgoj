package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetDBMiddleware(c *gin.Context) {
	//TODO: https://gorm.io/docs/context.html
	c.Next()
}
