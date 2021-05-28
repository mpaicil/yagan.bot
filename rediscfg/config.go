package rediscfg

import (
	"context"
	"fmt"
	"time"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"yagan.bot/config"
)

// User is a struct representing newly registered users
type User struct {
	Username string
	Email    string
}

// MarshalBinary encodes the struct into a binary blob
// Here I cheat and use regular json :)
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary decodes the struct into a User
func (u *User) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, u); err != nil {
		return err
	}
	return nil
}

func initRedis() *redis.Client {
	fmt.Println("Initialize Redis client")

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.LoadedConfig.Redis.Host + ":" + config.LoadedConfig.Redis.Port,
		//Addr:     "x.x.x.x:6379", // We connect to host redis, thats what the hostname of the redis service is set to in the docker-compose
		Password: "",                 // The password IF set in the redis Config file
		DB:       0,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		// Sleep for 3 seconds and wait for Redis to initialize
		time.Sleep(3 * time.Second)
		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	}

	return redisClient
}