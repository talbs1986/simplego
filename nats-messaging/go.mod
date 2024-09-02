module github.com/talbs1986/simplego/nats-messaging

go 1.20

replace github.com/talbs1986/simplego/app => ../app

replace github.com/talbs1986/simplego/messaging => ../messaging

require github.com/talbs1986/simplego/app v0.0.0-20240528101415-c854be60989c

require (
	github.com/google/uuid v1.6.0
	github.com/nats-io/nats.go v1.37.0
	github.com/talbs1986/simplego/messaging v0.0.0-20240528101415-c854be60989c
)

require (
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)
