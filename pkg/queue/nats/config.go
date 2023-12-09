package nats

import "github.com/nats-io/nats.go"

type NatsConfig struct {
	Host    string
	Options []nats.Option
}
