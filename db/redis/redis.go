package redis

import (
	"github.com/go-redis/redis"
)

var (
	Address  = "localhost:6379"
	Password = "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"
)

type Database struct {
	Client *redis.Client
}

func NewRedisDatabase() (*Database, error) {
	client := redis.NewClient(&redis.Options{Addr: Address, Password: Password, DB: 0})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}
