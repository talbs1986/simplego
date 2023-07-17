module github.com/talbs1986/simplego/app

go 1.20

replace github.com/talbs1986/simplego/logger => ../logger

replace github.com/talbs1986/simplego/zerolog-logger => ../zerolog-logger

replace github.com/talbs1986/simplego/configs => ../configs

replace github.com/talbs1986/simplego/goenv-configs => ../goenv-configs

require (
	github.com/talbs1986/simplego/configs v0.0.0-20230717062942-0331e9d59f6a
	github.com/talbs1986/simplego/goenv-configs v0.0.0-20230717062942-0331e9d59f6a
	github.com/talbs1986/simplego/logger v0.0.0-20230717062942-0331e9d59f6a
	github.com/talbs1986/simplego/zerolog-logger v0.0.0-20230717062942-0331e9d59f6a
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.29.2-0.20230618233044-9070d49a1a41 // indirect
	github.com/sethvargo/go-envconfig v0.9.1-0.20230214025939-d0a807644a16 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
