module github.com/talbs1986/simplego/scenarios/job-app

go 1.21

replace github.com/talbs1986/simplego/goenv-configs => ../../goenv-configs

replace github.com/talbs1986/simplego/configs => ../../configs

replace github.com/talbs1986/simplego/metrics => ../../metrics

replace github.com/talbs1986/simplego/prom-metrics => ../../prom-metrics

replace github.com/talbs1986/simplego/app => ../../app

replace github.com/talbs1986/simplego/zerolog-logger => ../../zerolog-logger

require (
	github.com/talbs1986/simplego/app v0.0.0-20240819061751-864645887469
	github.com/talbs1986/simplego/configs v0.0.0-20240819061703-e4a65cc20bb3
	github.com/talbs1986/simplego/goenv-configs v0.0.0-20240819061751-864645887469
	github.com/talbs1986/simplego/metrics v0.0.0-20240820052917-4c1fbac69c95
	github.com/talbs1986/simplego/prom-metrics v0.0.0-20240820052917-4c1fbac69c95
	github.com/talbs1986/simplego/zerolog-logger v0.0.0-00010101000000-000000000000
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/sethvargo/go-envconfig v1.1.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
