package rediscfg

import (
	"context"
	"log"
	"strings" 
	"yagan.bot/telegram"
)

func Subscribe() {
	// Create a new Redis Client
	redisClient := initRedis()

	ctx := context.Background()
	// Subscribe to the Topic given
	topic := redisClient.Subscribe(ctx, "marketTopic")
	// Get the Channel to use
	channel := topic.Channel()
	// Itterate any messages sent on the channel
	for msg := range channel {

		log.Println(msg)
		
		var message = strings.Split(msg.Payload, "#")

		log.Println("Mensaje recibido: "+message[1])
		telegram.SendNotification(message[1])

	}
}

func utf8_decode(str string)string {    
    var result string
    for i := range str {
        result += string(str[i])
    }    
    return result
}

func stripCtlAndExtFromBytes(str string) string {
	b := make([]byte, len(str))
	var bl int
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 32 && c < 127 {
			b[bl] = c
			bl++
		}
	}
	return string(b[:bl])
}