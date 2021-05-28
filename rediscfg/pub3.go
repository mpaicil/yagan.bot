package rediscfg

import (
	"context"
	"math/rand"
	"time"
)

func Publish() {
	// Create a new Redis Client
	redisClient := initRedis()
	
	// Generate a new background context that  we will use
	ctx := context.Background()
	// Loop and randomly generate users on a random timer
	for {
		// Publish a generated user to the new_users channel
		err := redisClient.Publish(ctx, "marketTopic", "Remensajeos!!!").Err()
		if err != nil {
			panic(err)
		}
		// Sleep random time
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		time.Sleep(time.Duration(n) * time.Second)
	}

}