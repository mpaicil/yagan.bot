package rediscfg

import (
	"context"
	"math/rand"
	"time"
)

// Names Some Non-Random name lists used to generate Random Users
var Names []string = []string{"Jasper", "Johan", "Edward", "Niel", "Percy", "Adam", "Grape", "Sam", "Redis", "Jennifer", "Jessica", "Angelica", "Amber", "Watch"}

// SirNames Some Non-Random name lists used to generate Random Users
var SirNames []string = []string{"Ericsson", "Redisson", "Edisson", "Tesla", "Bolmer", "Andersson", "Sword", "Fish", "Coder"}

// EmailProviders Some Non-Random email lists used to generate Random Users
var EmailProviders []string = []string{"Hotmail.com", "Gmail.com", "Awesomeness.com", "Redis.com"}

func examplePublish() {
	// Create a new Redis Client
	redisClient := initRedis()
	
	// Generate a new background context that  we will use
	ctx := context.Background()
	// Loop and randomly generate users on a random timer
	for {
		// Publish a generated user to the new_users channel
		err := redisClient.Publish(ctx, "market", GenerateRandomUser()).Err()
		if err != nil {
			panic(err)
		}
		// Sleep random time
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		time.Sleep(time.Duration(n) * time.Second)
	}

}

// GenerateRandomUser creates a random user, dont care too much about this.
func GenerateRandomUser() *User {
	rand.Seed(time.Now().UnixNano())
	nameMax := len(Names)
	sirNameMax := len(SirNames)
	emailProviderMax := len(EmailProviders)

	nameIndex := rand.Intn(nameMax-1) + 1
	sirNameIndex := rand.Intn(sirNameMax-1) + 1
	emailIndex := rand.Intn(emailProviderMax-1) + 1

	return &User{
		Username: Names[nameIndex] + " " + SirNames[sirNameIndex],
		Email:    Names[nameIndex] + SirNames[sirNameIndex] + "@" + EmailProviders[emailIndex],
	}
}