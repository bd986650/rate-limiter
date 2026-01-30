package main

import (
	"log"
	"net/http"

	"github.com/bd986650/rate-limiter/api"
	"github.com/bd986650/rate-limiter/config"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})

	limiter := api.NewRateLimiter(rdb)
	producer := api.NewKafkaProducer()

	app := api.NewAPI(limiter, producer)

	http.HandleFunc("/ping", app.PingHandler)

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
