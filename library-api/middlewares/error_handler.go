package middlewares

import (
	"library-api/httperror"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // run handler and middleware chains

	if len(c.Errors) == 0 {
		return
	}
	firstErr := c.Errors[0].Err
	appErr, isAppError := firstErr.(httperror.AppError) // type assertion
	if isAppError {
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		return
	}
	serverErr := httperror.InternalServerError(firstErr.Error())
	c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
}
