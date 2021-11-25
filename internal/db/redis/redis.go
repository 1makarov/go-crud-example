package redis

import (
	"github.com/1makarov/go-crud-example/internal/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func Open(cfg db.ConfigDB, salt string) (sessions.Store, error) {
	store, err := redis.NewStoreWithDB(10, "tcp", cfg.Address(), cfg.Password, cfg.Name, []byte(salt))
	if err != nil {
		return nil, err
	}

	return store, nil
}
