package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var ctx = context.Background()
var client *redis.Client

func setKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var keyVal map[string]string
	if err := json.NewDecoder(r.Body).Decode(&keyVal); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for key, val := range keyVal {
		err := client.Set(ctx, key, val, 0).Err()
		if err != nil {
			http.Error(w, "Failed to set key-value pair", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func getKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Failed to get value for key", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Value for key %s: %s", key, val)
}

func delKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var keys map[string]string
	if err := json.NewDecoder(r.Body).Decode(&keys); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for key := range keys {
		err := client.Del(ctx, key).Err()
		if err != nil {
			http.Error(w, "Failed to delete key", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Forbidden", http.StatusForbidden)
}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	http.HandleFunc("/set_key", setKeyHandler)
	http.HandleFunc("/get_key", getKeyHandler)
	http.HandleFunc("/del_key", delKeyHandler)
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
