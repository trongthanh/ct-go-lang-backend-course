package entity

import "io"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SelfRequest struct {
	Userid string `json:"userid"`
}

type FileInterface interface {
	io.Reader
}

type UploadImageRequest struct {
	Userid   string
	Filename string
	File     FileInterface
}

type ChangePasswordRequest struct {
	Userid          string `json:"userid"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=32"`
	RepeatPassword  string `json:"repeat_password"`
}
