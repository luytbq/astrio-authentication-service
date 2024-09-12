package auth

import "time"

type User struct {
	ID           uint64 `json:",omitempty"`
	Email        string
	Password     string    `json:"-"`
	CreateAt     time.Time `json:",omitempty"`
	PasswordSalt string    `json:"-"`
}
