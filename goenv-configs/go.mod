module github.com/talbs1986/simplego/goenv-configs

go 1.24

replace github.com/talbs1986/simplego/configs => ../configs

require (
	github.com/sethvargo/go-envconfig v1.3.0
	github.com/stretchr/testify v1.11.1
	github.com/talbs1986/simplego/configs v0.0.0-20230717062942-0331e9d59f6a
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
