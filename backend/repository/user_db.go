package repository

import (
	"database/sql"
	"fmt"
)

type userRepositoryDB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userRepositoryDB { //* Constructor
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(email string, password string, secret string) (*User, error) {
	// Insert document into database
	insert, err := r.db.Exec("INSERT INTO users (email, password, secret) VALUES (?, ?, ?)", email, password, secret)
	if err != nil {
		return nil, err
	}
	userId, err := insert.LastInsertId()

	// Create user object
	var user = User{
		Id:       userId,
		Email:    email,
		Password: password,
		Secret:   secret,
	}
	return &user, nil
}

func (r userRepositoryDB) CheckUser(email string) (*User, error) {
	// implement me
	fmt.Print("User Fetched")
	return nil, nil
}

func (r userRepositoryDB) GetUsers() ([]*User, error) {
	// implement me
	fmt.Print("Users Fetched")
	return nil, nil
}
