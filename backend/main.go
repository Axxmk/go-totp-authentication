package main

import (
	"database/sql"
	"github.com/axxmk/go-totp-authentication/config"
	"github.com/axxmk/go-totp-authentication/handler"
	"github.com/axxmk/go-totp-authentication/repository"
	"github.com/axxmk/go-totp-authentication/serivce"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	db, err := sql.Open("mysql", config.C.DB_HOST)
	if err != nil {
		panic(err)
	}
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/signup", userHandler.SignUp)
	http.HandleFunc("/signin", userHandler.SignIn)
	http.HandleFunc("/listuser", userHandler.ListUsers)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

	defer db.Close()
}
