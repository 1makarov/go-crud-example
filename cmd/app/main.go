package main

import (
	"context"
	"github.com/1makarov/go-crud-example/docs"
	"github.com/1makarov/go-crud-example/internal/db"
	"github.com/1makarov/go-crud-example/internal/db/postgres"
	"github.com/1makarov/go-crud-example/internal/delivery/http/v1"
	"github.com/1makarov/go-crud-example/internal/pkg/auth"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/server"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Library App API
// @version 1.0
// @description API Server for Library Application

// @BasePath /

// @securityDefinitions.apikey AuthKey
// @in header
// @name Authorization

func main() {
	docs.SwaggerInfo.Host = "localhost" + os.Getenv("APP_PORT")

	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := db.ConfigDB{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	d, err := postgres.Open(cfg)
	if err != nil {
		logrus.Fatalf("error open db: %s\n", err.Error())
	}

	a, err := auth.New(os.Getenv("JWT_SIGNING_KEY"), 5*time.Hour)
	if err != nil {
		logrus.Fatalf("error create auth: %s\n", err.Error())
	}

	repo := repository.New(d)
	service := services.New(repo, a)
	handler := v1.NewHandler(service)

	s := server.NewServer(os.Getenv("APP_PORT"), handler.Init())
	go func() {
		if err = s.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("LibraryApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = s.Shutdown(context.Background()); err != nil {
		logrus.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err = d.Close(); err != nil {
		logrus.Printf("error occured on db connection close: %s", err.Error())
	}
}
