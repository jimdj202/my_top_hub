module hub

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe
	github.com/jinzhu/gorm v1.9.12
	golang.org/x/text v0.3.0
)

replace golang.org/x/net v0.0.0-20200202094626-16171245cfb2 => github.com/golang/net v0.0.0-20200202094626-16171245cfb2

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200323165209-0ec3e9974c59

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20200327173247-9dae0f8f5775
