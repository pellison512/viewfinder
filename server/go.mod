module github.com/pellison512/viewfinder/server/main/v2

go 1.15

replace (
	github.com/pellison512/viewfinder/server/data/v2 => ./data
	github.com/pellison512/viewfinder/server/handlers/v2 => ./handlers
	github.com/pellison512/viewfinder/server/helpers/v2 => ./helpers
)

require (
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/pellison512/viewfinder/server/data/v2 v2.0.0-20201108085254-253204a8495d // indirect
	github.com/pellison512/viewfinder/server/handlers/v2 v2.0.0-20201108085254-253204a8495d
	github.com/pellison512/viewfinder/server/helpers/v2 v2.0.0-20201108085254-253204a8495d // indirect
)
