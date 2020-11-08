module github.com/pellison512/viewfinder/server/handlers/v2

go 1.15

replace (
	github.com/pellison512/viewfinder/server/data/v2 => ../data
	github.com/pellison512/viewfinder/server/helpers/v2 => ../helpers
)

require (
	github.com/pellison512/viewfinder/server/data/v2 v2.0.0-20201108085254-253204a8495d
	github.com/pellison512/viewfinder/server/helpers/v2 v2.0.0-20201108083957-9e863abd7804
)
