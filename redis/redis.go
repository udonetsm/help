package redis

import (
	"github.com/go-redis/redis"
	"github.com/udonetsm/help/helper"
	"github.com/udonetsm/help/models"
)

func NewStore(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

func NewOpts(username, password string) *redis.Options {
	return &redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	}
}

func SetVals(key, value string, auth models.Auth) error {
	store := NewStore(NewOpts(auth.Email, auth.Password))
	err := store.Set(store.Context(), key, value, 0).Err()
	helper.Errors(err, "set(setvals)")
	return err
}

func GetVals(key string, auth models.Auth) interface{} {
	store := NewStore(NewOpts(auth.Email, auth.Password))
	vals, err := store.Get(store.Context(), key).Result()
	helper.Errors(err, "get(getvals)")
	return vals
}
