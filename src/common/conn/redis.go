package commonConn

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func RedisConnect() (*redis.Client, error) {
	rOpts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%v", os.Getenv("RDB_HOST"), os.Getenv("RDB_PORT")),
		Password: os.Getenv("RBD_PASSWORD"),
		DB:       0,
	}

	rdb := redis.NewClient(rOpts)
	res := rdb.Ping(context.Background())
	if res.Err() != nil {
		return nil, res.Err()
	}
	return rdb, nil
}
