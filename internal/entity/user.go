package entity

import (
	"time"
)

type User struct {
	Id   string
	Username string
	Password string
	Created_at time.Time
	Updated_at time.Time
}

func (u User) GetID() string {
	return u.Id
}

func (u User) GetName() string {
	return u.Username
}

func (u User) GetPassword() string {
	return u.Password
}
