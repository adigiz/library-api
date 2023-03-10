package server

import (
	"fmt"
	"library-api/config"
	"library-api/db"
	repository "library-api/repositories"
	"library-api/services"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	bookRepository := repository.NewBookRepository(&repository.BROpts{
		DB: db.Get(),
	})
	bookService := services.NewBookService(&services.BSOpts{
		BookRepository: bookRepository,
	})
	userRepository := repository.NewUserRepository(&repository.UROpts{
		DB: db.Get(),
	})
	userService := services.NewUserService(&services.USOpts{
		UserRepository: userRepository,
	})
	borrowingRepository := repository.NewBorrowingRepository(&repository.BRWOpts{
		DB: db.Get(),
	})
	borrowingService := services.NewBorrowingService(&services.BRWOpts{
		BorrowingRepository: borrowingRepository,
		BookRepository:      bookRepository,
	})
	authService := services.NewAuthService(&services.AuthSConfig{UserRepository: userRepository, Config: config.Config})
	return NewRouter(&RouterOpts{
		BookService:      bookService,
		UserService:      userService,
		BorrowingService: borrowingService,
		AuthService:      authService,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("error while running server", err)
		return
	}
}
