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

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}

type ProfilesResponse struct {
	Profiles []Profile `json:"profiles"`
}

type PostsResponse struct {
	Posts []PostRes `json:"posts"`
}

type CreatePostResponse struct {
	Postid string `json:"postid"`
}

type LikePostResponse struct {
	LikesTotal int `json:"likes_total"`
}

type DeletePostResponse struct {
	Postid string `json:"postid"`
}
