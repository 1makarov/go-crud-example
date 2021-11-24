package main

import (
	"context"
	"github.com/1makarov/go-cache"
	"github.com/1makarov/go-crud-example/config"
	"github.com/1makarov/go-crud-example/internal/db/postgres"
	"github.com/1makarov/go-crud-example/internal/pkg/signaler"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/server"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/1makarov/go-crud-example/internal/transport/http"
	"github.com/sirupsen/logrus"
)

// @title Library App API
// @version 1.0
// @description API Server for Library Application

const configsDir = "configs"

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg, err := config.Init(configsDir)
	if err != nil {
		logrus.Fatalln(err)
	}

	db, err := postgres.Open(cfg.DB)
	if err != nil {
		logrus.Fatalf("error open db: %s\n", err.Error())
	}

	memCache := cache.NewWithInterval(cfg.CacheTTL)

	repo := repository.New(db)
	service := services.New(repo, memCache)
	handler := http.NewHandler(service)

	srv := server.NewServer(cfg.HTTP, handler.Init(cfg))
	go func() {
		if err = srv.Run(); err != nil {
			logrus.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("LibraryApp started")

	signaler.Wait()

	if err = srv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
