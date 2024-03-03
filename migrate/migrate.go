package main

import (
	"example/go-api-gin-gorm/initializers"
	"example/go-api-gin-gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
