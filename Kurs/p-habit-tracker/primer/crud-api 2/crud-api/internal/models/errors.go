package models

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidInput       = errors.New("name, email and password are required")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrPasswordTooLong    = errors.New("password must be at most 72 bytes")
	ErrInvalidToken       = errors.New("invalid token")

	// post errors
	ErrPostAlreadyExists  = errors.New("post with this title already exists")
	ErrInvalidPostTitle   = errors.New("title is required and must be at most 70 characters")
	ErrInvalidPostContent = errors.New("content is required and must be at most 500 characters")
)
