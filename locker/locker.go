package locker

import (
	"context"
	"sync"
	"time"
)

type Locker interface {
	Lock(ctx context.Context, name string, expire time.Duration) bool
	Unlock(ctx context.Context, name string) bool
}

var (
	client *Client
	once   sync.Once
)

func Init() {
	once.Do(func() {
		client = New()
	})
}

func Lock(ctx context.Context, name string, expire time.Duration) bool {
	if client == nil {
		return false
	}
	return client.Lock(ctx, name, expire)
}

func Unlock(ctx context.Context, name string) bool {
	if client == nil {
		return false
	}
	return client.Unlock(ctx, name)
}

type Client struct {
	names map[string]struct{}
}

func New() *Client {
	return &Client{
		names: make(map[string]struct{}),
	}
}

func (c *Client) Lock(ctx context.Context, name string, expire time.Duration) bool {
	return true
}

func (c *Client) Unlock(ctx context.Context, name string) bool {
	return false
}
