package redis

import (
	"github.com/go-redis/redis"
	"github.com/udonetsm/help/helper"
)

func NewStore(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

func NewOpts(username, password string, db int) *redis.Options {
	return &redis.Options{
		Addr:     ":6379",
		Password: password,
		Username: username,
		DB:       db,
	}
}

func SetVals(key, value, dbusername, dbpassword string, db int) error {
	store := NewStore(NewOpts(dbusername, dbpassword, db))
	err := store.Set(store.Context(), key, value, 0).Err()
	helper.Errors(err, "set(setvals)")
	return err
}

func GetVals(key, dbusername, dbpassword string, db int) interface{} {
	store := NewStore(NewOpts(dbusername, dbpassword, db))
	vals, err := store.Get(store.Context(), key).Result()
	helper.Errors(err, "get(getvals)")
	return vals
}
