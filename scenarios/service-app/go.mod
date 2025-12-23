module github.com/talbs1986/simplego/scenarios/service-app

go 1.24

replace github.com/talbs1986/simplego/goenv-configs => ../../goenv-configs

replace github.com/talbs1986/simplego/configs => ../../configs

replace github.com/talbs1986/simplego/metrics => ../../metrics

replace github.com/talbs1986/simplego/prom-metrics => ../../prom-metrics

replace github.com/talbs1986/simplego/app => ../../app

replace github.com/talbs1986/simplego/zerolog-logger => ../../zerolog-logger

replace github.com/talbs1986/simplego/server => ../../server

replace github.com/talbs1986/simplego/chi-server => ../../chi-server

require (
	github.com/go-chi/chi/v5 v5.2.3
	github.com/talbs1986/simplego/app v0.0.0-20240819061751-864645887469
	github.com/talbs1986/simplego/chi-server v0.0.0-00010101000000-000000000000
	github.com/talbs1986/simplego/configs v0.0.0-20240819061703-e4a65cc20bb3
	github.com/talbs1986/simplego/goenv-configs v0.0.0-20240819061751-864645887469
	github.com/talbs1986/simplego/metrics v0.0.0-20240820052917-4c1fbac69c95
	github.com/talbs1986/simplego/prom-metrics v0.0.0-20240820052917-4c1fbac69c95
	github.com/talbs1986/simplego/server v0.0.0-20240528101415-c854be60989c
	github.com/talbs1986/simplego/zerolog-logger v0.0.0-00010101000000-000000000000
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.66.1 // indirect
	github.com/prometheus/procfs v0.16.1 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	github.com/sethvargo/go-envconfig v1.3.0 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	golang.org/x/sys v0.35.0 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)
