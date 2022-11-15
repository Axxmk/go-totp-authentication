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
	// implement me
	fmt.Print("User Inserted")
	return nil, nil
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
