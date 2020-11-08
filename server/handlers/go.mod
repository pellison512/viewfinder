module github.com/pellison512/viewfinder/server/handlers/v2

go 1.15

replace (
	github.com/pellison512/viewfinder/server/data/v2 => ../data
	github.com/pellison512/viewfinder/server/helpers/v2 => ../helpers
)

require (
	github.com/pellison512/viewfinder/server/data/v2 v2.0.0-00010101000000-000000000000
	github.com/pellison512/viewfinder/server/helpers/v2 v2.0.0-00010101000000-000000000000
)
