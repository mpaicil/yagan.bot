module yagan.bot/plot

go 1.16

replace yagan.bot/mysql => ../mysql

replace yagan.bot/config => ../config

require (
	gonum.org/v1/plot v0.9.0
	yagan.bot/mysql v0.0.0-00010101000000-000000000000
	yagan.bot/telegram v0.0.0-00010101000000-000000000000
)

replace yagan.bot/telegram => ../telegram
