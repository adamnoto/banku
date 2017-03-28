package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

var (
	Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
)

func main() {
	act := flag.String("act", "producer", "Either: producer or consumer")
	partition := flag.String("partition", "0",
		"Partition which the consumer program will be subscribing")

	flag.Parse()

	fmt.Printf("Welcome to Banku service: %s\n\n", *act)

	switch *act {
	case "producer":
		// command part
		mainProducer()
	case "consumer":
		// query part
		if part32int, err := strconv.ParseInt(*partition, 10, 32); err == nil {
			mainConsumer(int32(part32int))
		}
	}
}
