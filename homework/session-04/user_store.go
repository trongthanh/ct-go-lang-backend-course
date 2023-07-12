package main

import (
	"fmt"
	"sync"
)

var userStore = NewUserStore()

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

func (u *UserStore) Save(info UserInfo) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, existed := u.data[info.Username]; existed {
		return fmt.Errorf("User %s existed", info.Username)
	}

	u.data[info.Username] = info
	fmt.Println("UserStore.Save", info)
	return nil
}

func (*UserStore) Get(username string) (UserInfo, error) {
	u, ok := userStore.data[username]
	if !ok {
		// fmt.Println("UserStore.Get", username, "not found")
		return UserInfo{}, fmt.Errorf("User not found")
	}
	// fmt.Println("UserStore.Get", username, u)
	return u, nil
}

type UserInfo struct {
	Username string `json:"username" validate:"required,min=2,max=32"`
	FullName string `json:"full_name" validate:"required"`
	Address  string `json:"address"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
