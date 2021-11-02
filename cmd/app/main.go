package main

import (
	d "github.com/1makarov/go-crud-example/internal/db"
	"github.com/1makarov/go-crud-example/internal/db/postgres"
	v1 "github.com/1makarov/go-crud-example/internal/delivery/http/v1"
	"github.com/1makarov/go-crud-example/internal/repository"
	"github.com/1makarov/go-crud-example/internal/server"
	"github.com/1makarov/go-crud-example/internal/services"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	cfg := d.ConfigDB{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := postgres.Open(cfg)
	if err != nil {
		log.Fatalf("error open db: %s\n", err.Error())
	}
	defer db.Close()

	repo := repository.New(db)
	service := services.New(repo)
	handler := v1.NewHandler(service)

	s := server.NewServer(os.Getenv("PORT"), handler.Init())

	if err = s.Run(); err != nil {
		log.Fatalf("error running server: %s\n", err.Error())
	}
}