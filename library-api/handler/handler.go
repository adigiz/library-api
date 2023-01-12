package handler

import (
	s "git.garena.com/sea-labs-id/trainers/library-api/services"
)

type Handler struct {
	bookService      s.BookService
	userService      s.UserService
	borrowingService s.BorrowingService
	authService      s.AuthService
}

type Opts struct {
	BookService      s.BookService
	UserService      s.UserService
	BorrowingService s.BorrowingService
	AuthService      s.AuthService
}

func New(opts *Opts) *Handler {
	return &Handler{
		bookService:      opts.BookService,
		userService:      opts.UserService,
		borrowingService: opts.BorrowingService,
		authService:      opts.AuthService,
	}
}
