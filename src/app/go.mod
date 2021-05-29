module github.com/aapaca/aapaca_backend/src/app

go 1.14

replace domain => ./domain

replace infrastructure => ./infrastructure

replace interfaces => ./interfaces

replace test => ./test

replace usecases => ./usecases

require (
	domain v0.0.0-00010101000000-000000000000 // indirect
	github.com/aapaca/aapaca_backend/src/app/interfaces v0.0.0-20210110064323-cbd706943604 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	infrastructure v0.0.0-00010101000000-000000000000
	interfaces v0.0.0-00010101000000-000000000000
	test v0.0.0-00010101000000-000000000000 // indirect
	usecases v0.0.0-00010101000000-000000000000 // indirect
)
