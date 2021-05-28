package rediscfg

import (
	"context"
	"fmt"
)

func exampleSubscribe() {
	// Create a new Redis Client
	redisClient := initRedis()

	ctx := context.Background()
	// Subscribe to the Topic given
	topic := redisClient.Subscribe(ctx, "market")
	// Get the Channel to use
	channel := topic.Channel()
	// Itterate any messages sent on the channel
	for msg := range channel {
		u := &User{}
		// Unmarshal the data into the user
		err := u.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			panic(err)
		}

		fmt.Println(u)
	}
}

func (u *User) String() string {
	return "User: " + u.Username + " registered with Email: " + u.Email
}