package scrape

import (
	"context"
	"encoding/json"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Dispatch interface {
	Router() string
	Consume(ctx context.Context, msg []byte) error
}

func NewService(handles ...Dispatch) *Service {
	s := &Service{
		routes: map[string]func(ctx context.Context, msg []byte) error{},
	}
	for _, v := range handles {
		s.routes[v.Router()] = v.Consume
	}
	return s
}

type Service struct {
	routes map[string]func(ctx context.Context, msg []byte) error
}

func (s *Service) Handle(ctx context.Context, msg pulsar.Message) error {
	handle, ok := s.routes[msg.Key()]
	if !ok {
		handle = s.defaultHandle
	}
	return handle(ctx, msg.Payload())
}

func (s *Service) defaultHandle(ctx context.Context, msg []byte) error {
	m, _ := json.Marshal(msg)
	log.Println(m)
	return nil
}
