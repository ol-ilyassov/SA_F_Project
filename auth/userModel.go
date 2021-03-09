package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

var m = &UserModel{}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := "INSERT INTO users (name, email, hashed_password, created) VALUES($1, $2, $3, $4)"

	_, err = m.DB.Exec(context.Background(), stmt, name, email, string(hashedPassword), time.Now())
	if err != nil {
		if strings.Contains(err.Error(), "users_uc_email") {
			return ErrDuplicateEmail
		}
		return err
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, hashed_password FROM users WHERE email = $1 AND active = TRUE"
	row := m.DB.QueryRow(context.Background(), stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, ErrNoRecord
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	u := &User{}
	stmt := `SELECT id, name, email, created, active FROM users WHERE id = $1`
	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&u.ID, &u.Name, &u.Email, &u.Created, &u.Active)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return u, nil
}
