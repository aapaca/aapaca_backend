module github.com/aapaca/aapaca_backend/src/app

go 1.14

replace domain => ./domain

replace infrastructure => ./infrastructure

replace interfaces => ./interfaces

replace usecase => ./usecase

require (
	domain v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	infrastructure v0.0.0-00010101000000-000000000000
	interfaces v0.0.0-00010101000000-000000000000 // indirect
	usecase v0.0.0-00010101000000-000000000000 // indirect
)
