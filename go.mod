module hub

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe
	github.com/jinzhu/gorm v1.9.12
)

replace golang.org/x/net v0.0.0-20200202094626-16171245cfb2 => github.com/golang/net v0.0.0-20200202094626-16171245cfb2