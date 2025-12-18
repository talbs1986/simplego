module github.com/talbs1986/simplego/chi-server

go 1.24

replace github.com/talbs1986/simplego/server => ../server

replace github.com/talbs1986/simplego/app => ../app

require (
	github.com/go-chi/chi/v5 v5.2.3
	github.com/talbs1986/simplego/app v0.0.0-20240528101415-c854be60989c
	github.com/talbs1986/simplego/server v0.0.0-20240528101415-c854be60989c
)
