package main

import (
	"yagan.bot/config"
	"yagan.bot/rediscfg"
	"yagan.bot/mysql"
	"yagan.bot/plot"
	"yagan.bot/mindicator"
)

func main() {
	config.LoadConfig()
	go mysql.Connection()
	
	go plot.LastHourPlot()
	go mindicator.PrintIndicators()

	rediscfg.Subscribe()
}
