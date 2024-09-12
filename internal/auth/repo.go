package auth

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (repo *Repo) GetUserByEmail(email string) (*User, error) {
	stmt := `select id, email, password, password_salt, create_at from users where email = $1`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user := User{}
	err := repo.DB.QueryRowContext(ctx, stmt, email).Scan(&user.ID, &user.Email, &user.Password, &user.PasswordSalt, &user.CreateAt)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user email not found")
			return nil, nil
		}
		log.Printf("error get user by email: %s", err.Error())
		return nil, err
	}

	return &user, nil
}

func (repo *Repo) InsertUser(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	stmt := `insert into users (email, password, password_salt, create_at)
		values ($1, $2, $3, $4) returning id`

	err := repo.DB.QueryRowContext(ctx, stmt, user.Email, user.Password, user.PasswordSalt, time.Now()).Scan(&user.ID)
	if err != nil {
		log.Printf("error inserting user: %s", err.Error())
	}
	return err
}
