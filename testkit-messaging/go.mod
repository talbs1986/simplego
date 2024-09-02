module github.com/talbs1986/simplego/testkit-messaging

go 1.21

replace github.com/talbs1986/simplego/messaging => ../messaging

replace github.com/talbs1986/simplego/app => ../app

require (
	github.com/google/uuid v1.6.0
	github.com/talbs1986/simplego/messaging v0.0.0-20240528111019-aad0f97d2d90
)

require github.com/talbs1986/simplego/app v0.0.0-20240528101415-c854be60989c // indirect
