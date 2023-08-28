package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	segment "testingavito"
	"testingavito/pkg/handler"
	"testingavito/pkg/repository"
	"testingavito/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title TestingAvito Slug App API
// @version 1.0
// @description API Server for TestingAvito Slug Application

// @host localhost:8080
// @BasePath /

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error intializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewDB(repository.Config{
		Host:   viper.GetString("db.Host"),
		Port:   viper.GetString("db.Port"),
		Uname:  viper.GetString("db.Uname"),
		Pass:   os.Getenv("DB_PASS"),
		NameDB: viper.GetString("db.NameDB"),
		SSL:    viper.GetString("db.SSL"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(segment.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Slug App Has Been Activated")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Slug App Has Been Downed")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("an error occurred while shutting down the server:%s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("an error occurred while closing the database connection: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
