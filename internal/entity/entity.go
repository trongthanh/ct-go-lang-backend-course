package entity

import (
	"time"
)

type User struct {
	Id             string `json:"_id,omitempty"`
	Email          string `json:"email"           bson:"email" unique:"true"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
	Active         bool   `json:"active"          bson:"active"`
}

type Profile struct {
	Id           string    `json:"_id,omitempty"`
	Userid       string    `json:"userid"        bson:"userid"`
	Bio          string    `json:"bio"           bson:"bio"`
	AccountType  string    `json:"account_type"  bson:"account_type"`
	Website      string    `json:"website"       bson:"website"`
	Name         string    `json:"name"          bson:"name"`
	Username     string    `json:"username"      bson:"username" unique:"true"`
	Gender       string    `json:"gender"        bson:"gender"`
	Birthday     time.Time `json:"birthday"      bson:"birthday"`
	CloseFriends []string  `json:"close_friends" bson:"close_friends"`
	Photo        string    `json:"photo"         bson:"photo"`
	Followers    []Profile `json:"followers"     bson:"followers"`
	Following    []Profile `json:"following"     bson:"following"`
	Requests     []Profile `json:"requests"      bson:"requests"`
}

type Image struct {
	Filename string `json:"filename" bson:"filename"`
	URL      string `json:"url"      bson:"url"`
	Path     string `json:"path"     bson:"path"`
	Size     int64  `json:"size"     bson:"size"`
}

type Post struct {
	Id        string    `json:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt"           bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"           bson:"updated_at"`
	Image     Image     `json:"image"               bson:"image"`
	Userid    string    `json:"userid"              bson:"userid"`
	Caption   string    `json:"caption"             bson:"caption"`
	Location  string    `json:"location"            bson:"location"`
	HashTags  []string  `json:"hashtags,nilasempty" bson:"hashtags"`
	Likes     []string  `json:"likes,nilasempty"    bson:"likes"`
	Comment   []Comment `json:"comment,nilasempty"  bson:"comment"`
}

type PostRes struct {
	Post
	Profile Profile `json:"profile"  bson:"profile"`
}

type Comment struct {
	Userid  string `json:"userid"    bson:"userid"`
	Comment string `json:"comment"   bson:"comment"`
}
