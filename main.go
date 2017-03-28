package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var (
	Redis = initRedis()
)

func initRedis() *redis.Client {
	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		redisUrl = "127.0.0.1:6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})
}

func main() {
	act := flag.String("act", "producer", "Either: producer or consumer")
	partition := flag.String("partition", "0",
		"Partition which the consumer program will be subscribing")

	flag.Parse()

	fmt.Printf("Welcome to Banku service: %s\n\n", *act)

	switch *act {
	case "producer":
		mainProducer()
	case "consumer":
		if part32int, err := strconv.ParseInt(*partition, 10, 32); err == nil {
			mainConsumer(int32(part32int))
		}
	}
}
