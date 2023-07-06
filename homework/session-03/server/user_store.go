package main

import (
	"fmt"
	"sync"
)

// TODO #1: implement in-memory user store ✅

var userStore = NewUserStore()

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

func (u *UserStore) Save(info UserInfo) error {
	if _, existed := u.data[info.Username]; existed {
		return fmt.Errorf("user existed")
	}

	u.data[info.Username] = info
	fmt.Println("UserStore.Save", info)
	return nil
}

func (*UserStore) Get(username string) (UserInfo, error) {
	u, ok := userStore.data[username]
	if !ok {
		// fmt.Println("UserStore.Get", username, "not found")
		return UserInfo{}, fmt.Errorf("user not found")
	}
	// fmt.Println("UserStore.Get", username, u)
	return u, nil
}

type UserInfo struct {
	// Id         int
	Username   string
	FullName   string
	Address    string
	Password   string
}

