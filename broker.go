package grpc_pubsub_broker

import (
	"sync"

	"github.com/tidwall/redcon"
)

type Broke struct {
	bind string
	mu   sync.Mutex
}

func (b *Broke) Run() error {
	panic("implement me")
}

func (b *Broke) Close(conn redcon.Conn, err error) error {
	panic("implement me")
}

func (b *Broke) Publish(i interface{}) error {
	panic("implement me")
}

func (b *Broke) Subscribe(i interface{}) error {
	panic("implement me")
}

func (b *Broke) Handler(conn redcon.Conn, cmd redcon.Command) {
	panic("implement me")
}

func (b *Broke) Accept(conn redcon.Conn) bool {
	panic("implement me")
}

type Broker interface {
	Run() error
	Close(conn redcon.Conn, err error) error
	Publish(interface{}) error
	Subscribe(interface{}) error
	Handler(conn redcon.Conn, cmd redcon.Command)
	Accept(conn redcon.Conn) bool
}
