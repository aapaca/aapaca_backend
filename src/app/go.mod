module github.com/aapaca/aapaca_backend/src/app

go 1.14

replace domain => ./domain

replace infrastructure => ./infrastructure

replace interfaces => ./interfaces

replace test => ./test

replace usecases => ./usecases

require (
	domain v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.7.0 // indirect
	infrastructure v0.0.0-00010101000000-000000000000
	interfaces v0.0.0-00010101000000-000000000000 // indirect
	test v0.0.0-00010101000000-000000000000 // indirect
	usecases v0.0.0-00010101000000-000000000000 // indirect
)
