package entity

import "time"

type User struct {
	Email          string `json:"email"           bson:"email" unique:"true"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
	Active         bool   `json:"active"          bson:"active"`
}

type Profile struct {
	Userid       string    `json:"userid"        bson:"userid"`
	Bio          string    `json:"bio"           bson:"bio"`
	AccountType  string    `json:"account_type"  bson:"account_type"`
	Website      string    `json:"website"       bson:"website"`
	Name         string    `json:"name"          bson:"name"`
	Username     string    `json:"username"      bson:"username"`
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
	URL      string `json:"url"  bson:"url"`
	Path     string `json:"path" bson:"path"`
}

type Post struct {
	Profile  Profile   `json:"profile"  bson:"profile"`
	Caption  string    `json:"caption"  bson:"caption"`
	Location string    `json:"location" bson:"location"`
	HashTags []string  `json:"hashtags" bson:"hashtags"`
	Likes    []Profile `json:"likes"    bson:"likes"`
	Image    []Image   `json:"image"    bson:"image"`
	Comment  []Comment `json:"comment"  bson:"comment"`
}

type Comment struct {
	Commenter Profile `json:"commenter" bson:"commenter"`
	Comment   string  `json:"comment"   bson:"comment"`
}
