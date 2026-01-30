package api

import (
	"fmt"
	"net/http"
)

type API struct {
	limiter  *RateLimiter
	producer *KafkaProducer
}

func NewAPI(limiter *RateLimiter, producer *KafkaProducer) *API {
	return &API{limiter: limiter, producer: producer}
}

func (a *API) PingHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	if userID == "" {
		userID = "anonymous"
	}
	allowed, err := a.limiter.Allow(userID)
	if err != nil {
		http.Error(w, "Redis error", 500)
		return
	}
	if !allowed {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Rate limit exceeded"))
		return
	}
	a.producer.Send(userID, fmt.Sprintf("user %s called /ping", userID))
	w.Write([]byte("pong"))
}
