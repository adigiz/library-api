package middlewares

import (
	"git.garena.com/sea-labs-id/trainers/library-api/httperror"
	"github.com/gin-gonic/gin"
)

func RequestMiddleware(model any) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			badReqErr := httperror.BadRequestError(err.Error(), "")
			c.AbortWithStatusJSON(badReqErr.StatusCode, badReqErr)
			return
		}
		c.Set("payload", model)
		c.Next()
	}
}
