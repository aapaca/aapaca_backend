module github.com/aapaca/aapaca_backend/src/app

go 1.14

replace domain => ./domain

replace infrastructure => ./infrastructure

replace interfaces => ./interfaces

replace usecase => ./usecase

require (
	domain v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	infrastructure v0.0.0-00010101000000-000000000000
	interfaces v0.0.0-00010101000000-000000000000 // indirect
	usecase v0.0.0-00010101000000-000000000000 // indirect
)
