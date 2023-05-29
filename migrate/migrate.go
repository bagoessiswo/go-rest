package main

import (
	"github.com/bagoessiswo/go-rest/initializers"
	"github.com/bagoessiswo/go-rest/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
