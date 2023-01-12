package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	name := c.Query("name")
	users, err := h.userService.GetUsers(name)
	
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}
