package main

import (
	"yagan.bot/config"
	"yagan.bot/rediscfg"
	"yagan.bot/mysql"
)

func main() {
	config.LoadConfig()
	mysql.Connection()
	rediscfg.Subscribe()
}
