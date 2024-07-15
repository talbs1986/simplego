module github.com/talbs1986/simplego/prom-metrics

go 1.21

replace github.com/talbs1986/simplego/metrics => ../metrics

replace github.com/talbs1986/simplego/app => ../app

require (
	github.com/prometheus/client_golang v1.17.0
	github.com/stretchr/testify v1.9.0
	github.com/talbs1986/simplego/app v0.0.0-20240528101415-c854be60989c
	github.com/talbs1986/simplego/metrics v0.0.0-20240528111019-aad0f97d2d90
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
