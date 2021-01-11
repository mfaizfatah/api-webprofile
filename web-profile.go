package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mfaizfatah/api-webprofile/app/adapter"
	"github.com/mfaizfatah/api-webprofile/app/config"
	"github.com/mfaizfatah/api-webprofile/app/controllers"
	"github.com/mfaizfatah/api-webprofile/app/repository"
	"github.com/mfaizfatah/api-webprofile/app/routes"
	"github.com/mfaizfatah/api-webprofile/app/usecases"
)

func init() {
	service := "web-profile-api"

	config.LoadConfig(service)
}

func main() {
	db := adapter.DBSQL()

	repo := repository.NewRepo(db)
	uc := usecases.NewUC(repo)
	ctrl := controllers.NewCtrl(uc)

	router := routes.NewRouter(ctrl)
	router.Router(os.Getenv("SERVER_PORT"))
}
