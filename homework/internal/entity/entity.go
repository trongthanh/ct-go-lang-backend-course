package entity

import (
	"io"
)

type ImageInfo struct {
	FileName string `json:"file_name"`
	Path     string `json:"path"`
	URL      string `json:"url"`
	Username string `json:"username"`
}

type UserInfo struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=2,max=32"`
	FullName string `json:"full_name" validate:"required"`
	Address  string `json:"address"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type RegisterResponse struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
}

type SelfRequest struct {
	Username string `json:"username"`
}

type SelfResponse struct {
	Username string          `json:"username"`
	FullName string          `json:"full_name"`
	Address  string          `json:"address"`
	Images   []ImageResponse `json:"images"`
}

type ImageResponse struct {
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type FileInterface interface {
	io.Reader
}

type UploadImageRequest struct {
	Username string
	Filename string
	File     FileInterface
}

type UploadImageResponse struct {
	URL string `json:"url"`
}