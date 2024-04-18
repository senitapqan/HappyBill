package main

import (
	happyBill "happyBill"

	"happyBill/pkg/handler"
	"happyBill/pkg/repository"
	"happyBill/pkg/service"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@title			HappyBill
//	@version		1.0
//	@description	This is a server for order billboards.

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							headers
//	@name						Authorization

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("some error with initializiing: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		log.Fatalf("error with data base: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(happyBill.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error with run server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
