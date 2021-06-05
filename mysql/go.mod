module yagan.bot/mysql

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	yagan.bot/config v0.0.0-00010101000000-000000000000
)

replace yagan.bot/config => ../config
