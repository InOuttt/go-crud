package entity

import (
	"time"
)

// Profile represents an profile record.
type Profile struct {
	Id        string    `json:"id"`
	User_id    string    `json:"userId"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}