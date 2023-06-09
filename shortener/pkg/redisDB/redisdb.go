package redisDB

import (
	"context"
	"github.com/go-redis/redis/v8"
	"grpcService/config"
	"grpcService/pkg/logger"
	"strconv"
)

func GetClient(ctx context.Context, c config.Config) (*redis.Client, error) {
	log := logger.GetLogger()
	dbNum, _ := strconv.Atoi(c.RedisDB.DBNum)
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisDB.Addr, // адрес сервера Redis
		Password: "",             // пароль, если он установлен
		DB:       dbNum,          // номер базы данных Redis
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal().Err(err).Msg("Redis is unable to connect.")
		return nil, err
	}
	log.Info().Msg("Redis connected successfully.")
	return rdb, nil
}
