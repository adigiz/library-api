package handler

import (
	"library-api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	p, _ := c.Get("payload")
	payload, _ := p.(*dto.SignInReq)
	tokenResp, err := h.authService.SignIn(payload)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tokenResp)
}
