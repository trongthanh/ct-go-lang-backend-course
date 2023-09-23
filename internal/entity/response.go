package entity

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Response structs below are data field in Response base
type RegisterResponse struct {
	Id string `json:"id"`
}

type LoginResponse struct {
	Token string `json:"token"`
	// TODO: check front end if profile is needed
	Profile Profile `json:"profile"`
}

// Below not migrated
type SelfResponse struct {
	Profile Profile `json:"profile"`
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
