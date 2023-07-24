package entity

import "mime/multipart"

type ImageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"url"`
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
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type UploadImageRequest struct {
	File *multipart.FileHeader
}

type UploadImageResponse struct {
	URL string `json:"url"`
}
