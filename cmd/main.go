package main

import (
	"context"
	"github.com/1makarov/go-cache"
	"github.com/1makarov/go-crud-example/internal/config"
	"github.com/1makarov/go-crud-example/internal/db/postgres"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/server"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/1makarov/go-crud-example/internal/transport/http"
	"github.com/1makarov/go-crud-example/pkg/auth"
	"github.com/1makarov/go-crud-example/pkg/hash"
	"github.com/1makarov/go-crud-example/pkg/signaler"
	"github.com/sirupsen/logrus"
)

// @title Library App API
// @version 1.0
// @description API Server for Library Application

// @securityDefinitions.apikey AuthKey
// @in header
// @name Authorization

const configsDir = "configs"

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	ctx := context.Background()

	cfg, err := config.Init(configsDir)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	db, err := postgres.Open(cfg.DB)
	if err != nil {
		logrus.Errorln("error open db: %x\n", err)
		return
	}
	defer func() {
		if err = db.Close(); err != nil {
			logrus.Errorln(err)
		}
	}()

	authManager, err := auth.New(cfg.Auth.JWT.SigningKey, cfg.Auth.JWT.AccessTokenTTL)
	if err != nil {
		logrus.Errorln("error create auth: %x\n", err)
		return
	}

	hashManager := hash.New(cfg.Auth.PasswordSalt)
	memCache := cache.NewWithInterval(cfg.CacheTTL)

	repo := repository.New(db)
	service := services.New(repo, memCache, hashManager, authManager)
	handler := http.NewHandler(service).Init(cfg)

	srv := server.NewServer(cfg.HTTP, handler)
	defer func() {
		if err = srv.Stop(ctx); err != nil {
			logrus.Errorln(err)
		}
	}()

	go func() {
		if err = srv.Run(); err != nil {
			logrus.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("LibraryApp started")

	signaler.Wait()
}
