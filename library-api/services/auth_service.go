package services

import (
	"git.garena.com/sea-labs-id/trainers/library-api/config"
	"git.garena.com/sea-labs-id/trainers/library-api/dto"
	"git.garena.com/sea-labs-id/trainers/library-api/httperror"
	"git.garena.com/sea-labs-id/trainers/library-api/models"
	repository "git.garena.com/sea-labs-id/trainers/library-api/repositories"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthService interface {
	SignIn(*dto.SignInReq) (*dto.TokenResponse, error)
}

type authService struct {
	userRepository repository.UserRepository
	config         config.AppConfig
}

type AuthSConfig struct {
	UserRepository repository.UserRepository
	Config         config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		config:         c.Config,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

func (a *authService) generateJWTToken(user *models.User) (*dto.TokenResponse, error) {
	var idExp = a.config.JWTExpiryInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp

	claims := &idTokenClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.config.AppName,
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(tokenExp, 0)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.config.JWTSecretKey)
	if err != nil {
		return nil, httperror.UnauthorizedError()
	}
	return &dto.TokenResponse{IDToken: tokenString}, nil
}

func (a *authService) SignIn(req *dto.SignInReq) (*dto.TokenResponse, error) {
	user, err := a.userRepository.MatchingCredential(req.Email, req.Password)
	if err != nil || user == nil {
		return nil, httperror.UnauthorizedError()
	}
	token, err := a.generateJWTToken(user)
	return token, err
}
