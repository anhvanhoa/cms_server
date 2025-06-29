package bootstrap

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr        string
	Password    string
	DB          int
	Network     string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type RedisIntance struct {
	config RedisConfig
	client *redis.Client
	ctx    context.Context
}

type RedisConfigImpl interface {
	Set(key string, value any, time time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	Ping() error
}

func NewRedis(c RedisConfig) RedisConfigImpl {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.DB,
		Network:      c.Network,
		PoolSize:     c.MaxActive,
		MinIdleConns: c.MaxIdle,
		PoolTimeout:  time.Duration(c.IdleTimeout) * time.Second,
	})

	ri := &RedisIntance{
		config: c,
		client: client,
		ctx:    ctx,
	}
	_, err := ri.client.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return ri
}

func (ri *RedisIntance) Set(key string, value any, time time.Duration) error {
	err := ri.client.Set(ri.ctx, key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}

func (ri *RedisIntance) Get(key string) (string, error) {
	value, err := ri.client.Get(ri.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (ri *RedisIntance) Del(key string) error {
	err := ri.client.Del(ri.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (ri *RedisIntance) Ping() error {
	_, err := ri.client.Ping(ri.ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func NewRedisConfig(
	addr, password string, db int, network string,
	maxIde, maxActive, idleTimeout int,
) RedisConfig {
	return RedisConfig{
		Addr:        addr,
		Password:    password,
		DB:          db,
		Network:     network,
		MaxIdle:     maxIde,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}
}
