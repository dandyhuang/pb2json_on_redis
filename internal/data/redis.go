package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)
func CreateRedis()  *redis.ClusterClient{
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: viper.GetStringSlice("redis"),
		//MaxRedirects: c.config.MaxRetries,
		//ReadOnly:     c.config.ReadOnly,
		//Password:     c.config.Password,
		//MaxRetries:   c.config.MaxRetries,
		//DialTimeout:  c.config.DialTimeout,
		//ReadTimeout:  c.config.ReadTimeout,
		//WriteTimeout: c.config.WriteTimeout,
		//PoolSize:     c.config.PoolSize,
		//MinIdleConns: c.config.MinIdleConns,
		//IdleTimeout:  c.config.IdleTimeout,
	})
	if err := clusterClient.Ping(context.Background()).Err(); err != nil {
			log.Println("start cluster redis", err)
	}
	return clusterClient
}