package Redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"main.go/internal/Config"
)

type Redis interface {
	Set(key string, value interface{}) error
	GetString(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	SetField(key string, value string) error
}
type DataBase struct {
	client *redis.Client
}

func New(cfg Config.Config) Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return DataBase{
		client: client,
	}
}
func (c DataBase) Set(key string, value interface{}) error {
	ctx := context.Background()
	err := c.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c DataBase) GetString(key string) (string, error) {
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
func (c DataBase) GetBytes(key string) ([]byte, error) {
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	return val, nil
}
func (c DataBase) SetField(key string, value string) error {
	ctx := context.Background()

	err := c.client.HSet(ctx, key, "FrequencyOfNotifications", value).Err()
	fmt.Println(err)
	fmt.Println("error in Hset")
	if err != nil {
		return err
	}
	return nil
}
