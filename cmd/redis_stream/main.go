package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
	"github.com/tzota/nostromo-parker/internal/config"
	"github.com/tzota/nostromo-parker/internal/harvester"
	"github.com/tzota/nostromo-parker/pkg/harvest"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// pong, err := client.XRead(ctx, &redis.XReadArgs{Streams: []string{"mystream", "$"}, Block: 0, Count: 1}).Result()

	cfg, err := config.ReadFromFile("./config.json")
	if err != nil {
		panic(err)
	}

	log.Info("Press Ctrl-C to stop")

	harvest.Simple(cfg, func(message harvester.IMessage) {
		if err := message.ReportToRedisStream(client, "nostromo-brett"); err != nil {
			log.Error(err)
		}
	})

	select {}
}
