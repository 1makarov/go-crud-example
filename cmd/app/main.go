package main

import (
	"context"
	"fmt"
	"github.com/1makarov/go-cache"
	"github.com/1makarov/go-crud-example/docs"
	database "github.com/1makarov/go-crud-example/internal/db"
	"github.com/1makarov/go-crud-example/internal/db/postgres"
	"github.com/1makarov/go-crud-example/internal/delivery/http/v1"
	"github.com/1makarov/go-crud-example/internal/pkg/auth"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/server"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Library App API
// @version 1.0
// @description API Server for Library Application

// @securityDefinitions.apikey AuthKey
// @in header
// @name Authorization

func main() {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT"))

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s\n", err.Error())
	}

	cfg := database.ConfigDB{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := postgres.Open(cfg)
	if err != nil {
		logrus.Fatalf("error open db: %s\n", err.Error())
	}

	authManager, err := auth.New(os.Getenv("JWT_SIGNING_KEY"), viper.GetDuration("auth.ttl"))
	if err != nil {
		logrus.Fatalf("error create auth: %s\n", err.Error())
	}

	memCache := cache.NewWithInterval(viper.GetDuration("cache.ttl"))

	repo := repository.New(db)
	service := services.New(repo, memCache)
	handler := v1.NewHandler(service, authManager)

	srv := server.NewServer(os.Getenv("APP_PORT"), handler.Init())
	go func() {
		if err = srv.Run(); err != nil {
			logrus.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("LibraryApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
