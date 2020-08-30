module awesomeProject/cdr/main

require (
	awesomeProject/cdr/cdr v0.0.0
	awesomeProject/cdr/tools v0.0.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.60.2
)

replace awesomeProject/cdr/tools => ../tools

replace awesomeProject/cdr/cdr => ../cdr

go 1.13
