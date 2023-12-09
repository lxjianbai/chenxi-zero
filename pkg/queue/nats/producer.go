package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
)

type Producer struct {
	conn *nats.EncodedConn
}

func NewProducer(c *NatsConfig) (*Producer, error) {
	nc, err := nats.Connect(c.Host)
	if err != nil {
		return nil, err
	}

	nec, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn: nec,
	}, nil
}

func (p *Producer) Publish(subject string, v any) error {
	return p.conn.Publish(subject, v)
}

func (p *Producer) Close() {
	if p.conn != nil {
		p.conn.Close()
	}
}
