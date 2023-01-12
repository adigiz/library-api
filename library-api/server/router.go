package server

import (
	"git.garena.com/sea-labs-id/trainers/library-api/dto"
	"git.garena.com/sea-labs-id/trainers/library-api/handler"
	"git.garena.com/sea-labs-id/trainers/library-api/middlewares"
	"git.garena.com/sea-labs-id/trainers/library-api/services"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	BookService      services.BookService
	UserService      services.UserService
	BorrowingService services.BorrowingService
	AuthService      services.AuthService
}

func NewRouter(opts *RouterOpts) *gin.Engine {
	router := gin.Default() // internally using logger and recovery mdw

	h := handler.New(&handler.Opts{
		BookService:      opts.BookService,
		UserService:      opts.UserService,
		BorrowingService: opts.BorrowingService,
		AuthService:      opts.AuthService,
	})

	router.Static("/docs", "swaggerui")

	router.Use(middlewares.ErrorHandler) // global error middleware
	router.POST("/signin", middlewares.RequestMiddleware(&dto.SignInReq{}), h.SignIn)

	router.Use(middlewares.AuthorizeJWT)
	books := router.Group("/books")
	{
		books.GET("", h.GetBooks)
		books.POST("", middlewares.RequestMiddleware(&dto.BookReq{}), h.PostBook)
	}

	router.GET("/users", h.GetUsers)

	records := router.Group("/borrowing-records")
	{
		records.POST("", middlewares.RequestMiddleware(&dto.BorrowingReq{}), h.PostBorrowingRecord)
		records.PATCH("/:id", middlewares.RequestMiddleware(&dto.PatchBorrowingReq{}), h.PatchBorrowingRecord)
	}

	return router
}
