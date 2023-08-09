package userstore

import (
	"errors"
	"fmt"
	"sync"
	"thanhtran/internal/entity"
)

// var Store = NewUserStore()

func New() *UserStore {
	return &UserStore{data: make(map[string]entity.UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]entity.UserInfo
}

func (u *UserStore) Save(info entity.UserInfo) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, existed := u.data[info.Username]; existed {
		return fmt.Errorf("User %s existed", info.Username)
	}

	u.data[info.Username] = info
	fmt.Println("UserStore.Save", info)
	return nil
}

func (u *UserStore) Get(username string) (entity.UserInfo, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, ok := u.data[username]
	if !ok {
		// fmt.Println("UserStore.Get", username, "not found")
		return entity.UserInfo{}, ErrUserNotFound
	}
	// fmt.Println("UserStore.Get", username, u)
	return user, nil
}

var ErrUserNotFound = errors.New("user not found")
