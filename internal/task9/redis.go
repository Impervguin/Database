package task9

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	Client *redis.Client
}

func NewRedisStorage() (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	stat := client.Ping(context.TODO())
	if stat.Err() != nil {
		return nil, stat.Err()
	}
	return &RedisStorage{Client: client}, nil
}

const WealthClientExpire = time.Second * 30

func (s *RedisStorage) PutTopWealthyClients(wc []WealthyClient) error {
	data, err := json.Marshal(wc)
	if err != nil {
		return err
	}
	err = s.Client.Set(context.TODO(), "top_clients", data, WealthClientExpire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisStorage) GetTopWealthyClients() ([]WealthyClient, error) {
	wc := make([]WealthyClient, 0)
	data, err := s.Client.Get(context.TODO(), "top_clients").Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	s.Client.Expire(context.TODO(), "top_clients", WealthClientExpire)
	err = json.Unmarshal(data, &wc)
	if err != nil {
		return nil, err
	}
	return wc, nil
}

func (s *RedisStorage) DropTopWealthyClients() error {
	err := s.Client.Del(context.TODO(), "top_clients").Err()
	if err != nil {
		return err
	}
	return nil
}
