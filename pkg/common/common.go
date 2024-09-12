package common

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, body any) {
	c.JSON(code, body)
}
