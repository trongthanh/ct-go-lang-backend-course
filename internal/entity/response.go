package entity

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type RegisterResponse struct {
	Id string `json:"id"`
}

// Remove below once migrated
type LoginResponse struct {
	Token string `json:"token"`
}

type SelfResponse struct {
	Email    string          `json:"email"`
	FullName string          `json:"full_name"`
	Address  string          `json:"address"`
	Images   []ImageResponse `json:"images"`
}

type ImageResponse struct {
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type UploadImageResponse struct {
	URL string `json:"url"`
}

type ChangePasswordResponse struct {
	Success bool `json:"success"`
}
