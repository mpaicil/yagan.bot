package main

import (
	"yagan.bot/config"
	"yagan.bot/rediscfg"
	"yagan.bot/mysql"
	"yagan.bot/plot"
)

func main() {
	config.LoadConfig()
	mysql.Connection()
	
	go plot.LastHourPlot()

	rediscfg.Subscribe()
}
