module github.com/pellison512/viewfinder/server/main/v2

go 1.15

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/pellison512/viewfinder/server/handlers/v2 v2.0.0-20201108044217-62fd3330887d
	github.com/pellison512/viewfinder/server/helpers/v2 v2.0.0-00010101000000-000000000000 // indirect
)

replace (
	github.com/pellison512/viewfinder/server/handlers/v2 => ./handlers
	github.com/pellison512/viewfinder/server/helpers/v2 => ./helpers
)
