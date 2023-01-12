package middlewares

import (
	"encoding/json"
	"fmt"
	"library-api/config"
	"library-api/httperror"
	"library-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

const (
	schema = "Bearer"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, httperror.UnauthorizedError()
		}
		return config.Config.JWTSecretKey, nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	if config.Config.ENV == "testing" {
		fmt.Println("disable JWT authorization on dev env")
		return
	}

	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, fmt.Sprintf("%v ", schema))
	authError := httperror.UnauthorizedError()
	if len(s) < 2 {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}
	decodedToken := s[1]
	token, err := validateToken(decodedToken)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}

	userJson, _ := json.Marshal(claims["user"])
	var user models.User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}
	c.Set("user", user)
}
