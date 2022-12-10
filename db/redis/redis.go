package redis

import (
	"github.com/go-redis/redis"
)

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", DB: 0})
	return client
}

func SetKey(data string) error {
	client := GetClient()
	err := client.Set(data, data, 0).Err()
	return err
}

func GetKey(data string) (string, error) {
	client := GetClient()
	val, err := client.Get(data).Result()
	return val, err
}
