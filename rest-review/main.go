package main

import (
	"fmt"
	"net/http"
	"rest-review/controllers"
	"rest-review/middleware"
	"rest-review/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware())

	// load env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// connect db
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersHandler := controllers.Users{
		UserService: &userService,
	}

	r.POST("/v1/ms-paylater/register", usersHandler.HandleRegister)
	r.POST("/v1/ms-paylater/login", usersHandler.HandleLogin)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)

}
