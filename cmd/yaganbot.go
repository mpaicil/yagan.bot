package main

import (
	"yagan.bot/config"
	"yagan.bot/rediscfg"
)

func main() {
	config.LoadConfig()
	rediscfg.Subscribe()
}
