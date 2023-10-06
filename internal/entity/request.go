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

type ProfilesRequest struct {
	// empty for now
}

type PostsRequest struct {
	// empty
}

type PostsByUserRequest struct {
	Userid string `json:"userid"`
}

type CreatePostRequest struct {
	Post     Post
	Filename string
	File     io.Reader
}

type DeletePostRequest struct {
	Postid string `json:"postid"`
}

type LikePostRequest struct {
	Postid string `json:"postid"`
	Userid string `json:"userid"`
}

// type FileInterface interface {
// 	io.Reader
// }
//
// type UploadImageRequest struct {
// 	Userid   string
// 	Filename string
// 	File     FileInterface
// }
