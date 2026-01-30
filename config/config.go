package config

import "time"

var (
	RedisAddr    = "redis:6379"
	KafkaBrokers = []string{"kafka:9092"}
	KafkaTopic   = "requests"
	RateLimit    = 5                // 5 запросов
	RateWindow   = 60 * time.Second // окно времени 60 секунд
)
