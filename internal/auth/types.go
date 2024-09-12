package auth

import "time"

type User struct {
	ID           uint64    `json:"id" field:"id"`
	Email        string    `json:"email" field:"email"`
	Password     string    `json:"-" field:"password"`
	PasswordSalt string    `json:"-" field:"password_salt"`
	CreateAt     time.Time `json:"-" field:"create_at"`
}
