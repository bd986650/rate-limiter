package main

import "log"

func ProcessMessage(key string, value string) {
	log.Printf("Processed message: %s = %s\n", key, value)
}
