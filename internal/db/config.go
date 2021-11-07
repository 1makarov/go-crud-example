package db

import "fmt"

type ConfigDB struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func (cfg *ConfigDB) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.DBName,
	)
}
