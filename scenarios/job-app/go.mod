module github.com/talbs1986/simplego/scenarios/job-app

go 1.20

replace github.com/talbs1986/simplego/goenv-configs => ../../goenv-configs

replace github.com/talbs1986/simplego/configs => ../../configs

replace github.com/talbs1986/simplego/prom-metrics => ../../prom-metrics

replace github.com/talbs1986/simplego/app => ../../app

replace github.com/talbs1986/simplego/zerolog-logger => ../../zerolog-logger

require (
	github.com/talbs1986/simplego/app v0.0.0-20240819061751-864645887469
	github.com/talbs1986/simplego/configs v0.0.0-20240819061703-e4a65cc20bb3
	github.com/talbs1986/simplego/goenv-configs v0.0.0-20240819061751-864645887469
	github.com/talbs1986/simplego/zerolog-logger v0.0.0-00010101000000-000000000000
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/sethvargo/go-envconfig v1.1.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
